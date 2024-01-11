service := orchestrator
tag := 1.3.0

.PHONY: build
build:
	docker build -t $(service):$(tag) .

.PHONY: tag
tag:
	docker tag $(service):$(tag) $(DOCKER_REGISTRY)/$(service):$(tag)

.PHONY: push
push:
	docker push $(DOCKER_REGISTRY)/$(service):$(tag)

.PHONY: lint
lint:
	pylint main.py orchestrator/ test/ --rcfile .pylintrc

.PHONY: load
load:
	minikube image load $(DOCKER_REGISTRY)/$(service):$(tag) --overwrite=true

.PHONY: run
run:
	kubectl create job orchestrator-manual --from=cronjobs/orchestrator

