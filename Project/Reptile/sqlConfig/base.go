package sqlConfig

import (
	"example.com/m/v2/utils"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"strings"
)

var ip = "192.168.3.138"

const (
	userName = "root"
	password = ""
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
	if _ip, ok := utils.GetIPv4Addr().(string); ok {
		ip = _ip
	} else {
		panic(_ip)
	}
	var err error
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	Db, err = sqlx.Open("mysql", path)
	if err != nil {
		panic(err)
	}
	Db.SetConnMaxLifetime(100)
	Db.SetMaxIdleConns(10)
}
