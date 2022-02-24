## Lambda-sls

Functions with Lambda, DynamoDB to learn AWS serverless services.

### Create table local example

```
## Todo create
aws dynamodb create-table --table-name "todo" \
  --attribute-definitions AttributeName=id,AttributeType=S \
  AttributeName=name,AttributeType=S \
  --key-schema AttributeName=id,KeyType=HASH \
  AttributeName=name,KeyType=RANGE \
  --provisioned-throughput ReadCapacityUnits=1,WriteCapacityUnits=1 \
  --endpoint-url http://localhost:8000

## Todo delete
aws dynamodb delete-table --table-name "todo" --endpoint-url http://localhost:8000
```

### Development

```
docker compose up
```

#### .env

See [config.go](config/config.go)

```
### .env example
ENV: dev
SERVER_PORT: 4000
AWS_REGION: ap-northeast-1
AWS_ACCESS_KEY_ID: xxxxxxxxxxxxx
AWS_SECRET_ACCESS_KEY: xxxxxxxxxxxx
```

### Todo API request example

```
curl -X POST -H "Content-Type: application/json" \
http://localhost:4000/todos/create -d '{"id":"", "name":"new todo"}'

curl -X POST -H "Content-Type: application/json" \
http://localhost:4000/todos -d '{ "id":"", "name":"" }'

curl -X POST -H "Content-Type: application/json" \
http://localhost:4000/todos/update -d '{"id":"9d47755a-b7ce-4ba4-b581-a0533d9d2dd8", "name":"new name"}'

curl -X POST -H "Content-Type: application/json" \
http://localhost:4000/todos/delete -d '{"id":"9d47755a-b7ce-4ba4-b581-a0533d9d2dd8", "name":"todo"}'
```

### Image upload API request example

```
curl -X POST -F 'file=@/path/image.png' https://localhost:4000/upload
```

### Why POST method ?

<https://aws.amazon.com/premiumsupport/knowledge-center/api-gateway-lambda-template-invoke-error>

> If you have an API Gateway REST API with Lambda integration, then the API must invoke the backend Lambda function using the HTTP method,
> POST. If you use any other HTTP method (for example, ANY or GET) for the backend integration request, then invocation fails.

### Infrastructure example

<https://github.com/kis9a/terraform/tree/master/services/lambda-sls>

### Publish docker image

Push code to `release` branch.

GitHub secrets

- AWS_ACCESS_KEY_ID
- AWS_SECRET_ACCESS_KEY

[Publish GitHub Action](.github/workflows/publish.yml)

### Update Lambda function

```
## Example
aws lambda update-function-code --function-name lambda-sls \
--image-uri 298276046670.dkr.ecr.ap-northeast-1.amazonaws.com/lambda-sls:latest
```
