.PHONY: temporal-up temporal-down temporal-logs \
	01-using-structs-service 01-using-structs-worker 01-using-structs-start \
	01-using-structs-history 01-using-structs-history-detailed \
	02-durable-execution-service 02-durable-execution-worker 02-durable-execution-start \
	02-durable-execution-history 02-durable-execution-history-detailed

NAME ?= Andy
LANGUAGE ?= fr

temporal-up:
	docker compose -f docker-compose-dev.yml up -d

temporal-down:
	docker compose -f docker-compose-dev.yml down

temporal-logs:
	docker compose -f docker-compose-dev.yml logs temporal

01-using-structs-service:
	go run ./01-using-structs/microservice

01-using-structs-worker:
	go run ./01-using-structs/worker

01-using-structs-start:
	go run ./01-using-structs/start "$(NAME)" "$(if $(filter command line,$(origin LANG)),$(LANG),$(LANGUAGE))"

01-using-structs-history:
	docker compose -f docker-compose-dev.yml exec temporal \
		temporal workflow show --workflow-id translation-workflow

01-using-structs-history-detailed:
	docker compose -f docker-compose-dev.yml exec temporal \
		temporal workflow show --workflow-id translation-workflow --detailed

02-durable-execution-service:
	go run ./02-durable-execution/microservice

02-durable-execution-worker:
	go run ./02-durable-execution/worker

02-durable-execution-start:
	go run ./02-durable-execution/start "$(NAME)" "$(if $(filter command line,$(origin LANG)),$(LANG),$(LANGUAGE))"

02-durable-execution-history:
	docker compose -f docker-compose-dev.yml exec temporal \
		temporal workflow show --workflow-id durable-execution-workflow

02-durable-execution-history-detailed:
	docker compose -f docker-compose-dev.yml exec temporal \
		temporal workflow show --workflow-id durable-execution-workflow --detailed
