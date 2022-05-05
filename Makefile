.PHONY: init
init:
	docker-compose up --build

.PHONY: up-mid
up-mid:
	docker-compose up redis minio

.PHONY: up-worker
up-worker:
	docker-compose up worker

.PHONY: event
event:
	curl -XPOST "http://localhost:8080/2015-03-31/functions/function/invocations" -d '{}'