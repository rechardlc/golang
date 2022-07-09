package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func convertToBin(n int) string {
	result := ""
	for n > 0 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
		n /= 2
	}
	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	fmt.Println(
		convertToBin(5),
		convertToBin(10),
	)
	printFile("/Users/ruoshuiyiliao/Desktop/golang/ControlFlow/text.txt")
}
