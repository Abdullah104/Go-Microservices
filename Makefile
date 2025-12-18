install:
	go install github.com/swaggo/swag/cmd/swag@latest

swagger:
	swag fmt
	swag init -g handlers/products.go --parseDependency --parseInternal --parseDepth 1 -o ./swagger

generate-client:
	openapi-generator generate -i swagger/swagger.json -g go -o clients/products --additional-properties=packageName=products_client,isGoSubmodule=true --git-user-id=Abdullah104 --git-repo-id=Go-Microservices
