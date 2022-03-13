package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Job 实现算法：
//  计算一个数字的各个位数之和，例如数字123，结果为1+2+3=6
//  随机生成数字进行计算
type Job struct {
	Id, RandNum int
}
type Result struct {
	job *Job
	sum int
}

// 创建工作池
/**
num开辟协程数量、
*/
func createPool(num int, jobChan chan *Job, resultChan chan *Result) {
	for i := 0; i < num; i++ {
		go func(jobChan chan *Job, resultChan chan *Result) {
			for job := range jobChan {
				rNum := job.RandNum
				var sum int
				for rNum != 0 {
					tmp := rNum % 10
					sum += tmp
					rNum /= 10
				}
				r := &Result{
					job: job,
					sum: sum,
				}
				resultChan <- r
			}
		}(jobChan, resultChan)
	}
}

func main() {
	jobChan := make(chan *Job, 128)
	resultChan := make(chan *Result, 128)
	createPool(64, jobChan, resultChan)
	go func(resultChan chan *Result) {
		for result := range resultChan {
			fmt.Printf("job id:%v randnum:%v result:%d\n", result.job.Id, result.job.RandNum, result.sum)
		}
	}(resultChan)
	var id int
	for {
		id++
		rNum := rand.Int()
		job := &Job{id, rNum}
		jobChan <- job
		if id > 10 {
			time.Sleep(time.Second * 10)
			break
		}
	}
}
