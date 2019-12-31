package main

import (
	"dev-framework-go/pkg/util"
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
	a := util.GetNowTimeV2()
	fmt.Println(a)
}
