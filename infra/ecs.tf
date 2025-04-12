# 1つのECSクラスター内に複数のECSサービスを配置できる
resource "aws_ecs_cluster" "production-fargate-cluster" {
  name = var.ecs_cluster_name
}

resource "aws_alb" "ecs_cluster_alb" {
  name            = "${var.ecs_cluster_name}-ALB"
  # false = インターネット向けのパブリックALB
  # true = VPC内部向けのプライベートALB
  internal        = false
  security_groups = [aws_security_group.ecs_alb_security_group.id]
  subnets         = [
    aws_subnet.public-subnet-1.id,
    aws_subnet.public-subnet-2.id,
    aws_subnet.public-subnet-3.id
  ]

  tags = {
    Name = "${var.ecs_cluster_name}-ALB"
  }
}

resource "aws_alb_target_group" "default_target_group" {
  name     = "${var.ecs_cluster_name}-def-TG"
  port     = 80
  protocol = "HTTP"
  vpc_id   = aws_vpc.production-vpc.id

  tags = {
    Name = "${var.ecs_cluster_name}-def-TG"
  }
}

resource "aws_alb_listener" "ecs_alb_https_listener" {
  load_balancer_arn = aws_alb.ecs_cluster_alb.arn
  port              = 443
  protocol          = "HTTPS"
  ssl_policy        = "ELBSecurityPolicy-TLS13-1-2-2021-06"
  certificate_arn   = aws_acm_certificate.ecs_domain_certificate.arn

  default_action {
    type             = "forward"
    target_group_arn = aws_alb_target_group.default_target_group.arn
  }
}

resource "aws_route53_record" "ecs_load_balancer_record" {
  name = "*.${var.ecs_domain_name}"
  type = "A"
  zone_id = data.aws_route53_zone.ecs_domain.zone_id

  # なんか必要。これがないとカスタムドメインでアクセスできない
  alias {
    evaluate_target_health  = false
    name                    = aws_alb.ecs_cluster_alb.dns_name
    zone_id                 = aws_alb.ecs_cluster_alb.zone_id
  }
}


resource "aws_iam_role" "fargate_iam_role" {
  name = "${var.ecs_cluster_name}-Fargate-Task-Role"

  assume_role_policy = jsonencode({
    "Version": "2012-10-17",
    "Statement": [
    {
      "Effect": "Allow",
      "Principal": { // 誰に許可するか
        "Service": [
          "ecs.amazonaws.com",
          "ecs-tasks.amazonaws.com"
          ]
        },
        "Action": "sts:AssumeRole" // 何を許可するか。この記述は、このロールを「使用する」権限を与える
      }
    ]
  })
}

resource "aws_iam_role_policy" "fargate_iam_policy" {
  name = "${var.ecs_cluster_name}-Fargate-Task-Role-Policy"
  role = aws_iam_role.fargate_iam_role.id

  policy = jsonencode({
    "Version": "2012-10-17",
    "Statement": [
      {
        "Effect": "Allow",
        "Action": [
          "ecs:*",
          "ecr:*",
          "logs:*",
          "cloudwatch:*",
          "elasticloadbalancing:*",
          # ECS Execという機能を使用するために必要
          "ssmmessages:CreateControlChannel",
          "ssmmessages:CreateDataChannel",
          "ssmmessages:OpenControlChannel",
          "ssmmessages:OpenDataChannel",
          # DBのパスワードを取得するために必要
          "ssm:GetParameter",
          "ssm:GetParameters"
        ],
        "Resource": "*"
      }
    ]
  })
}

#
# Backendタスク定義(Gin)
#
locals {
  db_host = aws_db_instance.rds_instance.address
  ginapp_task_definition = templatefile("ginapp-task-definition.json", {
    ecs_service_name      = var.ecs_backend_service_name
    docker_image_url      = var.docker_backend_image_url
    container_port        = var.backend_container_port
    region                = var.region
    db_host               = local.db_host
    db_name               = var.db_name
    db_user               = var.db_username
    db_password_param     = var.db_password_ssm_parameter_name
    jwt_secret_key_param  = var.jwt_secret_key_param
  })
}

resource "aws_ecs_task_definition" "ginapp_task_def" {
  family                   = var.ecs_backend_service_name // タスク定義の名前グループ
  container_definitions     = local.ginapp_task_definition
  cpu                      = 512
  memory                   = var.memory
  requires_compatibilities = ["FARGATE"]
  network_mode             = "awsvpc"
  execution_role_arn       = aws_iam_role.fargate_iam_role.arn
  task_role_arn            = aws_iam_role.fargate_iam_role.arn
}

resource "aws_security_group" "ginapp_sg" {
  name        = "${var.ecs_backend_service_name}-SG"
  description = "Security group for backend Gin app"
  vpc_id      = aws_vpc.production-vpc.id

  ingress {
    from_port   = var.backend_container_port
    protocol    = "TCP"
    to_port     = var.backend_container_port
    cidr_blocks = [aws_vpc.production-vpc.cidr_block]
  }

  egress {
    from_port   = 0
    protocol    = "-1"
    to_port     = 0
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "${var.ecs_backend_service_name}-SG"
  }
}

resource "aws_alb_target_group" "ginapp_tg" {
  name        = "${var.ecs_backend_service_name}-TG"
  port        = var.backend_container_port
  protocol    = "HTTP"
  vpc_id      = aws_vpc.production-vpc.id
  target_type = "ip"

  health_check {
    path                = "/health"
    protocol            = "HTTP"
    matcher             = "200"
    interval            = 30
    timeout             = 5
    unhealthy_threshold = 3
    healthy_threshold   = 2
  }

  tags = {
    Name = "${var.ecs_backend_service_name}-TG"
  }
}
resource "aws_ecs_service" "ginapp_service" {
  name            = var.ecs_backend_service_name
  task_definition = aws_ecs_task_definition.ginapp_task_def.family
  desired_count   = var.desired_task_number
  cluster         = aws_ecs_cluster.production-fargate-cluster.name
  launch_type     = "FARGATE"
  enable_execute_command = true

  network_configuration {
    subnets         = [aws_subnet.private-subnet-1.id, aws_subnet.private-subnet-2.id, aws_subnet.private-subnet-3.id]
    security_groups = [aws_security_group.ginapp_sg.id]
    assign_public_ip = false
  }

  load_balancer {
    container_name   = var.ecs_backend_service_name
    container_port   = var.backend_container_port
    target_group_arn = aws_alb_target_group.ginapp_tg.arn
  }

  depends_on = [aws_ecs_task_definition.ginapp_task_def]
}

resource "aws_cloudwatch_log_group" "ginapp_log_group" {
  name = "${var.ecs_backend_service_name}-LogGroup"
}

resource "aws_alb_listener_rule" "ginapp_listener_rule" {
  listener_arn = aws_alb_listener.ecs_alb_https_listener.arn
  priority     = 101

  action {
    type             = "forward"
    target_group_arn = aws_alb_target_group.ginapp_tg.arn
  }

  condition {
    host_header {
      values = ["api.${var.ecs_domain_name}"]
    }
  }

  depends_on = [aws_alb_listener.ecs_alb_https_listener]
}

#
# Frontendタスク定義(Next.js)
#
locals {
  nextjsapp_task_definition = templatefile("nextjsapp-task-definition.json", {
    ecs_service_name      = var.ecs_frontend_service_name
    docker_image_url      = var.docker_frontend_image_url
    container_port        = var.frontend_container_port
    region                = var.region
    jwt_secret_key_param  = var.jwt_secret_key_param
    domain_name           = var.ecs_domain_name
  })
}

resource "aws_ecs_task_definition" "nextjsapp_task_def" {
  family                   = var.ecs_frontend_service_name
  container_definitions     = local.nextjsapp_task_definition
  cpu                      = 512
  memory                   = var.memory
  requires_compatibilities = ["FARGATE"]
  network_mode             = "awsvpc"
  execution_role_arn       = aws_iam_role.fargate_iam_role.arn
  task_role_arn            = aws_iam_role.fargate_iam_role.arn
}

resource "aws_security_group" "nextjsapp_sg" {
  name        = "${var.ecs_frontend_service_name}-SG"
  description = "Security group for frontend Next.js app"
  vpc_id      = aws_vpc.production-vpc.id

  ingress {
    from_port   = var.frontend_container_port
    protocol    = "TCP"
    to_port     = var.frontend_container_port
    cidr_blocks = [aws_vpc.production-vpc.cidr_block]
  }

  egress {
    from_port   = 0
    protocol    = "-1"
    to_port     = 0
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "${var.ecs_frontend_service_name}-SG"
  }
}

resource "aws_alb_target_group" "nextjsapp_tg" {
  name        = "${var.ecs_frontend_service_name}-TG"
  port        = var.frontend_container_port
  protocol    = "HTTP"
  vpc_id      = aws_vpc.production-vpc.id
  target_type = "ip"

  health_check {
    path                = "/"
    protocol            = "HTTP"
    matcher             = "200"
    interval            = 30
    timeout             = 5
    unhealthy_threshold = 3
    healthy_threshold   = 2
  }

  tags = {
    Name = "${var.ecs_frontend_service_name}-TG"
  }
}


resource "aws_ecs_service" "nextjsapp_service" {
  name            = var.ecs_frontend_service_name
  task_definition = aws_ecs_task_definition.nextjsapp_task_def.family
  desired_count   = var.desired_task_number
  cluster         = aws_ecs_cluster.production-fargate-cluster.name
  launch_type     = "FARGATE"
  enable_execute_command = true

  network_configuration {
    subnets         = [aws_subnet.private-subnet-1.id, aws_subnet.private-subnet-2.id, aws_subnet.private-subnet-3.id]
    security_groups = [aws_security_group.nextjsapp_sg.id]
    assign_public_ip = false
  }

  load_balancer {
    container_name   = var.ecs_frontend_service_name
    container_port   = var.frontend_container_port
    target_group_arn = aws_alb_target_group.nextjsapp_tg.arn
  }

  depends_on = [aws_ecs_task_definition.nextjsapp_task_def]
}

resource "aws_cloudwatch_log_group" "nextjsapp_log_group" {
  name = "${var.ecs_frontend_service_name}-LogGroup"
}

resource "aws_alb_listener_rule" "nextjsapp_listener_rule" {
  listener_arn = aws_alb_listener.ecs_alb_https_listener.arn
  priority     = 102

  action {
    type             = "forward"
    target_group_arn = aws_alb_target_group.nextjsapp_tg.arn
  }

  condition {
    host_header {
      values = ["www.${var.ecs_domain_name}"]
    }
  }

  depends_on = [aws_alb_listener.ecs_alb_https_listener]
}