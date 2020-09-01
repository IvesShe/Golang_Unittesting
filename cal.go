package main

import (
	"fmt"
)

// 一個被測試的函數
func addUpper(n int) int {
	res := 0
	for i := 1; i <= n; i++ {
		res += i
	}
	return res
}

// 求兩個數的差
func getSub(n1 int, n2 int) int {
	return n1 - n2
}

func hello() {
	fmt.Println("hello Golang")
}
