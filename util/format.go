package util

import (
	"fmt"
)

func FormatMsg(flag bool, text *string) (msg string) {
	if flag {
		msg = fmt.Sprintf("\t\tqBittorrent已开始下载：\n%s", *text)
	} else {
		msg = fmt.Sprintf("\t\tqBittorrent已完成下载：\n%s", *text)
	}
	//msg = url.QueryEscape(msg)
	return
}
