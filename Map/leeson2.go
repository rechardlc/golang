package main

import "fmt"

func main() {
	findMastSubstring("")
}
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
