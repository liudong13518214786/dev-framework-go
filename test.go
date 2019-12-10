package main

import (
	"dev-framework-go/conf"
	"dev-framework-go/pkg/util"
	"fmt"
	"runtime/debug"
	"strings"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			DebugStack := ""
			for _, v := range strings.Split(string(debug.Stack()), "\n") {
				DebugStack += v + "<br>"
			}
			fmt.Println(DebugStack)
			emailInfo := conf.EmailDetail
			body := strings.ReplaceAll(emailInfo, "[error_info]", fmt.Sprintf("%s", r))
			body = strings.ReplaceAll(body, "[request_time]", util.GetNowTime())
			body = strings.ReplaceAll(body, "[request_url]", "")
			body = strings.ReplaceAll(body, "[request_ua]", "")
			body = strings.ReplaceAll(body, "[error_debug]", DebugStack)

		}
	}()
	a := []string{"123"}

	fmt.Println(a[:len(a)-3])
	//fmt.Println(b)
}
