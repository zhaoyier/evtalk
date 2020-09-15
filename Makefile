clean:
	@rm -rf rpc/*

gen:
	@protoc -I apidoc apidoc/proto/etadmin/*.proto --go_out=plugins=grpc:rpc/
	@protoc -I apidoc apidoc/proto/etwebapi/*.proto --go_out=plugins=grpc:rpc/