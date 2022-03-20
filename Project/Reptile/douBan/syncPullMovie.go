package douBan

// 同步方法~在非协程下的结果：1.832252501s
import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

var ws sync.WaitGroup

func (a Axios) responseResult() (body string, err error) {
	request, err := http.NewRequest(a.Method, a.Url, nil)
	for k, v := range a.Header {
		request.Header.Set(k, v)
	}
	if err != nil {
		return "", err
	}
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		if err = Body.Close(); err != nil {
			panic(err)
		}
	}(resp.Body)
	if err != nil {
		return "", err
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func collectData(body string) []Movie {
	var movieTop250 = make([]Movie, 0)
	dom, _ := goquery.NewDocumentFromReader(strings.NewReader(body))
	dom.Find(".grid_view .item").Each(func(i int, selection *goquery.Selection) {
		Order := stringToInt(removeTrim(selection.Find(".pic em").Text()))
		Url, _ := selection.Find(".pic a").Attr("href")
		Url = removeTrim(Url)
		Title := removeTrim(selection.Find(".info .hd a .title:nth-child(1)").Text())
		SubTitle := removeTrim(selection.Find(".info .hd a .title:nth-child(2)").Text())
		SubTitle = removeSpecificString(SubTitle, []string{"/"})
		CanPlayable := strings.Contains(removeTrim(selection.Find(".info .hd span.playable").Text()), "可播放")
		personnel := removeTrim(selection.Find(".info .bd p:nth-child(1)").Text())
		details := splitPersonnelString(personnel)
		Score, _ := strconv.ParseFloat(removeTrim(selection.Find(".info .bd .star span.rating_num").Text()), 0)
		ValuationNum := removeTrim(selection.Find(".info .bd .star span:last-of-type").Text())
		valuationNum, _ := strconv.ParseInt(strings.TrimSuffix(ValuationNum, "人评价"), 10, 0)
		Describe := removeTrim(selection.Find(".info .bd .quote span.inq").Text())
		movieTop250 = append(movieTop250, Movie{
			Order:        uint(Order),
			Url:          Url,
			Title:        Title,
			SubTitle:     SubTitle,
			Score:        float32(Score),
			ValuationNum: int(valuationNum),
			Describe:     Describe,
			CanPlayable:  CanPlayable,
			Personnel:    details["personnel"],
			Year:         stringToInt(details["year"]),
			Countries:    details["countries"],
			Types:        details["types"],
		})
	})
	return movieTop250
}

func stringToInt(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}

func removeTrim(str string) string {
	return strings.TrimSpace(str)
}

func removeSpecificString(str string, seps []string) string {
	for _, v := range seps {
		str = strings.Trim(str, v)
	}
	return str
}

func removeSliceNull(str []string) []string {
	p := make([]string, 0)
	for _, v := range str {
		if v != "" {
			p = append(p, removeTrim(v))
		}
	}
	return p
}

func splitPersonnelString(str string) map[string]string {
	str = removeTrim(str)
	reYear := regexp.MustCompile(`\d{4}`)
	year := reYear.FindString(str)
	persons := reYear.Split(str, -1)
	persons = removeSliceNull(persons)
	personnel := persons[0]
	if len(persons) == 4 {
		persons = strings.Split(removeTrim(persons[3]), "/")
	} else {
		persons = strings.Split(removeTrim(persons[1]), "/")
	}
	persons = removeSliceNull(persons)
	var countries, types string
	if len(persons) == 3 {
		countries = removeTrim(persons[1])
		types = removeTrim(persons[2])
	} else {
		countries = removeTrim(persons[0])
		types = removeTrim(persons[1])
	}

	return map[string]string{
		"personnel": personnel,
		"year":      year,
		"countries": countries,
		"types":     types,
	}
}

func Entry() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("收集异常信息:", err)
		}
	}()
	var (
		i         = 0
		pullData  = make([]Movie, 0)
		startTime = time.Now()
	)
	for i <= 225 {
		fmt.Println(fmt.Sprintf("https://movie.douban.com/top250?start=%v&filter=", i))
		axios := Axios{
			Url: fmt.Sprintf("https://movie.douban.com/top250?start=%v&filter=", i),
			Header: map[string]string{
				"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36",
			},
		}
		result, _ := axios.responseResult()
		var movieTop250 = collectData(result)
		pullData = append(pullData, movieTop250...)
		i += 25
	}
	fmt.Println(pullData, len(pullData))
	doneTime := time.Now().Sub(startTime)
	fmt.Println("经过时间：", doneTime)
}
