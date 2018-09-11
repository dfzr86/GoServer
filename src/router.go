package main

import (
	"client"
	"fmt"
	"helper"
	"net/http"
)


type RouterMethod interface {
	Dispatch(method string, resp http.ResponseWriter, req *http.Request)
}

type Router struct {
	handler		map[string]RouterMethod  //  模块 : handler e.g :  user : userHandler
}

func NewRouter()  *Router {

	r := Router{}

	userClient := client.NewUserClient()
	articleClient := client.NewArticleClient()
	vipClient := client.NewVipClient()
	payClient := client.NewPayClient()
	commentClient := client.NewCommentClient()

	r.handler = make(map[string]RouterMethod)

	r.DispatchHandler("user", userClient)
	r.DispatchHandler("article", articleClient)
	r.DispatchHandler("vip", vipClient)
	r.DispatchHandler("pay", payClient)
	r.DispatchHandler("comment", commentClient)

	return &r
}

//模块派发
func (r *Router) DispatchHandler(module string, handler RouterMethod) {
	r.handler[module] = handler
}

func (t *Router) ServeHTTP (resp http.ResponseWriter, req *http.Request) {

	// TODO: 做签名校验/安全等等过滤, 然后分发
	fmt.Println("router --> ServeHTTP", req.RequestURI)

	t.ServeMyHTTP(resp, req)
}

//分发
func (r *Router) ServeMyHTTP (resp http.ResponseWriter, req *http.Request) {

	// 分发给对应的业务handler
	module, err := helper.GetModule(req)
	if err != nil {
		fmt.Println("getModule err", err)
		return
	}

	method, _ := helper.GetMethod(req)

	fmt.Println("module is ", module)

	if len(module) == 0 {
		fmt.Println("invalid module", req.RequestURI)
		return
	}
	handler, find := r.handler[module]
	if find {
		handler.Dispatch(method, resp, req)

	} else {
		fmt.Println("invalid handler", req.RequestURI)
		return
	}
}

func (t *Router) Start() {

	http.ListenAndServe("127.0.0.1:8080", t)
}