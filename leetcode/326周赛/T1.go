package main

import (
	"fmt"
)

func countDigits(num int) int {
	ans := 0
	ind := num
	for	ind != 0{
		k := ind % 10
		if num % k == 0 {
			ans++
		}
		ind = ind /10
	}
	return ans
}
func main() {
	fmt.Println(countDigits(121))
}