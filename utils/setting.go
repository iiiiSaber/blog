package utils

import (
	"fmt"

	"gopkg.in/ini.v1"
)

var (
	AppMode string
	HttpPort string
	JwtKey string

	Db string
	DbPort string
	DbHost string
	DbUser string
	DbPassWord string
	DbName string

	AccessKey string
	SecretKey string
	Bucket string
	MinioServer string
)

func init() {
	file,err := ini.Load("config/config.ini")
	if err!=nil{
		fmt.Println("config cannot be read,please check it:",err)
	}
	LoadServer(file)
	LoadData(file)
	LoadMinio(file)
}
func LoadServer(file *ini.File){
	AppMode=file.Section("server").Key("AppMode").MustString("debug")
	HttpPort=file.Section("server").Key("HttpPort").MustString(":3131")
	JwtKey=file.Section("server").Key("JwtKey").MustString("fey42weqrt43rrduee2322")
}
func LoadData(file *ini.File){
	Db=file.Section("database").Key("Db").MustString("mysql")
	DbHost=file.Section("database").Key("DbHost").MustString("127.0.0.1")
	DbPort=file.Section("database").Key("DbPort").MustString("3306")
	DbUser=file.Section("database").Key("DbUser").MustString("my_blog")
	DbPassWord=file.Section("database").Key("DbPassWord").MustString("2xpBYrzNBnzGHdy3")
	DbName=file.Section("database").Key("DbName").MustString("my_blog")
}
func LoadMinio(file *ini.File){
	AccessKey = file.Section("minio").Key("AccessKey").String()
	SecretKey = file.Section("minio").Key("SecretKey").String()
	Bucket = file.Section("minio").Key("Bucket").String()
	MinioServer = file.Section("minio").Key("MinioServer").String()
}