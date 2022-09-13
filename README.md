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