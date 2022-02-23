```
aws dynamodb create-table --table-name "todo" \
  --attribute-definitions AttributeName=id,AttributeType=S \
  AttributeName=name,AttributeType=S \
  --key-schema AttributeName=id,KeyType=HASH \
  AttributeName=name,KeyType=RANGE \
  --provisioned-throughput ReadCapacityUnits=1,WriteCapacityUnits=1 \
  --endpoint-url http://localhost:8000

aws dynamodb delete-table --table-name "todo" --endpoint-url http://localhost:8000
```

```
curl -X POST -H "Content-Type: application/json" http://localhost:4000/todos -d '{ "id": "1decaa0b-7cb1-45a5-80d3-81f1af825d93", "name": "apple" }'


curl -X POST -H "Content-Type: application/json" http://localhost:4000/delete -d '{"id":"9d47755a-b7ce-4ba4-b581-a0533d9d2dd8","name":"apple"}'
```
