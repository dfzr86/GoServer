package main

import "client"

type RequestHandler struct {
	userClient 			*client.UserClient
	articleClient 		*client.ArticleClient
	vipClient 			*client.VipClient
	payClient 			*client.PayClient
	commentClient 		*client.CommentClient
}

func NewRequestHandler(
	userClient *client.UserClient,
	articleClient *client.ArticleClient,
	vipClient *client.VipClient,
	payClient *client.PayClient,
	commentClient *client.CommentClient) *RequestHandler  {

		return &RequestHandler{
			userClient:userClient,
			articleClient:articleClient,
			vipClient:vipClient,
			payClient: payClient,
			commentClient:commentClient,
		}
}







