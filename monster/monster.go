package monster

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Monster struct {
	Name  string
	Age   int
	Skill string
}

// 給Monster綁定方法Store，可以將一個Monster變量(對象)，序列化後保存到文件中

func (this *Monster) Store() bool {
	// 先序列化
	data, err := json.Marshal(this)
	if err != nil {
		fmt.Println("marshal err = ", err)
		return false
	}

	// 保存到文件
	filePath := "./monster.ser"
	ioutil.WriteFile(filePath, data, 0666)
	if err != nil {
		fmt.Println("write file err = ", err)
		return false
	}
	return true
}

// 給Monster綁定方法ReStore，可以將一個序列化的Monster，從文件中讀取，並反序列化為Monster對象，檢查反序列化，名字正確
// 編程測試用例文件store_test.go，編寫測試用例函數TestStore和TestRestore進行測試。

func (this *Monster) ReStore() bool {
	// 先從文件中，讀取序列化的字符串
	filePath := "./monster.ser"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("ReadFile err = ", err)
		return false
	}

	// 使用讀取到的data []byte，對其反序列化
	err = json.Unmarshal(data, this)
	json.Marshal(this)
	if err != nil {
		fmt.Println("UnMarshal err = ", err)
		return false
	}
	return true
}
