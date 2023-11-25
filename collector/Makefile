service := collector
tag := 1.1.1

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
	pylint main.py collector/ test/ --rcfile .pylintrc

.PHONY: load
load:
	minikube image load $(DOCKER_REGISTRY)/$(service):$(tag) --overwrite=true

.PHONY: protoc
protoc:
	protoc --go_out=proto/golang --go_opt=paths=source_relative --go-grpc_out=proto/golang --go-grpc_opt=paths=source_relative proto/collection/collection.proto
	python -m grpc_tools.protoc --proto_path=proto --python_out=collector --pyi_out=collector --grpc_python_out=collector proto/collection/collection.proto

.PHONY: test
test:
	python -m pytest test
