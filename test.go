package main

const tagName = "validate"

type employee struct {
	ID       int     `json:"id"`
	Name     string  `json:"名字" validate:"presence,min=2,max=40"`
	Age      int     `json:"年龄"`
	Desc     string  `json:"描述" back:"好看否"`
	weight   float64 `json:"weight" 单位:"kg"`
	Salary   float64 `json:"-"`
	Email    string  `validate:"email,required"`
	MateName string  `json:"mate_name,omitempty"`
}

func main() {
	a := "2019-11-28T14:19:37.13769Z"

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
