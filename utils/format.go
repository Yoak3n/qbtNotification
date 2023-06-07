package utils

import (
	"fmt"
	"net/url"
)

func FormatMsg(flag bool, text *string) (msg string) {
	if flag {
		msg = fmt.Sprintf("qBittorrent已开始下载：\n《%s》", url.QueryEscape(*text))
	} else {
		msg = fmt.Sprintf("qBittorrent已完成下载：\n《%s》", url.QueryEscape(*text))
	}
	return

}
