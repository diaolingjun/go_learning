package main

import "fmt"

func main() {
	data := []string{"我", "爱", "学", "习"}
	fmt.Println(data)
	reverseString(data)
}
func reverseString(s []string) {
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
