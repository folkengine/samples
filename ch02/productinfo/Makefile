go_tidy:
	cd go/server && go mod tidy
	cd go/client && go mod tidy

go_build_service:
	cd go/server && go build -v -o bin/server

go_run_service:
	cd go/server && ./bin/server

go_build_client:
	cd go/client && go build -v -o bin/client

go_run_client:
	cd go/client && ./bin/client

go_gen:
	protoc -I proto/ proto/product_info.proto --go_out=plugins=grpc:go/server/ecommerce