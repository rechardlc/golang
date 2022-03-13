package main

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
			fmt.Println(r, "catch~error~~~")
		}
	}()
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	Db, _ = sqlx.Open("mysql", path)
	Db.SetConnMaxLifetime(100)
	Db.SetMaxIdleConns(10)
	fmt.Println("connect mysql success~")
	//createTieBaTable()
}

func createTieBaTable() {
	fmt.Println(Db, "Db~~~~result")
	schema := `CREATE TABLE place (
    id int primary key auto_increment,
    country varchar(50),
    city varchar(50) NULL default '',
    telcode int);`
	res, err := Db.Exec(schema)
	fmt.Println("res:", res, "err:", err)
}

func main() {
	//createTieBaTable()
	//fmt.Println("main")
	//insertIntoPlace("中国", "广州", 998)
	//updatePlaceRow(100, "武汉", 2)
	query()
}
func insertIntoPlace(country, city string, telCode int) {
	defer Db.Close()
	countryCitySql := `insert into place (country, city, telcode) values (?,?,?)`
	result, err := Db.Exec(countryCitySql, country, city, telCode)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	id, _ := result.LastInsertId()
	fmt.Println(id)
}
func updatePlaceRow(telcode int, city string, id int) {
	defer Db.Close()
	sql := `update place set telcode=?, city=? where id =?`
	res, err := Db.Exec(sql, telcode, city, id)
	if err != nil {
		panic(res)
	}
	fmt.Println(res)
}
func query() {
	defer Db.Close()
	p := struct{}{}
	err := Db.Get(&p, "select * from place limit 10")
	fmt.Println(p, err)
}
