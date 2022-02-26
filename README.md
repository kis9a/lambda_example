## Lambda-sls

Serverless Functions with Lambda, DynamoDB, APIGatewayProxy... to learn.

### Development

```
docker compose up
```

#### .env

See [config.go](config/config.go)

```
### .env example
ENV: dev
AWS_REGION: ap-northeast-1
AWS_ACCESS_KEY_ID: AKxxxxxxxxxxxxxxxxxx
AWS_SECRET_ACCESS_KEY: 6txxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
S3_TODO_BUCKET: kis9a-lambda-upload
DB_TODO_TABLE: lambda-sls-todo
DB_ENDPOINT: http://127.0.0.1:8000
DB_DISABLE_SSL: false
SERVER_PORT: 4000
```

### Create dynamodb table local example

```
## Todo create
aws dynamodb create-table --table-name "lambda-sls-todo" \
  --attribute-definitions AttributeName=id,AttributeType=S \
  --key-schema AttributeName=id,KeyType=HASH \
  --provisioned-throughput ReadCapacityUnits=1,WriteCapacityUnits=1 \
  --endpoint-url http://localhost:8000

## Todo delete
aws dynamodb delete-table --table-name "lambda-sls-todo" --endpoint-url http://localhost:8000
```

### Todo API request example

```
# create todo item
curl -X POST -H "Content-Type: application/json" \
http://localhost:4000/todos/create -d '{"name":"new todo"}'

# read todo
curl -X POST -H "Content-Type: application/json" \
"http://localhost:4000/todos?limit=16"

# update todo
curl -X POST -H "Content-Type: application/json" \
http://localhost:4000/todos/update -d '{"id":"9d47755a-b7ce-4ba4-b581-a0533d9d2dd8", "name":"new name"}'

# delete todo
curl -X POST -H "Content-Type: application/json" \
http://localhost:4000/todos/delete -d '{"id":"9d47755a-b7ce-4ba4-b581-a0533d9d2dd8", "name":"todo"}'
```

### Image upload API request example

```
# upload image
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
