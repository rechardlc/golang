package fetcher

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

func randUserAgent() (string, *url.URL) {
	headers := []string{
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/104.0.0.0 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.93 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.93 Safari/537.36",
		"Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.93 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.181 Safari/537.36",
	}
	ips := []string{
		"117.86.11.34:8888",
		"51.158.152.223:3128",
		"165.225.56.117:10605",
		"103.145.45.6:55443",
		"195.138.73.54:44017",
		"94.102.201.1:1500",
		"89.237.33.129:37647",
		"49.12.77.56:3128",
		"165.225.112.67:10605",
		"154.85.58.149:80",
		"192.162.192.148:55443",
		"36.66.103.75:8080",
		"165.225.38.92:10605",
		"185.142.43.217:8080",
		"185.195.69.164:3128",
		"182.52.83.111:8080",
		"103.142.21.197:8080",
		"203.150.128.9:8080",
		"46.188.53.7:8009",
		"181.114.206.92:9090",
		"175.101.85.33:8080",
		"89.108.157.106:8080",
		"202.62.52.4:8080",
		"203.32.121.242:80",
	}
	rand.Seed(time.Now().Unix())
	randHeaderInt := rand.Intn(len(headers))
	randIpInt := rand.Intn(len(ips))
	proxy, _ := url.Parse("http://" + ips[randIpInt])
	return headers[randHeaderInt], proxy
}

/*
	Fetch 会存在异常，不可能每一个都成功，所以需要返回一个error
*/
var rateLimiter = time.Tick(time.Millisecond * 10)

func Fetch(uri string) ([]byte, error) {
	//<-rateLimiter
	// 自定义客服端~封装新的Request
	userAgent, proxy := randUserAgent()
	//proxy, _ := url.Parse("http://39.175.75.15:30001/")
	client := http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(nil),
		},
	}
	fmt.Println("代理IP addr：", proxy)
	request, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		log.Printf("%s\n", err.Error())
		return nil, err
	}
	request.Header.Set("User-Agent", userAgent)
	resp, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d", resp.StatusCode)
	}
	return ioutil.ReadAll(resp.Body)
}
