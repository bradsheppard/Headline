service := backend
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

.PHONY: protoc
protoc:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/article/article.proto
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/interest/interest.proto
	python -m grpc_tools.protoc --proto_path=proto --python_out=proto/python --pyi_out=proto/python --grpc_python_out=proto/python proto/article/article.proto
	python -m grpc_tools.protoc --proto_path=proto --python_out=proto/python --pyi_out=proto/python --grpc_python_out=proto/python proto/interest/interest.proto
	protoc --js_out=import_style=commonjs:. --grpc-web_out=import_style=commonjs,mode=grpcwebtext:. proto/article/article.proto
	protoc --js_out=import_style=commonjs:. --grpc-web_out=import_style=commonjs,mode=grpcwebtext:. proto/interest/interest.proto

.PHONY: load
load:
	minikube image load $(DOCKER_REGISTRY)/$(service):$(tag)

.PHONY: test
test:
	go clean -testcache
	go test ./...

