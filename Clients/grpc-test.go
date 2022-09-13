package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
	"mygo/services"
)

var tsname = "server.io"

func doClientWork() {
	certificate, err := tls.LoadX509KeyPair("../crt/client.crt", "../crt/client.key")
	if err != nil {
		log.Fatal(err)
	}

	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("../crt/ca.crt")
	if err != nil {
		log.Fatal(err)
	}
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatal("failed to append ca certs")
	}

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{certificate},
		ServerName:   tsname, // NOTE: this is required!
		RootCAs:      certPool,
	})

	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := services.NewUserClient(conn)

	r, err := c.SayHello(context.Background(), &services.RequestMessage{Id: 1})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("doClientWork: %s", r.Name)
}

func main() {
	doClientWork()
}
