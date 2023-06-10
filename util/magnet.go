package util

import (
	"bytes"
	"encoding/json"
	"github.com/kyoto44/rain/magnet"
	"io"
	"log"
	"net/http"
)

var magPrefix string

func init() {
	magPrefix = "magnet:?xt=urn:btih:"
}

func CheckHash(name string) bool {
	log.Println("检查文件名")
	uri := magPrefix + name
	_, err := magnet.New(uri)
	if err != nil {
		// true 即为已经是文件名，不需要进一步获取文件信息
		log.Println("解析失败,已经是文件名:", err)
		return true
	}
	//mag := m.String()
	//if uri == mag {
	//	log.Println("名字是hash值")
	//	return false
	//}
	return false
}

func ParseFileName(name string) (string, bool) {
	uri := magPrefix + name
	// magnet查询接口
	url := "https://api.magnet-vip.com/api2/magnetinfo"
	// 构建json请求
	type data struct {
		Url string `json:"url"`
	}
	d := &data{Url: uri}
	j, err := json.Marshal(d)
	if err != nil {
		return "", false
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(j))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Println("请求创建失败：", err)
		return "", false
	}
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Println("请求发送失败", err)
		return "", false
	}
	b, _ := io.ReadAll(res.Body)
	defer res.Body.Close()
	return parseJson(string(b)), true
}
