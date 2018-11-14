build:
	make build_go
	make build_js
build_go:
	protoc src/*.proto \
		--proto_path src \
		--go_out=plugins=grpc:.
build_js:
	grpc_tools_node_protoc src/*.proto \
		--js_out=import_style=commonjs,binary:proto \
		--proto_path src \
		--grpc_out=proto \
		--plugin=protoc-gen-grpc=`which grpc_tools_node_protoc_plugin`
clean:
	rm -rf proto/*
runServer:
	go run src/server.go
runClient:
	go run src/client.go
