package main

import (
	"crypto/tls"
	"crypto/x509"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	_ "mygo/Configs"
	"mygo/Lib/Crontabs"
	"mygo/Lib/Routers"
	"mygo/Lib/Routers/Common"
	"mygo/Lib/Routers/Goods"
	"mygo/Lib/Routers/Orders"
	"mygo/services"
	"net/http"
	"strings"
)

var (
	port       = ":8080"
	ginE       *gin.Engine
	creds      credentials.TransportCredentials
	grpcServer *grpc.Server
)

func main() {
	go Crontabs.Init()
	initHttpsServer()
	initGrpcServer()
	initCreds()
	startServer()
}

func initCreds() error {
	certificate, err := tls.LoadX509KeyPair("crt/server.crt", "crt/server.key")
	if err != nil {
		return err
	}

	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("crt/ca.crt")
	if err != nil {
		return err
	}
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		return err
	}

	creds = credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{certificate},
		ClientAuth:   tls.RequireAndVerifyClientCert, // NOTE: this is optional!
		ClientCAs:    certPool,
	})
	return err
}

func startServer() {
	// gRPC是建⽴在HTTP/2版本之上，如果HTTP不是HTTP/2协议则必然⽆法提供gRPC⽀持
	http.ListenAndServeTLS(port, "crt/server.crt", "crt/server.key", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 如果将gRPC和Web服务放在⼀起，会导致gRPC和Web路径的冲突，在处理时我们需要区分两类服务。
		// 每个gRPC调⽤请求的Content-Type类型会被标注为"application/grpc"类型
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			//if err := ginE.Run(); err != nil {
			//	fmt.Printf("startup service failed, err: %v \n", err)
			//}
			ginE.ServeHTTP(w, r)
		}
	}))
}

func initGrpcServer() {
	grpcServer = grpc.NewServer(grpc.Creds(creds))
	//TODO register grpc service eg :pb.RegisterGreeterServer(grpcServer, new(myGrpcServer))
	services.RegisterUserServer(grpcServer, new(services.HelloworldService))
}

func initHttpsServer() {
	// 加载多个APP的路由配置
	Routers.Include(Goods.Routers, Orders.Routers, Common.Routers)
	// 初始化路由
	ginE = Routers.Init()
}
