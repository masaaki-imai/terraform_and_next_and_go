resource "aws_db_subnet_group" "db_subnet_group" {
  name       = "${lower(var.ecs_cluster_name)}-db-subnet-group"
  subnet_ids = [
    aws_subnet.private-subnet-1.id,
    aws_subnet.private-subnet-2.id,
    aws_subnet.private-subnet-3.id
  ]

  tags = {
    Name = "${var.ecs_cluster_name}-db-subnet-group"
  }
}

resource "aws_security_group" "rds_sg" {
  name        = "${var.ecs_cluster_name}-RDS-SG"
  description = "Security group for RDS"
  vpc_id      = aws_vpc.production-vpc.id

  # ECSタスクからのアクセスを許可（MySQLの場合: 3306ポート）
  ingress {
    from_port       = 3306
    protocol        = "TCP"
    to_port         = 3306
    security_groups = [aws_security_group.ginapp_sg.id, aws_security_group.nextjsapp_sg.id]
    # ginapp_sgまたはnextjsapp_sgからのアクセスを許可
    # バックエンドのみがDBへアクセスするならginapp_sgのみでOK
  }

  egress {
    from_port   = 0
    protocol    = "-1"
    to_port     = 0
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "${var.ecs_cluster_name}-RDS-SG"
  }
}

resource "aws_db_instance" "rds_instance" {
  allocated_storage    = var.db_allocated_storage
  engine               = var.db_engine
  engine_version       = var.db_engine_version
  instance_class       = var.db_instance_class
  identifier           = var.db_name
  username             = var.db_username
  password             = data.aws_ssm_parameter.db_password.value
  skip_final_snapshot  = true
  db_subnet_group_name = aws_db_subnet_group.db_subnet_group.name
  vpc_security_group_ids = [aws_security_group.rds_sg.id]
  publicly_accessible = false
  multi_az = false

  tags = {
    Name = "${var.ecs_cluster_name}-RDS-Instance"
  }
}

data "aws_ssm_parameter" "db_password" {
  name = var.db_password_ssm_parameter_name
  with_decryption = true
}