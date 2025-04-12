variable "region" {
  default     = "ap-northeast-1"
  description = "AWS Region"
}

variable "vpc_cidr" {
  description = "VPC CIDR Block"
}


variable "public_subnet_1_cidr" {
  description = "CIDR Block for Public Subnet 1"
}

variable "public_subnet_2_cidr" {
  description = "CIDR Block for Public Subnet 2"
}

variable "public_subnet_3_cidr" {
  description = "CIDR Block for Public Subnet 3"
}

variable "private_subnet_1_cidr" {
  description = "CIDR Block for Private Subnet 1"
}

variable "private_subnet_2_cidr" {
  description = "CIDR Block for Private Subnet 2"
}

variable "private_subnet_3_cidr" {
  description = "CIDR Block for Private Subnet 3"
}


variable "ecs_cluster_name" {}
variable "internet_cidr_block" {}
variable "ecs_domain_name" {}

# Frontend(Next.js)用
variable "ecs_frontend_service_name" {}
variable "docker_frontend_image_url" {}
variable "frontend_container_port" {}

# Backend(Gin)用
variable "ecs_backend_service_name" {}
variable "docker_backend_image_url" {}
variable "backend_container_port" {}

variable "memory" {}
variable "desired_task_number" {}


variable "db_name" {
  description = "Database name"
}

variable "db_username" {
  description = "Database master username"
}

variable "db_password_ssm_parameter_name" {
  description = "SSM Parameter name where DB password is stored"
}

variable "jwt_secret_key_param" {
  description = "SSM Parameter name where JWT secret key is stored"
}

variable "db_engine" {
  default     = "mysql"
  description = "Database engine"
}

variable "db_engine_version" {
  default     = "8.0"
  description = "Database engine version"
}

variable "db_instance_class" {
  default     = "db.t3.micro"
  description = "Instance class for DB"
}

variable "db_allocated_storage" {
  default     = 20
  description = "Allocated storage in GB"
}