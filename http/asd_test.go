package asd

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"testing"

	"golang.org/x/net/http2"
)

func TestAsd(t *testing.T) {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	http2.ConfigureTransport(tr)
	client := &http.Client{Transport: tr}

	req, err := http.NewRequest("POST", "https://localhost:9005/service.OrderService/GetOrder", strings.NewReader("OrderId=123"))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Content-type", "application/grpc")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(body))
}
