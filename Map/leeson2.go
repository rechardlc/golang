package main

import (
	"fmt"
)

func main() {
	findMastSubstring("sdsdsdsdafdfd")
	lengthOfNonRepeatingSubstr("这个是John学习了Tom")
}

/*
	寻找最长不含有重复字符的子串
	golang思路：通过双循环，将字符串中的某一个字符位给map，若map中存在~则取对应的长度，若不存在~则继续追加到map中
*/
func findMastSubstring(s string) {
	substring := ""
	for i := 0; i < len(s); i++ {
		m := make(map[interface{}]interface{})
		for j := i; j < len(s); j++ {
			if _, ok := m[s[j]]; ok {
				if len(substring) < len(s[i:j]) {
					substring = s[i:j]
				}
				break
			}
			m[s[j]] = s[j]
		}
	}
	fmt.Println(substring)
}
func lengthOfNonRepeatingSubstr(s string) {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0
	for i, ch := range []rune(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	fmt.Println(maxLength)
}
