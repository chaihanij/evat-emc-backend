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

build-production:
	@echo "============= Docs -> https://docs.docker.com/compose/compose-v2/ ============="
	docker compose  -f "docker-compose-evat-emc-production.yaml" build emc-service-production

build-develop:
	@echo "============= Docs -> https://docs.docker.com/compose/compose-v2/ ============="
	docker compose  -f "docker-compose-evat-emc-develop.yaml" build emc-service-dev

build-local:
	@echo "============= Docs -> https://docs.docker.com/compose/compose-v2/ ============= "
	docker compose  -f "docker-compose-evat-emc-local.yaml" build  emc-service-local


production:
	@echo "============= Docs -> https://docs.docker.com/compose/compose-v2/ ============="
	docker compose  -f "docker-compose-evat-emc-production.yaml" build
	docker compose  -f "docker-compose-evat-emc-production.yaml" stop
	docker compose  -f "docker-compose-evat-emc-production.yaml" rm -f
	docker compose  -f "docker-compose-evat-emc-production.yaml" pull
	docker compose  -f "docker-compose-evat-emc-production.yaml" up -d

develop:
	@echo "============= Docs -> https://docs.docker.com/compose/compose-v2/ ============="
	docker compose  -f "docker-compose-evat-emc-develop.yaml" build
	docker compose  -f "docker-compose-evat-emc-develop.yaml" stop
	docker compose  -f "docker-compose-evat-emc-develop.yaml" rm -f
	docker compose  -f "docker-compose-evat-emc-develop.yaml" pull
	docker compose  -f "docker-compose-evat-emc-develop.yaml" up -d

local:
	@echo "============= Docs -> https://docs.docker.com/compose/compose-v2/ ============= "
	docker compose  -f "docker-compose-evat-emc-local.yaml" build
	docker compose  -f "docker-compose-evat-emc-local.yaml" stop
	docker compose  -f "docker-compose-evat-emc-local.yaml" rm -f
	docker compose  -f "docker-compose-evat-emc-local.yaml" pull
	docker compose  -f "docker-compose-evat-emc-local.yaml" up -d


clean: down
	@echo "=============cleaning up============="
	docker system prune -f
	docker volume prune -f
	docker images prune -f

format:
	go fmt ./app/...

dep: ## Get the dependencies
	@go get -v -d ./...
	@go get -u github.com/golang/lint/golint

migrate-create:
	migrate create -ext sql -dir app/migrations $(name)

mockgen:
	mockgen -source=./app/layers/repositories/$(module_name)/init.go -destination=./app/mocks/$(module_name)/repo.go
	mockgen -source=./app/layers/usecases/$(module_name)/init.go -destination=./app/mocks/$(module_name)/use_case.go
