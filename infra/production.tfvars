# VPC variables for production
vpc_cidr = "10.0.0.0/16"
public_subnet_1_cidr = "10.0.1.0/24"
public_subnet_2_cidr = "10.0.2.0/24"
public_subnet_3_cidr = "10.0.5.0/24"
private_subnet_1_cidr = "10.0.3.0/24"
private_subnet_2_cidr = "10.0.4.0/24"
private_subnet_3_cidr = "10.0.6.0/24"

ecs_domain_name = "lecsite.com"# ※YOU HAVE TO CHANGE THIS!
ecs_cluster_name = "Production-ECS-Cluster"
internet_cidr_block = "0.0.0.0/0"

# Frontend設定
ecs_frontend_service_name = "myapp-next"
docker_frontend_image_url = "" # ※YOU HAVE TO CHANGE THIS!
frontend_container_port = 3000

# Backend設定
ecs_backend_service_name = "myapp-go"
docker_backend_image_url = "" # ※YOU HAVE TO CHANGE THIS!
backend_container_port = 8080

memory = 1024
desired_task_number = "2"


db_name = "myappdb"
db_username = "dbuser"
db_password_ssm_parameter_name = "/myapp/prod/db_password"

jwt_secret_key_param = "/myapp/prod/jwt_secret_key"