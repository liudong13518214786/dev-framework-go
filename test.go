package main

import (
	"os"
	"os/signal"
)

func main() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Kill)

	//	zhangsan := employee{
	//		ID:       1,
	//		Name:     "张三",
	//		Age:      18,
	//		Desc:     "秀色可餐",
	//		weight:   48.0,
	//		Salary:   12.0,
	//		MateName: "Prince",
	//	}
	//
	//	fmt.Println(zhangsan)
	//	re, _ := json.Marshal(zhangsan)
	//	fmt.Println(string(re))
	//
	//	t := reflect.ValueOf(zhangsan)
	//	//c :=reflect.Ptr
	//	//fmt.Println("Type: ",t.Name())
	//	fmt.Println("Kind: ", t.Kind())
	//	fmt.Println("Kind: ", t.NumField())
	//	//fmt.Println("Kind: ",t.Kind())
	//	//fmt.Println(t.NumField())
	//	for i := 0; i < t.NumField(); i++ {
	//		//field := t.Field(i)
	//		tag := t.Type().Field(i)
	//		fmt.Println(tag.Tag.Get("json"))
	//		//fmt.Printf("%d. %v(%v), tag:'%v'\n", i+1, field.Name, field.Type.Name(), tag)
	//	}
	//}
}
