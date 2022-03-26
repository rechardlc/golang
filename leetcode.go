package main

import (
	"sync"
)
import "fmt"

var ws sync.WaitGroup
var lock sync.Mutex

func main() {
	n := countBinarySubstrings("00110011")
	fmt.Println(n)
}
func countBinarySubstrings(s string) int {
	num, subLen, sLen := 0, 2, len(s)
	if sLen < 2 {
		return num
	}
	for subLen < sLen {
		//ws.Add(1)
		func(_subLen int, _s string, _n *int) {
			for i := 0; i < len(_s); i++ {
				end := i + _subLen
				if end > len(_s) {
					break
				}
				temp := s[i:end]
				compare := ""
				if string(temp[0]) == "0" {
					for j := 0; j < len(temp); j++ {
						if len(compare) >= len(temp)/2 {
							compare += "1"
						} else {
							compare += "0"
						}
					}
				} else {
					for j := 0; j < len(temp)/2; j++ {
						if len(compare) > len(temp)/2 {
							compare += "0"
						} else {
							compare += "1"
						}
					}
				}
				if compare == temp {
					fmt.Println(compare, temp)
					//lock.Lock()
					*_n++
					//lock.Unlock()
				}
			}
			//ws.Done()
		}(subLen, s, &num)
		subLen += 2
	}
	//ws.Wait()
	return num
}
