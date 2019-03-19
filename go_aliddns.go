package main

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
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

// 文件相关操作
// ------------------------------------------------
const ipFileName = "ip.txt"

// 判断文件是否存在
func CheckFileIsExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

// 读取文件
func ReaderFile(filename string) string {
	ip, err := ioutil.ReadFile(filename)
	if err != nil {
		return ""
	}
	return string(ip)
}

// 写入文件
func WriteFile(filename, str string) bool {
	err := ioutil.WriteFile(ipFileName, []byte(str), 0777)
	if err != nil {
		return false
	}
	return true
}

// ------------------------------------------------
// ip相关
// ------------------------------------------------
// 获取公网ip地址
func GetIPAddressByHttp() string {
	resp, err := http.Get("http://members.3322.org/dyndns/getip")
	if err != nil {
		fmt.Println("获取公网ip失败!")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("获取公网ip失败!")
	}
	ip := body[:len(body)-1] // 返回的数据最后有一个空格
	return string(ip)
}

// 获取ip
func GetIP() (ip string) {
	ip = GetIPAddressByHttp()

	if CheckFileIsExist(ipFileName) { // ip地址文件不存在
		WriteFile(ipFileName, ip)
	} else {
		if tmpIp := ReaderFile(ipFileName); ip != tmpIp { // 如果存在且地址变化了 重新写入
			WriteFile(ipFileName, ip)
			fmt.Printf("上次解析的ip:%s", tmpIp)
			fmt.Printf("本次解析的ip:%s", ip)
		}
	}
	return
}

// ------------------------------------------------

// 阿里云相关操作

// ------------------------------------------------

// 检测解析记录是否生效
func ChecCheckDomainRecordk() bool {
	client, err := alidns.NewClientWithAccessKey("cn-hangzhou", accessKeyId, accessSecret)

	request := alidns.CreateCheckDomainRecordRequest()
	request.DomainName = DomainName
	request.Value = Value
	request.Type = Type
	request.RR = RR

	response, err := client.CheckDomainRecord(request)
	if err != nil {
		fmt.Print("检测解析记录失败")
		return false
	}
	fmt.Printf("response is %#v\n", response)
	fmt.Println(response.IsExist)
	if response.IsExist {
		fmt.Println("解析记录已生效")
		return true
	} else {
		fmt.Println("解析记录未生效")
		return false
	}
}

// ------------------------------------------------

// 初始化相关操作
// ------------------------------------------------

func initViper() {
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
}

func init() {
	initViper()

	DomainName = viper.GetString("DomainName")
	RR = viper.GetString("RR")
	Type = viper.GetString("Type")
	accessKeyId = viper.GetString("accessKeyId")
	accessSecret = viper.GetString("accessSecret")
	TTL = viper.GetInt32("TTL")
	Value = GetIP()
}

// ------------------------------------------------

// main
func main() {
	ChecCheckDomainRecordk()
	fmt.Println(DomainName)
	fmt.Println(RR)
	fmt.Println("this is go-aliddns project")
}
