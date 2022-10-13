# shorturl


参考 https://github.com/zeromicro/zero-doc/blob/main/doc/shorturl.md


```

# 安装etcd
sudo apt-get install etcd

# 通过docker安装mysql
docker run --name=mysql-test -itd -p 3306:3306 -e MYSQL_ROOT_PASSWORD=root123456 -d mysql
# 进入mysql
docker exec -it mysql-test /bin/bash
mysql -uroot -p
GRANT ALL PRIVILEGES ON *.* TO 'root'@'%' WITH GRANT OPTION;
GRANT ALL PRIVILEGES ON *.* TO 'root'@'localhost' WITH GRANT OPTION;
FLUSH PRIVILEGES;

# 安装 redis
docker run -itd --name redis-test -p 6379:6379 redis
```