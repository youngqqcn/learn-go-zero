start-etcd:
	etcd

start-rpc:
	go run ./rpc/transform.go -f ./rpc/etc/transform.yaml

start-api:
	go run ./api/shorturl.go -f ./api/etc/shorturl-api.yaml

# all: start-etcd start-rpc start-api

test:
	curl -i "http://localhost:8888/shorten?url=http://www.xiaoheiban.cn"
	curl -i "http://localhost:8888/expand?shorten=f35b2a"

.PHONY: clean

clean:
	pkill etcd && rm -rf default.etcd

