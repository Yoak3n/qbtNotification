package util

import "github.com/tidwall/gjson"

func parseJson(j string) string {
	result := gjson.Parse(j)
	if statusCode := result.Get("success").Int(); statusCode != 1 {
		return ""
	}
	return result.Get("info.name").String()
}
