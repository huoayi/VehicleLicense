package main

import (
	"fmt"
	"math"
)

func main() {
fmt.Println(takeCharacters("aaabcbabaac",2))
}
//给你一个由字符 'a'、'b'、'c' 组成的字符串 s 和一个非负整数 k 。每分钟，你可以选择取走 s 最左侧 还是 最右侧 的那个字符。
//
//你必须取走每种字符 至少 k 个，返回需要的 最少 分钟数；如果无法取到，则返回 -1 。
func takeCharacters(s string, k int) int {
	if k == 0{
		return 0
	}
	nums := [3]int{}
	for _, v := range s{
		nums[v - 'a']++
	}
	for i := 0; i < 3; i++ {
		if nums[i] < k{
			return -1
		}
	}

	max := 0
	l := 0
	for i, v := range s{
		nums[v-'a']--

		for nums[0] < k || nums[1] < k || nums[2] < k{
			nums[s[l]-'a']++
			l++
		}
		max = int(math.Max(float64(i-l+1),float64(max)))
	}

	return len(s) - max
}