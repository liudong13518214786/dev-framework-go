package main

import (
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
		}
	}()
	a := []string{"123"}
	fmt.Println(a[:len(a)-3])
}
