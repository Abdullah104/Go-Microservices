install:
	go install github.com/swaggo/swag/cmd/swag@latest

swagger:
	swag fmt
	swag init -g handlers/products.go --parseDependency --parseInternal --parseDepth 1 -o ./swagger
