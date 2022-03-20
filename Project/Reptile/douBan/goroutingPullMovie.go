package douBan

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

// 异步方法~在协程下的结果：818.605646ms
var (
	goWs     sync.WaitGroup
	pageDone = make(chan []Movie, 10)
	client   = &http.Client{}
)

func (a Axios) goResponseResult() (body string, err error) {
	request, err := http.NewRequest(a.Method, a.Url, nil)
	for k, v := range a.Header {
		request.Header.Set(k, v)
	}
	if err != nil {
		return "", err
	}
	resp, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		if err = Body.Close(); err != nil {
			panic(err)
		}
	}(resp.Body)
	fmt.Println("resp.Status:", resp.Status, resp.StatusCode)
	if err != nil {
		return "", err
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func goCollectData(body string) []Movie {
	var movieTop250 = make([]Movie, 0)
	dom, _ := goquery.NewDocumentFromReader(strings.NewReader(body))
	dom.Find(".grid_view .item").Each(func(i int, selection *goquery.Selection) {
		Order := goStringToInt(goRemoveTrim(selection.Find(".pic em").Text()))
		Url, _ := selection.Find(".pic a").Attr("href")
		Url = goRemoveTrim(Url)
		Title := goRemoveTrim(selection.Find(".info .hd a .title:nth-child(1)").Text())
		SubTitle := goRemoveTrim(selection.Find(".info .hd a .title:nth-child(2)").Text())
		SubTitle = goRemoveSpecificString(SubTitle, []string{"/"})
		CanPlayable := strings.Contains(goRemoveTrim(selection.Find(".info .hd span.playable").Text()), "可播放")
		personnel := goRemoveTrim(selection.Find(".info .bd p:nth-child(1)").Text())
		details := goSplitPersonnelString(personnel)
		Score, _ := strconv.ParseFloat(goRemoveTrim(selection.Find(".info .bd .star span.rating_num").Text()), 0)
		ValuationNum := goRemoveTrim(selection.Find(".info .bd .star span:last-of-type").Text())
		valuationNum, _ := strconv.ParseInt(strings.TrimSuffix(ValuationNum, "人评价"), 10, 0)
		Describe := goRemoveTrim(selection.Find(".info .bd .quote span.inq").Text())
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
			Year:         goStringToInt(details["year"]),
			Countries:    details["countries"],
			Types:        details["types"],
		})
	})
	return movieTop250
}

func goStringToInt(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}

func goRemoveTrim(str string) string {
	return strings.TrimSpace(str)
}

func goRemoveSpecificString(str string, seps []string) string {
	for _, v := range seps {
		str = strings.Trim(str, v)
	}
	return str
}

func goRemoveSliceNull(str []string) []string {
	p := make([]string, 0)
	for _, v := range str {
		if v != "" {
			p = append(p, goRemoveTrim(v))
		}
	}
	return p
}

func goSplitPersonnelString(str string) map[string]string {
	str = goRemoveTrim(str)
	reYear := regexp.MustCompile(`\d{4}`)
	year := reYear.FindString(str)
	persons := reYear.Split(str, -1)
	persons = goRemoveSliceNull(persons)
	personnel := persons[0]
	if len(persons) == 4 {
		persons = strings.Split(goRemoveTrim(persons[3]), "/")
	} else {
		persons = strings.Split(goRemoveTrim(persons[1]), "/")
	}
	persons = goRemoveSliceNull(persons)
	var countries, types string
	if len(persons) == 3 {
		countries, types = goRemoveTrim(persons[1]), goRemoveTrim(persons[2])
	} else {
		countries, types = goRemoveTrim(persons[0]), goRemoveTrim(persons[1])
	}

	return map[string]string{
		"personnel": personnel,
		"year":      year,
		"countries": countries,
		"types":     types,
	}
}

func GoEntry() {
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
		goWs.Add(1)
		go func(i int) {
			fmt.Println(fmt.Sprintf("https://movie.douban.com/top250?start=%v&filter=", i))
			axios := Axios{
				Url: fmt.Sprintf("https://movie.douban.com/top250?start=%v&filter=", i),
				Header: map[string]string{
					//"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36",
					"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.2 Safari/605.1.15",
					"Cookie":     "bid=EiFoZyBObDM; ll=\"118318\"; __utmc=30149280; gr_user_id=58c7fa73-6653-4fa2-9966-f842a9456abe; _ga=GA1.2.1269460877.1647445204; viewed=\"35523099\"; ct=y; __utmc=223695111; _vwo_uuid_v2=D0157CDFE68CA4A52FFA15137B156426D|3197525a93babfd52b401c9e4ceea6c5; Hm_lvt_16a14f3002af32bf3a75dfe352478639=1647662684; Hm_lpvt_16a14f3002af32bf3a75dfe352478639=1647662968; __utma=30149280.1269460877.1647445204.1647763393.1647787713.7; __utmz=30149280.1647787713.7.2.utmcsr=google|utmccn=(organic)|utmcmd=organic|utmctr=(not%20provided); __utmb=30149280.3.10.1647787713; dbcl2=\"157813292:H9Mi+0YsVyA\"; ck=DKsQ; _pk_ref.100001.4cf6=%5B%22%22%2C%22%22%2C1647788481%2C%22https%3A%2F%2Fopen.weixin.qq.com%2F%22%5D; _pk_id.100001.4cf6=93716d5b326545aa.1647662542.6.1647788481.1647763748.; _pk_ses.100001.4cf6=*; __utma=223695111.1269460877.1647445204.1647763393.1647788481.6; __utmb=223695111.0.10.1647788481; __utmz=223695111.1647788481.6.2.utmcsr=open.weixin.qq.com|utmccn=(referral)|utmcmd=referral|utmcct=/; push_noty_num=0; push_doumail_num=0",
				},
			}
			result, _ := axios.goResponseResult()
			var movieTop250 = goCollectData(result)
			pageDone <- movieTop250
			fmt.Printf("第%v页爬取完成！\n", i/25)
			fmt.Println("爬取结果：", movieTop250)
			goWs.Done()
		}(i)
		i += 25
	}
	goWs.Wait()
	close(pageDone)
	for movies := range pageDone {
		pullData = append(pullData, movies...)
	}
	//fmt.Println(pullData, len(pullData))
	doneTime := time.Now().Sub(startTime)
	fmt.Println("经过时间：", doneTime)
	fmt.Println("开始入库~")
	insertIntoMovie(pullData) // 入库执行函数
}
func init() {
	createMovieTable()
}
