genmodel:
	cd service/user/model && goctl model mysql ddl -src user.sql -dir . -c && cd -


gen-user-api:
	cd service/user/api && goctl api go -api user.api -dir .  && cd -

gen-search-api:
	cd service/search/api && goctl api go -api search.api -dir .  && cd -

start-user-api:
	 go run service/user/api/user.go -f service/user/api/etc/user-api.yaml

start-search-api:
	 go run service/search/api/search.go -f service/search/api/etc/search-api.yaml


# 测试用户登录，返回jwt
test-user-login:
	curl -i -X POST http://127.0.0.1:8888/user/login -H 'Content-Type: application/json' -d '{ "username":"666","password":"123456"}'

# 注意其中Authorization是用登录的jwt
test-jwt:
	curl -i -X GET 'http://127.0.0.1:8889/search/do?name=%E8%A5%BF%E6%B8%B8%E8%AE%B0'  -H 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjU2NjUxNTgsImlhdCI6MTY2NTY2MTU1OCwidXNlcklkIjoxfQ.Hqxn5D46ex7eDZhIG2F5AHCKAgBcBsrdcoyBFEE-ImY'