package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"unicode"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func isHan(r rune) bool {

	return unicode.Is(unicode.Han, r)

}
func main() {
	viper.SetConfigName("config") //获取配置文件
	viper.AddConfigPath(".")      //添加配置文件所在的路径
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("打开文件失败: %s\n", err)
		os.Exit(1)
	}
	//获取配置文件
	DbHost := viper.GetString("mysql.host")
	DbUsername := viper.GetString("mysql.username")
	DbPassword := viper.GetString("mysql.password")
	DbName := viper.GetString("mysql.dbname")
	DbCharset := viper.GetString("mysql.charset")
	Dbport := viper.GetString("mysql.port")
	mysqlpath := strings.Join([]string{DbUsername, ":", DbPassword, "@tcp(", DbHost, ":", Dbport, ")/", DbName, "?charset=", DbCharset}, "") //链接配置文件拼接

	ruler := viper.GetStringMap("config.ruler")

	data := []map[string]interface{}{}                          //初始化数据表
	db, err := gorm.Open(mysql.Open(mysqlpath), &gorm.Config{}) //链接数据库

	tableName := []map[string]interface{}{}
	db.Raw("show tables").Scan(&tableName)

	// db.Debug().Table("user").Find(&data) //查找数据

	for _, v := range tableName { //循环从数据库取出的表map
		for _, s := range v { //循环表map得到键值对
			sstring := s.(string) //转换数据库名称为字符串
			fmt.Println("正在查询表：", sstring)
			db.Debug().Table(sstring).Find(&data) //查找数据map
			for _, dataFor := range data {        //循环返回数据map
				for dateListName, dataForOne := range dataFor { //循环单条数据
					for rulerName, rulerFor := range ruler { //循环出整个规则列表
						// fmt.Println("type:", reflect.TypeOf(dataForOne))
						var dataForOneString string
						switch dataForOne.(type) {
						case string:
							// fmt.Println("is string", dataForOneType)
							dataForOneString = dataForOne.(string)
						case int:
							// fmt.Println("is int ", dataForOneType)
							dataForOneString = dataForOne.(string)
						case float64:
							// fmt.Println("is float64 ", dataForOneType)
							dataForOneString = dataForOne.(string)
						case int32:
							// fmt.Println("is int32 ", dataForOneType)
							dataForOneString = string(dataForOne.(int32))
						case int64:
							// fmt.Println("is int64 ", dataForOneType)

							dataForOneString = string(dataForOne.(int64))
						}

						// fmt.Println(rulerName, rulerFor.(string), dataForOneString, dateListName)
						matchDigit, _ := regexp.MatchString(rulerFor.(string), dataForOneString)
						if matchDigit == true {
							fmt.Println(matchDigit)
							fmt.Println(rulerName, sstring, dateListName)
						}
					}
					// fmt.Printf("%+v\n", dataForOne)
				}
			}

		}
	}
}
