package douBan

import (
	"time"
)

type Book struct {
	Title       string    `json:"title" db:"title"`
	SubTitle    string    `json:"subTitle" db:"sub_title"`
	Author      string    `json:"author" db:"author"`
	PublishTime time.Time `json:"publishTime" db:"publish_time"`
	Price       string    `json:"price" db:"price"`
	Score       float32   `json:"score" db:"score"`
	CommentNum  int       `json:"commentNum" db:"comment_num"`
	ImgUrl      string    `json:"imgUrl" db:"img_url"`
}

type Movie struct {
	Id           uint    `json:"id" db:"id"`
	Order        uint    `json:"order" db:"m_order"`
	Title        string  `json:"title" db:"title"`
	Url          string  `json:"url" db:"url"`
	SubTitle     string  `json:"subTitle" db:"sub_title"`
	Score        float32 `json:"score" db:"score"`
	ValuationNum int     `json:"valuationNum" db:"valuation_num"`
	Describe     string  `json:"describe" db:"m_describe"`
	CanPlayable  bool    `json:"canPlayable" db:"can_playable"`
	Year         int     `json:"year" db:"m_year"`
	Personnel    string  `json:"personnel" db:"personnel"`
	Countries    string  `json:"countries" db:"countries"`
	Types        string  `json:"types" db:"types"`
}

//type Movie struct {
//	Order        uint    `db:"order"`
//	Title        string  `db:"title"`
//	Url          string  `db:"url"`
//	SubTitle     string  `db:"sub_title"`
//	Score        float32 `db:"score"`
//	ValuationNum int     `db:"valuation_num"`
//	Describe     string  `db:"describe"`
//	CanPlayable  bool    `db:"can_playable"`
//	Year         int     `db:"year"`
//	Personnel    string  `db:"personnel"`
//	Countries    string  `db:"countries"`
//	Types        string  `db:"types"`
//}

type Axios struct {
	Header      map[string]string
	Url, Method string
}

var Method = map[string]string{
	"GET":    "GET",
	"POST":   "POST",
	"PUT":    "PUT",
	"DELETE": "DELETE",
	"HEAD":   "HEAD",
}

var Header = map[string]string{
	//"Host":               "movie.douban.com",
	//"Connection":         "keep-alive",
	//"Pragma":             "no-cache",
	//"Cache-Control":      "no-cache",
	//"sec-ch-ua":          `" Not A;Brand";v="99", "Chromium";v="99", "Google Chrome";v="99"`,
	//"sec-ch-ua-mobile":   " ?0",
	"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36",
	//"sec-ch-ua-platform": "macOS",
	//"Accept":             "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
	//"Sec-Fetch-Site":     "same-origin",
	//"Sec-Fetch-Mode":     "cors",
	//"Referer":            "https://movie.douban.com/top250?start=0&filter=",
	//"Accept-Encoding":    "gzip, deflate, br",
	//"Accept-Language":    "en,zh-CN;q=0.9,zh;q=0.8",
	//"Cookie":             "bid=EiFoZyBObDM; ll=\"118318\"; __utmc=30149280; __utmz=30149280.1647445204.1.1.utmcsr=baidu|utmccn=(organic)|utmcmd=organic; gr_user_id=58c7fa73-6653-4fa2-9966-f842a9456abe; _ga=GA1.2.1269460877.1647445204; viewed=\"35523099\"; ct=y; __utma=30149280.1269460877.1647445204.1647445204.1647662118.2; _pk_ref.100001.4cf6=%5B%22%22%2C%22%22%2C1647662542%2C%22https%3A%2F%2Fwww.douban.com%2F%22%5D; __utma=223695111.1269460877.1647445204.1647662542.1647662542.1; __utmc=223695111; __utmz=223695111.1647662542.1.1.utmcsr=douban.com|utmccn=(referral)|utmcmd=referral|utmcct=/; _vwo_uuid_v2=D0157CDFE68CA4A52FFA15137B156426D|3197525a93babfd52b401c9e4ceea6c5; Hm_lvt_16a14f3002af32bf3a75dfe352478639=1647662684; Hm_lpvt_16a14f3002af32bf3a75dfe352478639=1647662968; _pk_id.100001.4cf6=93716d5b326545aa.1647662542.1.1647662996.1647662542",
}
