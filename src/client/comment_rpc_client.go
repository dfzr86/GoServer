package client

import (
	"fmt"
	"net/http"
)

type CommentClient struct {
}

func NewCommentClient() *CommentClient  {
	return nil
}



func (c *CommentClient) Dispatch(method string, resp http.ResponseWriter, req *http.Request) {
	fmt.Println("comment收到请求")

}