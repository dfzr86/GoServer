package client

import (
	"article"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type ArticleClient struct {
	funcList map[string]funcHandler
}

func NewArticleClient() *ArticleClient {

	a := ArticleClient{}
	a.funcList = make(map[string]funcHandler)

	a.Register("article", a.Get)
	a.Register("articleList", a.GetList)
	a.Register("deleteArticle", a.Delete)
	return &a
}

//方法注册函数
func (a *ArticleClient) Register(method string, f funcHandler) {
	a.funcList[method] = f
}

func (a *ArticleClient) Dispatch(method string, resp http.ResponseWriter, req *http.Request) {
	if f, find := a.funcList[method]; find {
		f(resp, req)
	} else {
		fmt.Println("invalid method", req.RequestURI)
	}
}

//********************* Func IMP *********************//
func (a *ArticleClient) Get(w http.ResponseWriter, r *http.Request) {

		vars := r.URL.Query()

		id := vars["id"][0]

		resp, err := article.Get(id)
		if err != nil {
			fmt.Fprintln(w, "", err.Error())
			return
		}
		respJson, err := json.MarshalIndent(&resp, "", "\t\t")
		if err != nil {
			fmt.Println("序列化Article对象失败", err.Error())

		}

		w.Write(respJson)
}

func (a *ArticleClient) GetList(w http.ResponseWriter, r *http.Request) {

	vars := r.URL.Query()

	size := vars["size"][0]
	cursor := vars["cursor"][0]

	sizeInt, _ := strconv.Atoi(size)

	fmt.Println("size = ", size, "cursor = ", cursor)

	articleList, err := article.GetArticleList(sizeInt, cursor)
	if err != nil {
		fmt.Fprintln(w, "getArticleListErr", err)
		return
	}

	resp, _ := json.Marshal(articleList)
	w.Write(resp)
}


func (a *ArticleClient) Delete(w http.ResponseWriter, r *http.Request) {
	vars := r.URL.Query()
	id := vars["id"][0]

	err := article.Delete(id)
	if err != nil {
		fmt.Println("删除文章失败", err.Error())
	}

	resp, _ := json.Marshal("ok")

	w.Write(resp)
}