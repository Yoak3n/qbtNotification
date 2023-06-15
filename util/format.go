package util

import (
	"fmt"
)

func FormatMsg(flag bool, text *string) (msg string) {
	if flag {
		msg = fmt.Sprintf("⚡qBittorrent已开始下载：\n%s", *text)
	} else {
		msg = fmt.Sprintf("⭕qBittorrent已完成下载：\n%s", *text)
	}
	//msg = url.QueryEscape(msg)
	return
}
