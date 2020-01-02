package main

import (
	"fmt"
)

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

type Users struct {
	Name string
}

func main() {
	a := []string{"123", "123", "2"}
	fmt.Println(a[:2])
}
func RepeatArr(tar []string) []string {
	result := make([]string, 0, len(tar))
	tmp := map[string]struct{}{}
	for _, item := range tar {
		if _, ok := tmp[item]; !ok {
			tmp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}
