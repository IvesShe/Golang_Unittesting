package monster

import (
	"testing"
)

func TestStore(t *testing.T) {

	// 先創建一個Monster實例
	monster := &Monster{
		Name:  "ivesshe",
		Age:   20,
		Skill: "Golang",
	}

	res := monster.Store()
	if !res {
		t.Fatalf("monster.Store() 錯誤，希望為=%v 實際為=%v", true, res)
	}

	t.Logf("monster.Store() 測試成功!")
}

func TestReStore(t *testing.T) {

	// 先創建一個Monster實例
	monster := &Monster{}

	res := monster.ReStore()
	if !res {
		t.Fatalf("monster.ReStore() 錯誤，希望為=%v 實際為=%v", true, res)
	}

	// 進一步判斷
	if monster.Name != "ivesshe" {
		t.Fatalf("monster.ReStore() 錯誤，希望為=%v 實際為=%v", "ivesshe", monster.Name)
	}
	if monster.Age != 20 {
		t.Fatalf("monster.ReStore() 錯誤，希望為=%v 實際為=%v", 20, monster.Age)
	}
	if monster.Skill != "Golang" {
		t.Fatalf("monster.ReStore() 錯誤，希望為=%v 實際為=%v", "Golang", monster.Skill)
	}

	t.Logf("monster.Store() 測試成功!")
	//t.Logf("monster.Name 正確，希望為%v 實際為%v", "ivesshe", monster.Name)
}
