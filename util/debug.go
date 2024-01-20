package util

import (
	"fmt"
	"io"
	"net/http"
)

func checkPublicIP() (ip string, err error) {
	res, err := http.Get("http://myexternalip.com/raw")
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	content, _ := io.ReadAll(res.Body)
	ip = string(content)
	return
}

func debugNetwork() {
	ip, err := checkPublicIP()
	if err != nil {
		fmt.Println("网络请求错误")
	}
	fmt.Printf("您当前IP为%s\n", ip)
}
