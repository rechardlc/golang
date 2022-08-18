package main

import (
	"fmt"
	"regexp"
)

const text = `My email is ccmous1@gmail.com.com My email is ccmous1@gmail.com.com My email is ccmous1@gmail.com.com`

func main() {
	re := regexp.MustCompile(`\w+@[\w.]+\.\w+`)
	fmt.Println(re.FindAllString(text, -1))
}
