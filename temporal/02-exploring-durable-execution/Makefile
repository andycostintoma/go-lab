.PHONY: temporal-up temporal-down temporal-logs \
	01-using-structs-service 01-using-structs-worker 01-using-structs-start \
	01-using-structs-history 01-using-structs-history-detailed \
	02-durable-execution-service 02-durable-execution-worker 02-durable-execution-start \
	02-durable-execution-history 02-durable-execution-history-detailed \
	03-testing-code-service 03-testing-code-worker 03-testing-code-start \
	03-testing-code-history 03-testing-code-history-detailed \
	03-testing-code-test 03-testing-code-test-activity \
	03-testing-code-test-workflow 03-testing-code-test-mock \
	04-debug-activity-worker 04-debug-activity-start \
	04-debug-activity-history 04-debug-activity-history-detailed \
	04-debug-activity-test 04-debug-activity-test-activity 04-debug-activity-test-workflow

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

03-testing-code-service:
	go run ./03-testing-code/microservice

03-testing-code-worker:
	go run ./03-testing-code/worker

03-testing-code-start:
	go run ./03-testing-code/start "$(NAME)" "$(if $(filter command line,$(origin LANG)),$(LANG),$(LANGUAGE))"

03-testing-code-history:
	docker compose -f docker-compose-dev.yml exec temporal \
		temporal workflow show --workflow-id translation-workflow

03-testing-code-history-detailed:
	docker compose -f docker-compose-dev.yml exec temporal \
		temporal workflow show --workflow-id translation-workflow --detailed

03-testing-code-test:
	go test -v ./03-testing-code/...

03-testing-code-test-activity:
	go test -v ./03-testing-code -run '^(TestSuccessfulTranslateActivity|TestFailedTranslateActivity)'

03-testing-code-test-workflow:
	go test -v ./03-testing-code -run '^TestSuccessfulCompleteFrenchTranslation$$'

03-testing-code-test-mock:
	go test -v ./03-testing-code -run '^TestSuccessfulTranslationWithMocks$$'

04-debug-activity-worker:
	go run ./04-debug-activity/worker

04-debug-activity-start:
	go run ./04-debug-activity/start

04-debug-activity-history:
	docker compose -f docker-compose-dev.yml exec temporal \
		temporal workflow show --workflow-id pizza-workflow-order-Z1238

04-debug-activity-history-detailed:
	docker compose -f docker-compose-dev.yml exec temporal \
		temporal workflow show --workflow-id pizza-workflow-order-Z1238 --detailed

04-debug-activity-test:
	go test -v ./04-debug-activity/...

04-debug-activity-test-activity:
	go test -v ./04-debug-activity -run '^Test(GetDistance|SendBill)'

04-debug-activity-test-workflow:
	go test -v ./04-debug-activity -run '^Test(SuccessfulPizzaOrder|FailedPizzaOrderCustomerOutsideDeliveryArea)$$'
