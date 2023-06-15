package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/Yoak3n/qbtNotification/util"
	"log"
	"net/http"
	"time"
)

var id int64
var status string
var name string
var token string
var host string
var group int64
var check bool

func init() {
	flag.StringVar(&token, "t", "", "access_token")
	flag.Int64Var(&id, "id", 0, "QQ号")
	flag.StringVar(&status, "s", "end", "状态")
	flag.StringVar(&name, "n", "", "下载完成的内容")
	flag.StringVar(&host, "host", "127.0.0.1:5700", "go-cqhttp的http地址及端口号")
	flag.Int64Var(&group, "group", 0, "QQ群号")
	flag.BoolVar(&check, "check", false, "是否检查文件名为hash值")
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
		if ok {
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
		for _, item := range res {
			if item.StatusCode == 200 {
				log.Println("成功发送消息")
				return
			} else {
				count++
				if count == 1 {
					go util.DebugNetwork()
				}
				if item.StatusCode == 502 {
					log.Println("服务器可能屏蔽了当前IP的网络请求，请检查当前的网络配置")
				}
				fmt.Printf("消息发送失败正在重试，已失败次数：%d\n失败原因：%s\n", count, item.Status)
				if count == 10 {
					log.Println("消息发送失败已达10次，放弃发送！")
					return
				}
				time.Sleep(time.Second * 5)
			}
		}
	}
}

func sendPrivate(msg string) *http.Response {
	l := fmt.Sprintf("http://%s/send_private_msg", host)
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
		log.Fatalln(err)
	}
	return res
}

func sendGroup(msg string) *http.Response {
	l := fmt.Sprintf("http://%s/send_group_msg", host)
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
		log.Fatalln(err)
	}
	return res
}

func send(msg string) (res []*http.Response) {
	if group != 0 {
		res = append(res, sendGroup(msg))
	}
	if id != 0 {
		res = append(res, sendPrivate(msg))
	}
	return res
}
