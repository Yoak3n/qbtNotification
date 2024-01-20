package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/Yoak3n/qbtNotification/util"
	"log"
	"net/http"
)

var (
	id     int64
	status string
	name   string
	token  string
	host   string
	group  int64
	check  bool
	origin string
	kind   string
)

func init() {
	flag.StringVar(&token, "token", "", "access_token")
	flag.Int64Var(&id, "id", 0, "QQ号")
	flag.StringVar(&status, "status", "end", "状态")
	flag.StringVar(&name, "name", "", "下载完成的内容")
	flag.StringVar(&host, "host", "127.0.0.1:5700", "go-cqhttp的http地址及端口号")
	flag.Int64Var(&group, "group", 0, "QQ群号")
	flag.BoolVar(&check, "check", false, "是否检查文件名为hash值")
	flag.StringVar(&origin, "origin", "", "通知信息来源，如nas")
	flag.StringVar(&kind, "kind", "text", "通知信息类型，如text")
	flag.Parse()
	if group+id == 0 {
		panic("请指定通知对象：私聊的QQ号或群聊的群号")
	}
}

func main() {
	var tn string // 种子真正的名字
	if check && !util.CheckHash(name) {
		log.Println("需要进行解析")
		n, ok := util.ParseFileName(name)
		if ok && n != "" {
			tn = n
		} else {
			tn = name
			log.Println("仍然无法获得种子文件名")
		}
	} else {
		tn = name
	}
	msg := util.FormatMsg(status == "start", &tn)
	count := 0
	fmt.Printf("向%d %d发送：%s\n", id, group, tn)
	for {
		res := send(msg)
		if ok := util.Retry(res, count); ok {
			return
		}
		continue
	}
}

func sendPrivate(msg string) (*http.Response, error) {
	l := fmt.Sprintf("http://%s/send_private_msg?kind=%s&origin=%s", host, kind, origin)
	type Post struct {
		UserID  int64  `json:"user_id"`
		Message string `json:"message"`
	}
	post := &Post{
		UserID:  id,
		Message: msg,
	}
	data, _ := json.Marshal(post)
	req, _ := http.NewRequest("POST", l, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func sendGroup(msg string) (*http.Response, error) {
	l := fmt.Sprintf("http://%s/send_group_msg?kind=%s&origin=%s", host, kind, origin)
	type Post struct {
		GroupID int64  `json:"user_id"`
		Message string `json:"message"`
	}
	post := &Post{
		GroupID: group,
		Message: msg,
	}
	data, _ := json.Marshal(post)
	req, _ := http.NewRequest("POST", l, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func send(msg string) (res []*http.Response) {
	if group != 0 {
		rg, err := sendGroup(msg)
		if err != nil {
			log.Fatalln("发送群组消息出错：", err)
		} else {
			res = append(res, rg)
		}
	}
	if id != 0 {
		rg, err := sendPrivate(msg)
		if err != nil {
			log.Fatalln("发送私聊消息出错：", err)
		} else {
			res = append(res, rg)
		}
	}
	return res
}
