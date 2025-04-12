resource "aws_security_group" "ecs_security_group" {
  name        = "${var.ecs_cluster_name}-SG"
  description = "Security group for ECS tasks"
  vpc_id      = aws_vpc.production-vpc.id

  ingress {
    from_port   = 32768
    protocol    = "TCP"
    to_port     = 65535
    cidr_blocks = [aws_vpc.production-vpc.cidr_block]
  }

  egress {
    from_port   = 0
    protocol    = "-1" # すべてのプロトコル
    to_port     = 0
    cidr_blocks = [var.internet_cidr_block]
  }

  tags = {
    Name = "${var.ecs_cluster_name}-SG"
  }
}

resource "aws_security_group" "ecs_alb_security_group" {
  name        = "${var.ecs_cluster_name}-ALB-SG"
  description = "Security group for ALB to traffic for ECS cluster"
  vpc_id      = aws_vpc.production-vpc.id

  ingress {
    from_port   = 443
    protocol    = "TCP"
    to_port     = 443
    cidr_blocks = [var.internet_cidr_block]
  }

  egress {
    #     from_port = 0 と to_port = 0: すべてのポート範囲を許可
    # protocol = "-1": すべてのプロトコルを許可（"-1"は全プロトコルを示すAWS特有の表記）
    from_port   = 0
    protocol    = "-1"
    to_port     = 0
    cidr_blocks = [var.internet_cidr_block]
  }
}