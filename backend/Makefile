service := backend
tag := 1.2.0

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
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/topic/topic.proto
	python3 -m grpc_tools.protoc --proto_path=proto --python_out=proto/python --pyi_out=proto/python --grpc_python_out=proto/python proto/article/article.proto
	python3 -m grpc_tools.protoc --proto_path=proto --python_out=proto/python --pyi_out=proto/python --grpc_python_out=proto/python proto/topic/topic.proto
	protoc --js_out=import_style=commonjs,binary:. --grpc-web_out=import_style=typescript,mode=grpcweb:. proto/article/article.proto
	protoc --js_out=import_style=commonjs,binary:. --grpc-web_out=import_style=typescript,mode=grpcweb:. proto/topic/topic.proto

.PHONY: load
load:
	minikube image load $(DOCKER_REGISTRY)/$(service):$(tag)

.PHONY: test
test:
	go clean -testcache
	go test -v ./...

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: build_envoy
build_envoy:
	docker build -t grpcweb/envoy -f Dockerfile_envoy .
