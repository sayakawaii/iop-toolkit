/*
# ------------------------------------------------------------
# -- redis.go
# --
# -- Huang Minghe
# -- 2022-8-16
# ------------------------------------------------------------
*/

package global

import (
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	MysqlDB *gorm.DB
)

type MysqlOptions struct {
	User     string
	Addr     string
	Password string
	DB       string
}

func MysqlNewClient(opt *MysqlOptions) *gorm.DB {
	//DSN
	DSN := "" + opt.User + ":" + opt.Password + "@tcp(" + opt.Addr + ")/" + opt.DB + "?charset=utf8&parseTime=True&loc=Local"
	//指定驱动
	const DRIVER = "mysql"
	db, err := gorm.Open(DRIVER, DSN)
	if err != nil {
		panic(err)
	}
	return db
}

func InitMysql() error {
	MysqlDB = MysqlNewClient(&MysqlOptions{
		User:     AppConf.Mysql.User,
		Addr:     AppConf.Mysql.Addr,
		Password: AppConf.Mysql.Password,
		DB:       AppConf.Mysql.DefaultDB,
	})
	return nil
}
