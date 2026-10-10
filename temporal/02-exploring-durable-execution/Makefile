.PHONY: up logs service worker start history history-detailed

NAME ?= Andy
LANGUAGE ?= fr

up:
	docker compose -f docker-compose-dev.yml up -d

logs:
	docker compose -f docker-compose-dev.yml logs temporal

service:
	go run ./01-using-structs/microservice

worker:
	go run ./01-using-structs/worker

start:
	go run ./01-using-structs/start "$(NAME)" "$(if $(filter command line,$(origin LANG)),$(LANG),$(LANGUAGE))"

history:
	docker compose -f docker-compose-dev.yml exec temporal \
		temporal workflow show --workflow-id translation-workflow

history-detailed:
	docker compose -f docker-compose-dev.yml exec temporal \
		temporal workflow show --workflow-id translation-workflow --detailed
