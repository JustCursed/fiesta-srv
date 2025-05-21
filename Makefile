PROTO_DIR = ./protos
PB_DIR = ./pb
MODULE = module=void-studio.net/fiesta/pb

setup:
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest; \
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest; \
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest; \
	proto

proto:
	@[ -n "$$(ls -A $(PB_DIR))" ] && rm -rf $(PB_DIR)/*; \
	protoc --go_out=$(PB_DIR) --go_opt=$(MODULE) \
	--go-grpc_out=$(PB_DIR) --go-grpc_opt=$(MODULE) \
	$(PROTO_DIR)/*.proto;

# download google api for proto (working only for linux)
googleapis:
	git clone https://github.com/googleapis/googleapis.git; \
	sudo mkdir -p /usr/include/google; \
	sudo cp -r googleapis/google/api /usr/include/google; \
	rm -rf googleapis;
