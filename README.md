# go_code
Simple go code

# 自动生成数据库:reverse
    go get xorm.io/reverse
    编写yml数据库链接文件db.yml
    执行：reverse -f db.yml
    注意，不支持json名称，需手动改
# 自动生成数据库:xorm
    go get github.com/go-xorm/cmd/xorm
    go get github.com/go-xorm/xorm
    xorm/xorm reverse mysql root:123456@tcp\(127.0.0.1:3308\)/test?charset=utf8 xorm/templates/goxorm Entities/Databases/
    注意：xorm方式不支持unsigned,可以修改测试环境字段类型后执行

# 测试grpc gin
    GODEBUG=x509ignoreCN=0 go run grpc-test.go
    GODEBUG=x509ignoreCN=0 go run https-test.go

# rabbitmq
    docker run -d --hostname my-rabbit --name rabbitmq -p 15672:15672 -p 5672:5672 -p 25672:25672 -e RABBITMQ_DEFAULT_USER=rbmq -e RABBITMQ_DEFAULT_PASS=rbmq rabbitmq:3-management
    5672：连接生产者、消费者的端口。 
    15672：WEB管理页面的端口。
    25672：分布式集群的端口。

# mysql
    docker run --name mysql -d -p 3306:3306 -v /Users/missyourlove/docker_data/mysql/data:/var/lib/mysql \
    -v /Users/missyourlove/docker_data/mysql/conf/my.cnf:/etc/my.cnf \
    -e MYSQL_ROOT_PASSWORD=密码 mysql