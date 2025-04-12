provider "aws" {
  region = var.region
}

terraform {
  # 空の中括弧 {} は、具体的な設定を別途行うことを示す
  backend "s3" {}
}
