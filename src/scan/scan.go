package main

import "fmt"

func main() {
	var s1, s2 string

	n, _ := fmt.Scan(&s1, &s2)
	fmt.Println("입력 개수:", n)
	fmt.Println(s1, s2)
}
