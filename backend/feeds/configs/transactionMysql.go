package configs

import (
	"github.com/jinzhu/gorm"
	"github.com/pranay999000/feeds/utils"
)

var (
	transactionDB *gorm.DB
)

func TransactionConnect() {
	writeMySQLHost, err := EnvMap("write_mysql_host")
	utils.FailOnError(err, "Unable to find write_mysql_host")

	writeMySQLPort, err := EnvMap("write_mysql_port")
	utils.FailOnError(err, "Unable to find write_mysql_port")

	transactiondb, err := gorm.Open("mysql", "root:masterpassword@tcp(" + writeMySQLHost + ":" + writeMySQLPort + ")/masterdatabase?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err.Error())
	}

	transactionDB = transactiondb
}

func GetTransactionDB() *gorm.DB {
	return transactionDB
}
