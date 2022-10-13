all:genapi

genapi:
	goctl api go --api shorturl.api --dir ./


run:
	go run shorturl.go -f etc/shorturl-api.yaml

test:
	curl -i http://localhost:8888/from/you