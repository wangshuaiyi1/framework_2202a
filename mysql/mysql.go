package mysql

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 连接mysql数据库
func Client(handler func(db *gorm.DB) error) error {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("mysql.username"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetInt("mysql.port"),
		viper.GetString("mysql.dbname"),
	)
	cli, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	//获取数据库连接对象
	db, err := cli.DB()
	if err != nil {
		return err
	}
	//延迟结束
	defer db.Close()

	fmt.Println("mysql连接成功")
	return handler(cli)

}
