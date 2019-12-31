package main

import (
	"encoding/json"
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
	//x := 1
	//y := 2
	//tp1 := calc("B", x, y)
	//defer calc("A", x, tp1)
	//x = 3
	//tp2 := calc("D", x, y)
	//defer calc("C", x, tp2)
	//y = 4
	a := []string{"ads", "bb"}
	res, _ := json.Marshal(a)
	fmt.Println(string(res))
}
