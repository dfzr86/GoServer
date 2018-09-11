package helper

import (
	"fmt"
	"net/http"
	"strings"
)

func GetModule(req *http.Request) (string, error) {

	list := strings.Split(req.RequestURI, "/")
	fmt.Println("------list ", list, "len = ", len(list))
	if len(list) < 1 {
		fmt.Println("invalid module", req.RequestURI)
		return "", fmt.Errorf("invalid module", req.RequestURI)
	}

	fmt.Println("------list element: ", list[0], list[1], list[2])
	return list[1], nil
}


func GetMethod(req *http.Request) (string, error) {

	a := strings.Split(req.RequestURI, "?")

	fmt.Println("<>", a)

	uri := a[0]
	list := strings.Split(uri, "/")
	fmt.Println("------list ", list, "len = ", len(list))
	if len(list) < 1 {
		fmt.Println("invalid module", req.RequestURI)
		return "", fmt.Errorf("invalid module", req.RequestURI)
	}

	var method string
	temp := list[2:]
	for _, part := range temp {
		method += part
		method += "/"
	}
	method = strings.TrimRight(method, "/")
	return method, nil
}