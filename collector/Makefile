service := collector
tag := 1.0.0

.PHONY: build
build:
	docker build -t $(service):$(tag) .

.PHONY: tag
tag:
	docker tag $(service):$(tag) $(DOCKER_REGISTRY)/$(service):$(tag)

.PHONY: push
push:
	docker push $(DOCKER_REGISTRY)/$(service):$(tag)

.PHONY: load
load:
	minikube image load $(DOCKER_REGISTRY)/$(service):$(tag)
