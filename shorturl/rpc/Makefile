rpctemplate:
	goctl rpc -o transform.proto


genrpc:
	goctl rpc protoc transform.proto --go_out=. --go-grpc_out=. --zrpc_out=.


run:
	go run transform.go -f etc/transform.yaml