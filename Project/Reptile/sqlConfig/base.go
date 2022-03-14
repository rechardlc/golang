package sqlConfig

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"strings"
)

const (
	userName = "root"
	password = ""
	ip       = "192.168.3.138"
	port     = "3306"
	dbName   = "reptile"
)

var Db *sqlx.DB

func init() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r, "catch:链接数据库异常~")
		}
	}()
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	Db, _ = sqlx.Open("mysql", path)
	Db.SetConnMaxLifetime(100)
	Db.SetMaxIdleConns(10)
}
