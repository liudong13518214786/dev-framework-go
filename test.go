package main

import (
	"fmt"
	"github.com/wonderivan/logger"
)

func main() {

	code_log := logger.NewLogger(1)
	//logger.Register("app_log", code_log)
	err := logger.SetLogger(`{"Console":{"LogLevel": 1, "level": "DEBUG"}}`)
	fmt.Println(err)
	logger.Debug("123123")
}
