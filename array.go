package main

import "fmt"

func main() {
	var name [3]string

	name[0] = "David"
	name[1] = "Bahtiar"
	name[2] = "DBR"

	fmt.Println(name[0])
	fmt.Println(name[1])
	fmt.Println(name[2])

	var values = [3]int{
		90,
		80,
		70,
	}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])
}
