.PHONY: protoc
protoc:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/article/article.proto
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/interest/interest.proto
	
	python -m grpc_tools.protoc --proto_path=proto --python_out=proto/python --pyi_out=proto/python --grpc_python_out=proto/python proto/article/article.proto
	python -m grpc_tools.protoc --proto_path=proto --python_out=proto/python --pyi_out=proto/python --grpc_python_out=proto/python proto/interest/interest.proto

.PHONY: test
test:
	go clean -testcache
	go test ./...

