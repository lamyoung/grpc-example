package util

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"

	"google.golang.org/grpc/credentials"
)

// GetServerCredentials 服务端证书
func GetServerCredentials() credentials.TransportCredentials {
	cert, err := tls.LoadX509KeyPair("../cert/server.pem", "../cert/server.key")
	if err != nil {
		log.Fatalf("加载服务端证书失败, err: %v\n", err)
	}

	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("../cert/ca.pem")
	if err != nil {
		log.Fatalf("读取公钥文件失败: %v\n", err)
	}

	certPool.AppendCertsFromPEM(ca)

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	})
	return creds
}

// GetClientCredentials 客户端证书
func GetClientCredentials() credentials.TransportCredentials {
	cert, err := tls.LoadX509KeyPair("../cert/client.pem", "../cert/client.key")
	if err != nil {
		log.Fatalf("加载客户端证书失败, err: %v\n", err)
	}

	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("../cert/ca.pem")
	if err != nil {
		log.Fatalf("读取公钥文件失败: %v\n", err)
	}

	certPool.AppendCertsFromPEM(ca)

	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ServerName:   "localhost",
		RootCAs:      certPool,
	})
	return creds
}
