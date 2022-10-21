package core

import "testing"

func TestStatement(t *testing.T) {
	for i := 0; i <= 3; i++ {
		str, _ := Statement()
		if str == " " {
			t.Fatalf("函数出错")
		}
	}

}
func TestDiary(t *testing.T) {
	for i := 0; i <= 3; i++ {
		str, _ := Diary()
		if str == " " {
			t.Fatalf("函数出错")
		}
	}
}
