package client

import (
	"fmt"
	"net/http"
)

type VipClient struct {
}

func NewVipClient() *VipClient {

	return nil

}
func (v *VipClient) Dispatch(method string, resp http.ResponseWriter, req *http.Request) {
	fmt.Println("vip收到请求")

}