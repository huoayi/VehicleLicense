package main

import (
	"fmt"
	"math"
)

func main() {
	str := []string{"hsdqinnoha","mqhskgeqzr","zemkwvqrww","zemkwvqrww","daljcrktje","fghofclnwp","djwdworyka","cxfpybanhd","fghofclnwp","fghofclnwp"}
	fmt.Println(closetTarget(str,"zemkwvqrww",8))
}

//给你一个下标从 0 开始的 环形 字符串数组 words 和一个字符串 target 。环形数组 意味着数组首尾相连。
//
//形式上， words[i] 的下一个元素是 words[(i + 1) % n] ，而 words[i] 的前一个元素是 words[(i - 1 + n) % n] ，其中 n 是 words 的长度。
//从 startIndex 开始，你一次可以用 1 步移动到下一个或者前一个单词。
//
//返回到达目标字符串 target 所需的最短距离。如果 words 中不存在字符串 target ，返回 -1 。
func closetTarget(words []string, target string, startIndex int) int {
	ind := []int{}
	le := len(words)
	for i := 0;i < le;i++ {
		if words[i] == target{
			ind = append(ind, i)
		}
	}
	if len(ind) == 0 {
		return -1
	}
	max := math.MaxInt
	for _,v := range ind{
		if v > startIndex{
			max =  int(math.Min(math.Abs(float64(v-startIndex)),math.Min(math.Abs(float64(le+startIndex-v)),float64(max))))
		}else{
			max =  int(math.Min(math.Abs(float64(v-startIndex)),math.Min(math.Abs(float64(le-startIndex+v)),float64(max))))
		}

	}
	return max
}