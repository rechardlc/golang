package parser

import (
	"crawler/engine"
	"regexp"
)

//var cityRe = regexp.MustCompile(`<th><a href="(http://album.zhenai.com/u/\d+)"[^>]*>([^<]+)</a></th>`)
var cityRe = regexp.MustCompile(`<th><a href="(http://album.zhenai.com/u/\d+)" target="_blank">([^<]*)</a></th></tr> <tr><td width="180"><span class="grayL">性别：</span>([^<]*)</td>`)

// City 第三个解析：城市解析
func City(contents []byte) (result engine.ParseResult) {
	matches := cityRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		nickName := string(m[2])
		gender := string(m[3])
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParserFunc: func(contents []byte) engine.ParseResult {
				return ParserProfile(contents, map[string]interface{}{
					"nickName": nickName,
					"gender":   gender,
				})
			},
		})
		result.Items = append(result.Items, "User: "+nickName)
	}
	return
}
