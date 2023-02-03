example-docs-gen:
	@echo "============= Docs -> https://github.com/swaggo/swag ============= "
	swag init --dir ./example/docs-generator --swagger ./example/docs-generator/docs/swagger/

docs-gen:
	@echo "============= Docs -> https://github.com/swaggo/swag ============= "
	swag init --dir ./app --output ./app/docs/swagger/
	npx redoc-cli bundle ./app/docs/swagger/swagger.json --output ./app/docs/swagger/index.html

docs-gen-no-html:
	@echo "============= Docs -> https://github.com/swaggo/swag ============= "
	swag init --dir ./app --output ./app/docs/swagger/

