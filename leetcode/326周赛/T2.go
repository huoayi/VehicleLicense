package main

func isPrime (num int)[]bool{
	isBool := make([]bool,num+1)
	for i := 1;i <= num;i++{
		isBool[i] = true
	}
	for i := 2; i <= num/i; i++{
		for j := i;j <= num/i; j++ {
			isBool[i*j] = false
		}
	}
	return isBool
}
func distinctPrimeFactors(nums []int) int {
	isBool := isPrime(1000)
	num := 0
	set := make([]int,0)
	for i := 0; i < len(nums); i++ {
		num = num * nums[i]
	}
	for i := 1; i < len(isBool); i++ {
		if isBool[i] == true {
			set = append(set, i)
		}
	}

}
func main() {

}