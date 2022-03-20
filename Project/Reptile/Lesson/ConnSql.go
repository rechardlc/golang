package main

import (
	"encoding/json"
	"example.com/m/v2/utils"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"strings"
)

var ip = "192.168.3.138" // 默认地址
const (
	userName = "root"
	password = ""
	port     = "3306"
	dbName   = "reptile"
)

var Db *sqlx.DB

type iface interface{}

type Place struct {
	Id      int    `db:"id"`
	City    string `db:"city"`
	Telcode int    `db:"telcode"`
	Country string `db:"country"`
}

func init() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r, "catch~error~~~")
		}
	}()
	if _ip, ok := utils.GetIPv4Addr().(string); ok {
		ip = _ip
	}
	//net.IPv4()
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
	Db, _ = sqlx.Open("mysql", path)
	Db.SetConnMaxLifetime(100)
	Db.SetMaxIdleConns(10)
	fmt.Println("connect mysql success~")
	createTieBaTable()
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err, "catch~main~中的内容")
		}
	}()
	createTieBaTable()
	selectNameQuery()
	query()
	//fmt.Println("main")
	//insertIntoPlace("中国", "广州", 998)
	//updatePlaceRow(100, "武汉", 2)
	//query()
	//TestQueryx_Rowx()
	//batchInertInto()
}

func batchInertInto() {
	res, err := Db.NamedExec(
		`insert into place (country, city, telcode) values (:country, :city, :telcode)`,
		[]Place{{Telcode: 1, City: "南京", Country: "中国"}, {Telcode: 199, City: "伦敦", Country: "英国"}},
	)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}

// 在库中提供最常用的就是NamedQuery和NamedExec函数，一个是执行对查询参数命名并绑定，另一个则是对 CUD 操作的查询参数名的绑定：
//
//  selectNameQuery
//  @Description: 绑定查询
//  type Place struct {
//		Id      int    `db:"id"`
//		City    string `db:"city"`
//		Telcode int    `db:"telcode"`
//		Country string `db:"country"`
// }
//
func selectNameQuery() {
	sqlStr := `select id, city, telcode, country from place where id = :id`
	rows, err := Db.NamedQuery(sqlStr, map[string]interface{}{
		"id": 1,
	})
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var p Place
		if err = rows.StructScan(&p); err != nil {
			fmt.Println("struct failed:", err)
		}
		fmt.Println(p)
	}
}

func createTieBaTable() {
	fmt.Println(Db, "Db~~~~result")
	schema := `CREATE TABLE if not exists place (
    id int primary key auto_increment,
    country varchar(50),
    city varchar(50) NULL default '',
    telcode int);`
	res, err := Db.Exec(schema)
	fmt.Println("res:", res, "err:", err)
}

func TestQueryx_Rowx() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	defer Db.Close()
	s := make([]interface{}, 0)
	if rows, err := Db.Queryx("select * from place"); err != nil {
		panic(err)
	} else {
		for rows.Next() {
			//var (
			//	country, city string
			//	telcode       int
			//)
			//err = rows.Scan(&country, &city, &telcode)
			err = rows.StructScan(&Place{})
			v, _ := rows.SliceScan()
			s = append(s, v[0])
			var bb, _ = json.Marshal(s)
			fmt.Println(string(bb))
			//err, _ = rows.SliceScan(s)
		}
	}
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
	defer func(Db *sqlx.DB) {
		if err := Db.Close(); err != nil {
			panic(err)
		}
	}(Db)
	var p Place
	var total int
	if err := Db.Get(&total, "select count(*) from place"); err != nil {
		return
	}
	if err := Db.Get(&p, "select * from place where id=?", 4); err != nil {
		fmt.Println(err, "err~~")
		return
	}
	var pes []Place
	if err := Db.Select(&pes, "select * from place;"); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("total:", total, "\n", "p:", p, "\n", "pes:", pes)
}
