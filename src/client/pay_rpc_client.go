package client

import (
	"fmt"
	"helper"
	"net/http"
)


//微信支付api
//https://pay.weixin.qq.com/wiki/doc/api/wxa/wxa_api.php?chapter=7_4&index=3


type PayClient struct {
	ModuleName string
	funcList map[string]funcHandler
}

func NewPayClient() *PayClient {

	p := PayClient {
		ModuleName:"pay",
	}
	p.funcList = make(map[string]funcHandler)



	return &p
}

//方法注册函数
func (p *PayClient) Register (method string, f funcHandler) {
	p.funcList[method] = f
}


func (p *PayClient) Dispatch(method string, resp http.ResponseWriter, req *http.Request) {

	method, err := helper.GetMethod(req)
	if err != nil {
		fmt.Println("invalid method", req.RequestURI)
		return
	}

	if f, find := p.funcList[method]; find {
		f(resp, req)
	} else  {
		fmt.Println("method is not register [%T]", method)
	}

}

//统一下单接口, https://pay.weixin.qq.com/wiki/doc/api/wxa/wxa_api.php?chapter=9_1&index=1
func (p *PayClient) unitfiledOrder () {
	//
	//appId := helper.AppId
	//mch_id := "asdasd"//商户id
	//
	//nonce_str := helper.RandomStr(16) //随机字符串
	//
	//sign := "dsajdaskjls" //签名
	//
	//body := "商品描述"
	//
	//out_trade_no := "out_trade_no" //商户订单号
	//
	//total_fee := 123 //考核
	//
	//spbill_create_ip := "ds" //机器ip
	//
	//trade_type := "JSAPI"
	//
	//notify_url := "asd" //通讯地址










}