package main

import "fmt"

func main() {
	var (
		result = 10 + 10

		a = 10
		b = 20
		c = a * b

		negative = -100
		positif  = +100
	)

	var i = 10
	i += 10 // i = i + 10

	i++ // i = i + 1

	fmt.Println(result)
	fmt.Println(c)
	fmt.Println(i)
	fmt.Println(negative)
	fmt.Println(positif)

}
