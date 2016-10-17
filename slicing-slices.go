package main

import "fmt"

func main() {
	var a []byte
	s := []string{"2", "3", "5", "7", "11", "13"}
	a = s
	str := fmt.Sprintf("%s", s)
	str1 := string(s)
	fmt.Println("s ==: ", s)
	fmt.Println("s[1, 10] ==: ", s[1:4])

	//missing low index implies 0
	fmt.Println("s[:3] ==: ", s[:3])

	//missing high index implies len of slice
	fmt.Println("s[4:] ==: ", s[4:])
	fmt.Println("str: ", str)
	fmt.Println("str1: ", a)
}
