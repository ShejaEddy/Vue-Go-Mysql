package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

func Connect(
	host string,
	port string,
	user string,
	password string,
	database string,
) (db *xorm.Engine, err error) {
	db, err = xorm.NewEngine("mysql", user+":"+password+"@tcp("+host+":"+port+")/"+database+"?parseTime=true")
	if err != nil {
		return
	}
	db.SetConnMaxLifetime(1000)
	_, err = db.DBMetas()

	return
}
