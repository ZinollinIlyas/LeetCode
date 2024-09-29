package main

import (
	"fmt"
	"github.com/ZinollinIlyas/LeetCode/easy"
)

func main() {

	var arr []string

	arr = append(arr, "flower")
	arr = append(arr, "flow")
	arr = append(arr, "flight")

	res := easy.LongestCommonPrefix(arr)
	fmt.Println(res)
}
