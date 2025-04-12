AWS_ACCOUNT_ID ?= $(shell aws sts get-caller-identity --query "Account" --output text)
AWS_REGION ?= ap-northeast-1# YOU HAVE TO CHANGE THIS!
ECR_REPO_NAME = myapp-next
ECR_URI = $(AWS_ACCOUNT_ID).dkr.ecr.$(AWS_REGION).amazonaws.com/$(ECR_REPO_NAME)

login:
	aws ecr get-login-password --region $(AWS_REGION) | docker login --username AWS --password-stdin $(AWS_ACCOUNT_ID).dkr.ecr.$(AWS_REGION).amazonaws.com

build:
	docker build --platform linux/x86_64 -t $(ECR_REPO_NAME) .

tag:
	docker tag $(ECR_REPO_NAME) $(ECR_URI):latest

push:
	docker push $(ECR_URI):latest

echo_ecr_url:
	@echo "$(ECR_URI)"

deploy: build tag push