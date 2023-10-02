generate-swagger:
	swag init  --generalInfo "main.go" --output "./openapi" --outputTypes "json" --propertyStrategy "camelcase"
