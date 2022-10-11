package go_10_10_homework

import (
	"fmt"
	"testing"
)

func TestReverseString(t *testing.T) {
	data := []string{"我", "爱", "学", "习"}
	fmt.Println(data)
	ReverseString(data)
}
func ReverseString(s []string) {
	var i int
	var j int = len(s) - 1
	var tmp string
	for i < j {
		tmp = s[i]
		s[i] = s[j]
		s[j] = tmp
		i++
		j--
	}

	fmt.Printf("%v", s)
}
