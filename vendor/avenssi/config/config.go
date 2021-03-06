package config

import (
	"encoding/json"
	"os"
)

type Configuration struct {
	LBAddr  string `json:"lb_addr"`
	OssAddr string `json:"oss_addr"`
}

var configuration Configuration

func init() {
	file, _ := os.Open("./bin/conf.json")
	defer file.Close()
	// 根据io.Reader创建Decoder 然后调用Decode()方法将JSON解析成对象
	decoder := json.NewDecoder(file)
	configuration = Configuration{}

	err := decoder.Decode(&configuration)
	if err != nil {
		panic(err)
	}
}

func GetLbAddr() string {
	return configuration.LBAddr
}

func GetOssAddr() string {
	return configuration.OssAddr
}
