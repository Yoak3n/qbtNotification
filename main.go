package main

import (
	"flag"
	"fmt"
	"github.com/Yoak3n/qbtNotification/utils"
	"log"
	"net/http"
	"time"
)

var id string
var status string
var name string
var token string
var host string
var group string

func init() {
	flag.StringVar(&token, "t", "", "access_token")
	flag.StringVar(&id, "id", "", "QQ号")
	flag.StringVar(&status, "s", "end", "状态")
	flag.StringVar(&name, "n", "", "下载完成的内容")
	flag.StringVar(&host, "host", "127.0.0.1:5700", "go-cqhttp的http地址及端口号")
	flag.StringVar(&group, "group", "", "QQ群号")
	flag.Parse()
	fmt.Printf("向%s %s发送：%s\n", id, group, name)
	if group+id == "" {
		panic("请指定通知对象：私聊的QQ号或群聊的群号")
	}
}

func main() {
	msg := utils.FormatMsg(status == "start", &name)

	count := 0
	for {
		res := send(msg)
		for _, item := range res {
			if item.StatusCode == 200 {
				log.Println("成功发送消息")
				return
			} else {
				count++
				if count == 1 {
					go utils.DebugNetwork()
				}
				if item.StatusCode == 502 {
					fmt.Println("服务器可能屏蔽了当前IP的网络请求，请检查当前的网络配置")
				}
				fmt.Printf("消息发送失败正在重试，已失败次数：%d\n失败原因：%s\n", count, item.Status)
				if count == 10 {
					fmt.Println("消息发送失败已达10次，放弃发送！")
					return
				}
				time.Sleep(time.Second)
			}

		}

	}
}

func sendPrivate(msg string) *http.Response {
	l := fmt.Sprintf("http://%s/send_private_msg?access_token=%s&user_id=%s&message=%s", host, token, id, msg)
	res, err := http.Get(l)
	if err != nil {
		log.Fatal(err)
	}
	return res
}

func sendGroup(msg string) *http.Response {
	l := fmt.Sprintf("http://%s/send_group_msg?access_token=%s&group_id=%s&s&message=%s", host, token, group, msg)
	res, err := http.Get(l)
	if err != nil {
		log.Fatal(err)
	}
	return res
}

func send(msg string) (res []*http.Response) {
	if group != "" {
		res = append(res, sendGroup(msg))
	}
	if id != "" {
		res = append(res, sendPrivate(msg))
	}
	return res
}
