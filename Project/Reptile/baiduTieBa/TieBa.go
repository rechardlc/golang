package main

import (
	"encoding/json"
	"example.com/m/v2/sqlConfig"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"sync"
)

var wg sync.WaitGroup

const (
	reQQEmail    = `(\d+)@qq.com`
	rePersonName = `target="_blank">(.*)</a>`
	reLevel      = `title="本吧头衔(\d+)级，经验值(\d+)，点击进入等级头衔说明页">`
)

type Level struct {
	Value      string `json:"level"`
	EmpiricVal string `json:"empiricVal"`
}

type Person struct {
	Number interface{} `json:"number"` // 可以验证结构体~ https://blog.csdn.net/netdxy/article/details/78528211
	Email  string      `json:"email"`
	Name   string      `json:"name"`
	Level  Level       `json:"level"`
}

type Tiba struct {
	Number string `db:"number"`
	Email  string `db:"email"`
	Name   string `db:"name"`
}
type Tiba_Level struct {
	Value      string `db:"value"`
	EmpiricVal string `db:"empiric_val"`
	Tid        string `db:"t_id"`
}

func getEmail() {
	if err := recover(); err != nil {
		fmt.Println(err)
	}
	resp, err := http.Get("https://tieba.baidu.com/p/6051076813?red_tag=1573533731")
	HandleError(err, "http.Get url")
	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			panic(err)
		}
	}(resp.Body)
	pageBytes, err := ioutil.ReadAll(resp.Body)
	HandleError(err, "ioutil.ReadAll")
	pageStr := strings.Split(string(pageBytes), "j_p_postlist")
	pageStr = strings.Split(pageStr[1], "right_bright")
	pageStr = strings.Split(pageStr[0], "l_post_bright")
	dealStrings(pageStr)
}

func dealStrings(divs []string) {
	var ch = make(chan Person, 1024)
	var s = make([]Person, 0)
	for _, div := range divs {
		wg.Add(1)
		div := div
		go func() {
			pNameRes := matchResult(rePersonName, div)
			pLeRes := matchResult(reLevel, div)
			pQQRes := matchResult(reQQEmail, div)
			p := &Person{}
			if len(pNameRes) > 0 {
				p.Name = pNameRes[0][1]
				if len(pLeRes) > 0 {
					p.Level.Value = pLeRes[0][1]
					p.Level.EmpiricVal = pLeRes[0][2]
				}
				if len(pQQRes) > 0 {
					p.Email = pQQRes[0][0]
					p.Number = pQQRes[0][1]
				}
				ch <- *p
			}
			wg.Done()
		}()
	}
	wg.Wait()
	close(ch)
	for {
		v, ok := <-ch
		if !ok {
			break
		}
		s = append(s, v)
	}
	s = sliceRemoveRepeat(s)
	id, err := intoRecordToTieBa("richard", "1119641305", "1119641305@qq.com")
	if err != nil || id == -1 {
		return
	}
	intoRecordToTiebaLevel(id, "10", "100")
	jsonRes, _ := json.Marshal(s)
	ioutil.WriteFile("test.json", jsonRes, 0666)
}

func matchResult(regStr, div string) [][]string {
	re := regexp.MustCompile(regStr)
	result := re.FindAllStringSubmatch(div, -1)
	return result
}

// 切片去重
func sliceRemoveRepeat(sic []Person) []Person {
	var m = make(map[string]string, len(sic))
	var s = make([]Person, 0)
	for _, v := range sic {
		if _, ok := m[v.Name]; !ok {
			s = append(s, v)
			m[v.Name] = v.Name
		}
	}
	return s
}

func HandleError(err error, why string) {
	if err != nil {
		fmt.Println(why, err)
	}
}

//  createTableTieBa
//  @Description:
//
func createTableTieBa() {
	schema := `CREATE TABLE tieba (
		id int primary key auto_increment,
		name varchar(50),
		email varchar(50),
		number varchar(50)
	);`
	if _, err := sqlConfig.Db.Exec(schema); err != nil {
		panic(err)
		return
	}
	schema = `create table tieba_level(
		id int primary key auto_increment,
		value varchar(50),
		empiricVal varchar(100)   
	)`
	if _, err := sqlConfig.Db.Exec(schema); err != nil {
		panic(err)
		return
	}
}

//
//  intoRecordToTieBa
//  @author:
//  @Description:
//  @param name
//  @param number
//  @param email
//  @return interface{}
//  @return error
//
func intoRecordToTieBa(name, number, email string) (int, error) {
	sql := `insert into tieba (name, number, email) value(?,?,?)`
	exec, err := sqlConfig.Db.Exec(sql, name, number, email)
	if err != nil {
		return -1, err
	} else {
		var id, _ = exec.LastInsertId()
		return int(id), nil
	}
}
func intoRecordToTiebaLevel(id int, value, EmpiricVal string) error {
	sql := `insert into tieba_level (value, empiric_val, t_id) value (?,?,?)`
	_, err := sqlConfig.Db.Exec(sql, value, EmpiricVal, id)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err, "<---main catch~err")
		}
	}()
	//createTableTieBa()
	getEmail()
}
