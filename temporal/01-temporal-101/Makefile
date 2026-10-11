.PHONY: temporal-up temporal-down temporal-logs \
	01-hello-workflow-worker 01-hello-workflow-start \
	01-hello-workflow-history 01-hello-workflow-history-detailed \
	02-farewell-workflow-service 02-farewell-workflow-worker 02-farewell-workflow-start \
	03-finale-workflow-java-worker 03-finale-workflow-worker 03-finale-workflow-start

temporal-up:
	docker compose -f docker-compose-dev.yml up -d

temporal-down:
	docker compose -f docker-compose-dev.yml down

temporal-logs:
	docker compose -f docker-compose-dev.yml logs temporal

01-hello-workflow-worker:
	go run ./01-hello-workflow

01-hello-workflow-start:
	docker compose -f docker-compose-dev.yml exec temporal \
		temporal workflow execute \
		--type GreetSomeone \
		--task-queue greeting-tasks \
		--workflow-id hello-andy-1 \
		--input '"Andy"'

01-hello-workflow-history:
	docker compose -f docker-compose-dev.yml exec temporal \
		temporal workflow show --workflow-id hello-andy-1

01-hello-workflow-history-detailed:
	docker compose -f docker-compose-dev.yml exec temporal \
		temporal workflow show --workflow-id hello-andy-1 --detailed

02-farewell-workflow-service:
	go run ./02-farewell-workflow/microservice

02-farewell-workflow-worker:
	go run ./02-farewell-workflow/worker

02-farewell-workflow-start:
	go run ./02-farewell-workflow/start Andy

03-finale-workflow-java-worker:
	java -classpath ./03-finale-workflow/java-activity-and-worker-1.1.jar io.temporal.training.PdfCertWorker

03-finale-workflow-worker:
	go run ./03-finale-workflow/worker

03-finale-workflow-start:
	go run ./03-finale-workflow/start "Andy Toma"
