package main

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func get_config() {
	viper.SetConfigName("config") //把json文件换成yaml文件，只需要配置文件名 (不带后缀)即可
	viper.AddConfigPath(".")      //添加配置文件所在的路径
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("打开文件失败: %s\n", err)
		os.Exit(1)
	}

	urlValue := viper.Get("mysql.url")
	fmt.Println("mysql url:", urlValue)
	fmt.Printf("mysql url: %s\nmysql username: %s\nmysql password: %s", viper.Get("mysql.url"), viper.Get("mysql.username"), viper.GetString("mysql.password"))
}

func main() {
	get_config()
}
