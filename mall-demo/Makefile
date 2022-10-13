genrpc:
	cd mall/user/rpc && goctl rpc protoc user.proto --go_out=./types --go-grpc_out=./types --zrpc_out=. && cd -


genapi:
	cd order/api && goctl api go -api order.api -dir . && cd -

start-rpc:
	go run mall/user/rpc/user.go -f mall/user/rpc/etc/user.yaml


start-api:
	go run mall/order/api/order.go -f mall/order/api/etc/order.yaml


test:
	curl -i -X GET http://localhost:8888/api/order/get/1