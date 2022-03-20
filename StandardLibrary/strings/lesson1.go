package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	var str = `   dsfdsfsfdsfds fdsfds fdsfds fdsfds fd

sf/nds`
	str = strings.TrimFunc(str, func(r rune) bool {
		//return unicode.IsLetter(r)
		return unicode.IsSpace(r)
	})
	fmt.Println(str)
}
