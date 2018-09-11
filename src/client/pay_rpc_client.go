package client

import (
	"fmt"
	"net/http"
)

type PayClient struct {
}

func NewPayClient() *PayClient {

	return nil
}

func (p *PayClient) Dispatch(method string, resp http.ResponseWriter, req *http.Request) {
	fmt.Println("pay收到请求")

}