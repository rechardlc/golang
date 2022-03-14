package main

import (
	"context"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	WaitGroup   sync.WaitGroup
	Client      = &http.Client{}
	chanImgUrls = make(chan string, 9999)
)

func main() {
	for i := 1; i < 12; i++ {
		WaitGroup.Add(1)
		go func(index int) {
			if err := GetHttpHtmlContent(index); err != nil {
				fmt.Println(err)
			}
		}(i)
	}
	WaitGroup.Wait()
	close(chanImgUrls)
	fmt.Println(len(chanImgUrls))
	for {
		url, ok := <-chanImgUrls
		if !ok {
			break
		}
		WaitGroup.Add(1)
		go downloadUrls(url)

	}
	WaitGroup.Wait()
}
func GetHttpHtmlContent(index int) error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	chromeCtx, cancel := chromedp.NewContext(context.Background(), chromedp.WithLogf(log.Printf))
	defer cancel()
	timeoutCtx, cancel := context.WithTimeout(chromeCtx, 40*time.Second)
	defer cancel()
	var htmlContent, pageUrl string
	if index == 1 {
		pageUrl = "https://www.ivsky.com/tupian/"
	} else {
		pageUrl = "https://www.ivsky.com/tupian/index_" + strconv.Itoa(index) + ".html"
	}
	fmt.Println(pageUrl)
	err := chromedp.Run(
		timeoutCtx,
		chromedp.Navigate(pageUrl),
		chromedp.WaitVisible(".ali"),
		chromedp.OuterHTML(`document.querySelector("body")`, &htmlContent, chromedp.ByJSPath),
	)
	if err != nil {
		return err
	}
	GetImageUrls(htmlContent)
	return nil
}
func GetImageUrls(content string) {
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(content))
	if err != nil {
		fmt.Println(err)
	}
	dom.Find(".il_img a img").Each(func(i int, selection *goquery.Selection) {
		if val, ok := selection.Attr("src"); ok {
			chanImgUrls <- fmt.Sprintf("https:%s", val)
		}
	})
	WaitGroup.Done()
}
func downloadUrls(url string) {
	index := strings.LastIndex(url, "/")
	fileName := url[index+1:]
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36")
	resp, err := Client.Do(req)
	if err != nil {
		fmt.Println("load failed", err)
		return
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	path, err := filepath.Abs("")
	fileName = filepath.Join(path, "pictures/images/"+fileName)
	err = ioutil.WriteFile(fileName, bytes, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	WaitGroup.Done()
}
