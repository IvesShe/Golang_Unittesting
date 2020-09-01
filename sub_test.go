package main

import (
	"testing" // 引入go的testing框架包
)

// 編寫測試用例，去測試addUpper是否正確，首字母要大寫
func TestGetSub(t *testing.T) {

	// 調用
	res := getSub(10, 3)
	if res != 7 {
		t.Fatalf("AddUpper(10)執行錯誤，期望值=%v 實際值=%v\n", 7, res)
	}

	// 如果正確，輸出日誌
	t.Logf("getSub(10,3) 執行正確...")
}
