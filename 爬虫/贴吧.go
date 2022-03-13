package main

import (
	"encoding/json"
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
	//if err := ioutil.WriteFile("test.html", pageBytes, 0666); err != nil {
	//	panic(err) // 通过ioutil.WriteFile将数据写入文件中
	//}
	//pageStr := string(pageBytes)
	pageStr := strings.Split(string(pageBytes), "j_p_postlist")
	pageStr = strings.Split(pageStr[1], "right_bright")
	pageStr = strings.Split(pageStr[0], "l_post_bright")
	dealStrings(pageStr)
	//re := regexp.MustCompile(reQQEmail)              // 解析并返回一个正则表达式，全局正则表达式安全初始化
	//results := re.FindAllStringSubmatch(pageStr, -1) // n小于0，返回所有匹配想，n大于0，查找前n项
	//qq := make([]Person, 0)
	//for _, result := range results {
	//	if isInSlice(qq, result[1]) {
	//		qq = append(qq, Person{
	//			Number: result[1],
	//			Email:  result[0],
	//		})
	//	}
	//}
	//bytes, _ := json.Marshal(qq)
	//if err := ioutil.WriteFile("test.json", bytes, 0); err != nil {
	//	return
	//}
	//fmt.Println(string(bytes))
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

func main() {
	getEmail()
}
