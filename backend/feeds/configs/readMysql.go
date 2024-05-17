package configs

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/pranay999000/feeds/utils"
)

var (
	db *gorm.DB
)

func ReadConnect() {
	readMySQLHost, err := EnvMap("read_mysql_host")
	utils.FailOnError(err, "Unable to find read_mysql_host")

	readMySQLPort, err := EnvMap("read_mysql_port")
	utils.FailOnError(err, "Unable to find read_mysql_port")

	d, err := gorm.Open("mysql", "root:masterpassword@tcp(" + readMySQLHost + ":" + readMySQLPort + ")/masterdatabase?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		utils.FailOnError(err, "Unable to connect to read mysql")
	}
	db = d
}

func GetReadDB() *gorm.DB {
	return db
}