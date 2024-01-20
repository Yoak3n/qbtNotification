package util

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func Retry(res []*http.Response, count int) bool {
	if len(res) == 0 {
		return false
	}
	for _, item := range res {
		if item.StatusCode == 200 {
			log.Println("成功发送消息")
			return true
		} else {
			count++
			if count == 1 {
				go debugNetwork()
			}
			if item.StatusCode == 502 {
				log.Println("服务器可能屏蔽了当前IP的网络请求，请检查当前的网络配置")
			}
			fmt.Printf("消息发送失败正在重试，已失败次数：%d\n失败原因：%s\n", count, item.Status)
			if count == 10 {
				log.Println("消息发送失败已达10次，放弃发送！")
				return true
			}
			time.Sleep(time.Second * 5)
		}
	}
	return true
}
