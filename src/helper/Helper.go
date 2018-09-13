package helper

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"strings"
	"time"
)

const AppId = "wx48a9d80716d5e341"
const AppSecrety = "2d9202f4b54f9be364d96d27e99f4037"

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

//RandomStr 随机生成字符串
func RandomStr(length int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}


// LocalIP 获取机器的IP
func LocalIP() string {
	info, _ := net.InterfaceAddrs()
	for _, addr := range info {
		ipNet, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}
		if !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			return ipNet.IP.String()
		}
	}
	return ""
}

func MapStringToStruct(m map[string]string, i interface{}) error {
	bin, err := json.Marshal(m)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bin, i)
	if err != nil {
		return err
	}
	return nil
}