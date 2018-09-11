package client

import (
	"encoding/json"
	"fmt"
	"helper"
	"net/http"
	"user"
	_ "user"
)

type funcHandler func(w http.ResponseWriter, r *http.Request)

type UserClient struct {

	funcList	map[string]funcHandler
}

func NewUserClient() *UserClient {

	uc := UserClient{}
	uc.funcList = make(map[string]funcHandler)

	//openId
	uc.Register("getOpenId", uc.getOpenId)
	//用户信息
	uc.Register("userInfo", uc.userInfo)
	return &uc
}

//方法注册函数
func (u *UserClient) Register (method string, f funcHandler) {
	u.funcList[method] = f
}

func (u *UserClient) Dispatch(method string, resp http.ResponseWriter, req *http.Request) {

	fmt.Println("UserClient收到请求", u)

	method, err := helper.GetMethod(req)
	if err != nil {
		fmt.Println("invalid method", req.RequestURI)
		return
	}

	if f, find := u.funcList[method]; find {
		f(resp, req)
	} else  {
		fmt.Println("method is not register [%T]", method)
	}
}

//****************** Func IMP *******************

func (u *UserClient) getOpenId (resp http.ResponseWriter, req *http.Request)  {

	vars := req.URL.Query()

	code := vars["code"][0]

	openId, err := user.GetOpenId(code)
	if err != nil {
		fmt.Println(err)
		return;
	}

	respJson, err:= json.Marshal(openId)
	fmt.Println("resp:", respJson);
	if err != nil {
		fmt.Println("序列化openId失败")
	}

	resp.Write(respJson)
}


func (u *UserClient) userInfo (resp http.ResponseWriter, req *http.Request)  {

	vars := req.URL.Query()

	code := vars["code"][0]

	user, err := user.GetUserInfo(code)
	if err != nil {
		fmt.Println(err)
		return;
	}

	respJson, err:= json.MarshalIndent(&user, "", "\t\t")
	if err != nil {
		fmt.Println("序列化User对象失败")
	}

	resp.Write(respJson);
}