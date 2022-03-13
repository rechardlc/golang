package main

import "fmt"

func main() {
	arr := [2]int{1, 2}
	testArr(arr[:])
	fmt.Println(arr)
}
func testArr(nums []int) {
	nums[0] = 10
	fmt.Println(nums)
}
