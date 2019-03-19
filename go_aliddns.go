package main

import (
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"strings"
)

var (
	DomainName,
	RR,
	Type,
	Value,
	accessKeyId,
	accessSecret string
	TTL int32
)

const ipFileName = "ip.txt"

// 判断ip文件是否存在
func CheckFileIsExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

// 读取ip地址
func ReaderIPAddress() string {
	ip, err := ioutil.ReadFile(ipFileName)
	if err != nil {
		fmt.Println("无法读取ip")
		return ""
	}
	return string(ip)
}

// 写入ip地址
func WriteIPAddress(ip string) bool {
	err := ioutil.WriteFile(ipFileName, []byte(ip), 0777)
	if err != nil {
		fmt.Println("写入ip错误")
		return false
	}
	return true
}

func init() {
	const CONFIG = "config"
	viper.New()
	viper.SetEnvPrefix(CONFIG)
	viper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.SetConfigName(CONFIG)
	viper.AddConfigPath("./")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(fmt.Errorf("Fatal error when reading %s config file:%s", CONFIG, err))
		os.Exit(1)
	}

	DomainName = viper.GetString("DomainName")
	RR = viper.GetString("RR")
	Type = viper.GetString("Type")
	accessKeyId = viper.GetString("accessKeyId")
	accessSecret = viper.GetString("accessSecret")
	TTL = viper.GetInt32("TTL")

	if CheckFileIsExist(ipFileName) {
		Value = ReaderIPAddress()
	} else {
		WriteIPAddress("127.0.0.1")
	}
	Value = ReaderIPAddress()
}

func main() {
	fmt.Println(DomainName)
	fmt.Println(RR)
	fmt.Println(Value)
	fmt.Println("this is go-aliddns project")
}
