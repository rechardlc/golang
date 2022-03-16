package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	// 读取年月日时分秒
	year, month, day, hour, minute, second, ms := now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), now.Nanosecond()
	format := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d:%02d\n\n", year, month, day, hour, minute, second, ms)
	fmt.Println("format:", format)
	// 读取时间戳
	timeStamp1 := now.Unix()
	timeStamp2 := now.UnixNano()
	timeStamp3 := now.UnixMicro()
	timeStamp4 := now.UnixMilli()
	fmt.Println("timeStamp:", timeStamp1, timeStamp2, timeStamp3, timeStamp4)
	fmt.Println(time.September)
	fmt.Println(int64(time.Second / time.Millisecond))
	fmt.Println(time.Duration(10 * (time.Second)))
	fmt.Println(now.Add(time.Hour))
	fmt.Println(now.Sub(time.Now()))
	// 定时器
	//ticker := time.Tick(time.Second) // 定义个一个定时器~每一秒钟执行一次
	//fmt.Println(ticker)
	//for i := range ticker {
	//	fmt.Println(i)
	//}
	// 格式化
	fmt.Println(now.Format("2006/01/02"))
	fmt.Println(now)
	// 加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 按照指定时区和指定格式解析字符串时间
	timeObj, err := time.ParseInLocation("2006/01/02 15:04:05", "2019/08/04 14:15:20", loc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj)
	fmt.Println(timeObj.Sub(now), "time~")
}
