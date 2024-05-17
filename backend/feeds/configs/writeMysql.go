package configs

import (
	"github.com/jinzhu/gorm"
	"github.com/pranay999000/feeds/utils"
)

var (
	writedb *gorm.DB
)

func WriteConnect() {
	writeMySQLHost, err := EnvMap("write_mysql_host")
	utils.FailOnError(err, "Unable to find write_mysql_host")

	writeMySQLPort, err := EnvMap("write_mysql_port")
	utils.FailOnError(err, "Unable to find write_mysql_port")

	writed, err := gorm.Open("mysql", "root:masterpassword@tcp(" + writeMySQLHost + ":" + writeMySQLPort + ")/masterdatabase?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err.Error())
	}

	writedb = writed
}

func GetWriteDB() *gorm.DB {
	return writedb
}