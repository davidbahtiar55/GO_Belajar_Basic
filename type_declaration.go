package main

import "fmt"

func main() {
	type (
		NoKTP   string
		Married bool
	)

	var (
		noKTPDavid    NoKTP   = "4593488828483846"
		marriedStatus Married = true
	)

	fmt.Println(noKTPDavid)
	fmt.Println(marriedStatus)

}
