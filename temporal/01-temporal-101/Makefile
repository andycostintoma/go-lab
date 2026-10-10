.PHONY: up logs worker hello history history-detailed farewell-service farewell-worker farewell finale-java-worker finale-worker finale

up:
	docker compose -f docker-compose-dev.yml up -d

logs:
	docker compose -f docker-compose-dev.yml logs temporal

worker:
	go run ./01-hello-workflow

hello:
	docker compose -f docker-compose-dev.yml exec temporal \
		temporal workflow execute \
		--type GreetSomeone \
		--task-queue greeting-tasks \
		--workflow-id hello-andy-1 \
		--input '"Andy"'

history:
	docker compose -f docker-compose-dev.yml exec temporal \
		temporal workflow show --workflow-id hello-andy-1

history-detailed:
	docker compose -f docker-compose-dev.yml exec temporal \
		temporal workflow show --workflow-id hello-andy-1 --detailed

farewell-service:
	go run ./02-farewell-workflow/microservice

farewell-worker:
	go run ./02-farewell-workflow/worker

farewell:
	go run ./02-farewell-workflow/start Andy

finale-java-worker:
	java -classpath ./03-finale-workflow/java-activity-and-worker-1.1.jar io.temporal.training.PdfCertWorker

finale-worker:
	go run ./03-finale-workflow/worker

finale:
	go run ./03-finale-workflow/start "Andy Toma"
