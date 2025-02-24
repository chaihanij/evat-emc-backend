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

debug-develop:
	@echo "============= Docs -> https://docs.docker.com/compose/compose-v2/ ============="
	docker compose  -f "docker-compose-evat-emc-develop-debug.yaml" build
	docker compose  -f "docker-compose-evat-emc-develop-debug.yaml" stop
	docker compose  -f "docker-compose-evat-emc-develop-debug.yaml" rm -f
	docker compose  -f "docker-compose-evat-emc-develop-debug.yaml" pull
	docker compose  -f "docker-compose-evat-emc-develop-debug.yaml" up -d

local:
	@echo "============= Docs -> https://docs.docker.com/compose/compose-v2/ ============= "
	docker compose  -f "docker-compose-evat-emc-local.yaml" build
	docker compose  -f "docker-compose-evat-emc-local.yaml" stop
	docker compose  -f "docker-compose-evat-emc-local.yaml" rm -f
	docker compose  -f "docker-compose-evat-emc-local.yaml" pull
	docker compose  -f "docker-compose-evat-emc-local.yaml" up -d

up-production:
	@echo "============= Docs -> https://docs.docker.com/compose/compose-v2/ ============="
	docker compose  -f "docker-compose-evat-emc-production.yaml" up emc-service-production

up-develop:
	@echo "============= Docs -> https://docs.docker.com/compose/compose-v2/ ============="
	docker compose  -f "docker-compose-evat-emc-develop.yaml" up emc-service-dev

up-develop-debug:
	@echo "============= Docs -> https://docs.docker.com/compose/compose-v2/ ============="
	docker compose  -f "docker-compose-evat-emc-develop-debug.yaml" up emc-service-dev

up-local:
	@echo "============= Docs -> https://docs.docker.com/compose/compose-v2/ ============="
	docker compose  -f "docker-compose-evat-emc-local.yaml" up emc-service-local

stop-production:
	@echo "============= Docs -> https://docs.docker.com/compose/compose-v2/ ============="
	docker compose  -f "docker-compose-evat-emc-production.yaml" stop emc-service-production

stop-develop:
	@echo "============= Docs -> https://docs.docker.com/compose/compose-v2/ ============="
	docker compose  -f "docker-compose-evat-emc-develop.yaml" stop emc-service-dev

stop-develop-debug:
	@echo "============= Docs -> https://docs.docker.com/compose/compose-v2/ ============="
	docker compose  -f "docker-compose-evat-emc-develop-debug.yaml" stop emc-service-dev

stop-local:
	@echo "============= Docs -> https://docs.docker.com/compose/compose-v2/ ============="
	docker compose  -f "docker-compose-evat-emc-local.yaml" stop emc-service-local

ps-production:
	@echo "============= Docs -> https://docs.docker.com/compose/compose-v2/ ============="
	docker compose  -f "docker-compose-evat-emc-production.yaml" ps

ps-develop:
	@echo "============= Docs -> https://docs.docker.com/compose/compose-v2/ ============="
	docker compose  -f "docker-compose-evat-emc-develop.yaml" ps

ps-local:
	@echo "============= Docs -> https://docs.docker.com/compose/compose-v2/ ============= "
	docker compose  -f "docker-compose-evat-emc-local.yaml" ps

logs-production:
	@echo "============= Docs -> https://docs.docker.com/compose/compose-v2/ ============="
	docker compose  -f "docker-compose-evat-emc-production.yaml" logs emc-service-production

logs-develop:
	@echo "============= Docs -> https://docs.docker.com/compose/compose-v2/ ============="
	docker compose  -f "docker-compose-evat-emc-develop.yaml" logs  emc-service-dev

logs-local:
	@echo "============= Docs -> https://docs.docker.com/compose/compose-v2/ ============= "
	docker compose  -f "docker-compose-evat-emc-local.yaml" logs -f emc-service-local

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
