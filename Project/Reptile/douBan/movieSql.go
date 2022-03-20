package douBan

import (
	"database/sql/driver"
	"encoding/json"
	"example.com/m/v2/httpServer"
	"example.com/m/v2/sqlConfig"
	"fmt"
	"net/http"
)

func init() {
	createMovieTable()
}

// Value
//  @Description: 实现Movie结构上的Value方法，sqlx会默认查询方法
//  @receiver m
//  @return driver.Value
//  @return error
//
func (m Movie) Value() (driver.Value, error) {
	return []interface{}{
		m.CanPlayable, m.Order, m.Url, m.SubTitle, m.Score, m.ValuationNum,
		m.Countries, m.Year, m.Types, m.Title, m.Personnel,
	}, nil
}
func createMovieTable() {
	var err error
	schema := `create table if not exists movie (
    id int primary key auto_increment,
    m_order int null,
    title varchar(255) default '',
    url varchar(255) default '',
    sub_title varchar(9999) default '',
    score int not null default 0,
  	valuation_num int not null default 0,
  	m_describe varchar(9999) default '',
  	can_playable boolean not null default false,
  	m_year int default 1111,
  	personnel varchar(255) default '',
  	countries varchar(255) default '',
  	types varchar(255) default ''
);`
	_, err = sqlConfig.Db.Exec(schema)
	if err != nil {
		panic(err)
	}
}
func insertIntoMovie(movies []Movie) {
	fmt.Println("Movie:length", len(movies))
	sql := `insert into movie (m_order, title, url, sub_title, score,
	             valuation_num, m_describe, m_year, personnel, countries, types)
	             values (:m_order, :title, :url, :sub_title, :score,
	             :valuation_num, :m_describe, :m_year, :personnel, :countries, :types);`
	_, err := sqlConfig.Db.NamedExec(sql, movies)
	if err != nil {
		panic(err)
	}
}
func DoubanMovie(writer http.ResponseWriter, request *http.Request) {
	defer sqlConfig.Db.Close()
	var m []Movie
	sqlConfig.Db.Select(&m, `select * from movie;`)
	var r httpServer.RespStruct
	r.Status = 200
	r.Data.Data = m
	r.ErrMsg = ""
	r.Data.Page.Page = 1
	r.Data.Page.TotalCount = 250
	r.Data.Page.Count = 20
	marshal, _ := json.Marshal(r)
	writer.Write(marshal)
}
