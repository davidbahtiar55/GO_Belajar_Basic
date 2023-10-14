package main

import "fmt"

func main() {
	var name1 = "David"
	var name2 = "Bahtiar"
	var result bool = name1 == name2

	var (
		value1 = 100
		value2 = 300
	)

	fmt.Println(result)
	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)

}
