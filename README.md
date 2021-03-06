# index_indicators

## 環境構築
```
git clone

config.ini .env を設置

docker compose up
```

# テスト
 `cd server`
`go test -v -cover ./app/controllers/`
<br>
<br>
# ECS設定
### login 
`aws ecr get-login-password --region region | docker login --username AWS --password-stdin ID`
### build
`docker build -t index_indicators:v1 .`

### tag
`docker tag index_indicators:v1 ID/index_indicators:latest`

### ECR push
`docker push ID/index_indicators:latest`

### task definition
`aws ecs register-task-definition --cli-input-json file://task-definition.json`

### create service
`aws ecs create-service --cli-input-json file://ecs-service.json`

# アーキテクチャ
<img src="./assets/server.svg">
