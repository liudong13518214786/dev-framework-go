package main

import "encoding/base64"

func main() {
	a := "1232131sdfafs"
	c := base64.StdEncoding.EncodeToString([]byte(a))
	println(c)
	e, _ := base64.StdEncoding.DecodeString(c)

	println(string(e))
}
