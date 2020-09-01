package main

import (
	"fmt"
	"testing" // 引入go的testing框架包
)

// 編寫測試用例，去測試addUpper是否正確，首字母要大寫
func TestAddUpper(t *testing.T) {

	// 調用
	res := addUpper(10)
	if res != 55 {
		t.Fatalf("AddUpper(10)執行錯誤，期望值=%v 實際值=%v\n", 55, res)
	}

	// 如果正確，輸出日誌
	t.Logf("AddUpper(10) 執行正確...")
}

func TestHello(t *testing.T) {
	fmt.Println("TestHello被調用...")
}
