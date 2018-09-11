package main

import (
	"fmt"
)

/*
1）用户信息相关表，支持前端获取openid和身份及记录相关信息
2）支持前端获取课程列表、每集点播信息
3）后端能增、删、改课程信息、每集信息
4）答疑相关信息的增删改
5）支付相关接口联调
6）每集点播或答疑下面的用户留言支持
7）相关统计信息，比如播放次数、顶、踩等
8）后端支持查看用户的播放、下单、购买会员等统计
 */

func main() {

	server := NewRouter()
	if server == nil {
		fmt.Println("new router fail")
		return
	}

	server.Start()
}



