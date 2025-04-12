ドメイン取得
ESR用意
cd go
make login
make deploy

cd ..
cd next



terraform実行
production.tfvarsの編集

cd ..
cd infra
terraform plan -var-file=production.tfvars
terraform apply -var-file=production.tfvars

terraform destroy