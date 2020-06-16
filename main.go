package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configuration struct {
	LBAddr  string `json:"lb_addr"`
	OssAddr string `json:"oss_addr"`
}

var configuration *Configuration

func main() {
	file, _ := os.Open("./bin/conf.json")
	NAME := file.Name()
	fmt.Println(NAME)
	defer file.Close()
	// 根据io.Reader创建Decoder 然后调用Decode()方法将JSON解析成对象
	decoder := json.NewDecoder(file)
	configuration = &Configuration{}

	err := decoder.Decode(configuration)
	if err != nil {
		panic(err)
	}
	fmt.Println(configuration)
}
