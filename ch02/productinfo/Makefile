go_tidy:
	cd go/server && make tidy
	cd go/client && make tidy

go_build_service:
	cd go/server && make build

go_run_service:
	cd go/server && make run

go_build_client:
	cd go/client && make build

go_run_client:
	cd go/client && make run

go_gen:
	protoc -I proto/ proto/product_info.proto --go_out=plugins=grpc:go/server/ecommerce