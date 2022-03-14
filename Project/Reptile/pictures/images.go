package main

// 只能遍历那种静态渲染的图片页面，对于动态渲染的无法爬虫
import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	chanImageUrls chan string
	wg            sync.WaitGroup
	chanTask      chan string
	reImg         = `https?://[^"]+?(\.((jpg)|(png)|(jpeg)|(gif)|(bmp)))`
)

func handleError(err error, why string) {
	if err != nil {
		fmt.Println(why)
	}
}

func buildBase() {
	chanImageUrls = make(chan string, 100000)
	chanTask = make(chan string, 26)
	for i := 1; i <= 26; i++ {
		wg.Add(1)
		// 获取26页图片的url
		go getImgUrls("https://www.ivsky.com/tupian/index_" + strconv.Itoa(i+1) + ".html")
	}
	wg.Add(1)
	// 检查是否将所有的图片url都获取到了
	go CheckOk()
	for i := 0; i < 5; i++ {
		wg.Add(1)
		// 开启新的协程~去下载保存在chanImageUrls中的图片url
		go Downloading()
	}
	wg.Wait()
}

func getImgUrls(url string) {
	urls := getIms(url)
	for _, url := range urls {
		chanImageUrls <- url
	}
	chanTask <- url // 一个页面一个channel
	wg.Done()
}

//
//  getIms
//  @Description: 获取当前页面上所有的img~url
//  @param url
//  @return urls
//
func getIms(url string) (urls []string) {
	pageStr := GetPageStr(url)
	re := regexp.MustCompile(reImg)
	results := re.FindAllStringSubmatch(pageStr, -1)
	for _, result := range results {
		url := result[0]
		urls = append(urls, url)
	}
	return
}

func GetPageStr(url string) string {
	resp, err := http.Get(url)
	handleError(err, "http.Get url")
	defer resp.Body.Close()
	pageBytes, err := ioutil.ReadAll(resp.Body)
	return string(pageBytes)
}

func CheckOk() {
	var count int
	for {
		url := <-chanTask
		fmt.Printf("%s 完成爬取任务\n", url)
		count++
		if count == 26 {
			close(chanImageUrls)
			break
		}
	}
	wg.Done()
}

func Downloading() {
	for url := range chanImageUrls {
		filename := GetFileNameFormUrl(url)
		ok := DownloadFile(url, filename)
		if ok {
			fmt.Println("下载成功：", filename)
		} else {
			fmt.Println("下载失败", filename)
		}
	}
	wg.Done()
}

func GetFileNameFormUrl(url string) string {
	lastIndex := strings.LastIndex(url, "/")
	filename := url[lastIndex+1:]
	timePreFix := strconv.Itoa(int(time.Now().UnixNano()))
	filename = timePreFix + "_" + filename
	return filename
}

func DownloadFile(url, fileName string) bool {
	resp, err := http.Get(url)
	handleError(err, "Http.get.url")
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	fileName = "/Users/ruoshuiyiliao/Desktop/golang/Project/Reptile/pictures/" + fileName
	err = ioutil.WriteFile(fileName, bytes, 0666)
	if err != nil {
		return false
	}
	return true
}

func init() {
	fmt.Println("init")
}
func main() {
	fmt.Println("main")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("main function -> catch err:")
		}
	}()
	//buildBase()
	resp, _ := http.Get("https://www.ivsky.com/tupian/index_2.html")
	bytes, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bytes))
}
