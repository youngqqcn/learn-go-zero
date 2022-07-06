all:genapi

genapi:
	goctl api go --api shuturl.api --dir ./


run:
	go run shuturl.go -f etc/shuturl-api.yaml

test:
	curl -i http://localhost:8888/from/you