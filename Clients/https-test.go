package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetHttps(url, caCertPath, certFile, keyFile string) ([]byte, error) {

	// 创建证书池及各类对象
	var pool *x509.CertPool // 我们要把一部分证书存到这个池中
	var client *http.Client
	var resp *http.Response
	var body []byte
	var err error

	var caCrt []byte // 根证书
	caCrt, err = ioutil.ReadFile(caCertPath)
	pool = x509.NewCertPool()
	if err != nil {
		return nil, err
	}
	pool.AppendCertsFromPEM(caCrt)

	var cliCrt tls.Certificate // 具体的证书加载对象
	cliCrt, err = tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, err
	}

	// 把上面的准备内容传入 client
	client = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:      pool,
				Certificates: []tls.Certificate{cliCrt},
				ServerName:   "server.io",
			},
		},
	}

	// Get 请求
	resp, err = client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer client.CloseIdleConnections()

	return body, nil
}

func main() {
	rs, err := GetHttps("https://localhost:8080/username?id=1", "../crt/ca.crt", "../crt/client.crt", "../crt/client.key")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(rs))
}
