ドメイン取得
ECR用意
cd go
make login
make deploy

cd ..
cd next

make deploy



terraform実行
production.tfvarsの編集
s3で適当な名前のバケットを作成
infrastructure-prod.configを編集

parameter storeに
/myapp/prod/db_password
password
/myapp/prod/jwt_secret_key
jwt-secret-key

cd ..
cd infra
terraform init -backend-config="infrastructure-prod.config"
terraform plan -var-file=production.tfvars
terraform apply -var-file=production.tfvars

terraform destroy -var-file=production.tfvars


www.domain