package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "David",
		"address": "Karawang",
	}

	person["title"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	var book map[string]string = make(map[string]string)
	book["title"] = "Belajar G0-Lang"
	book["author"] = "David"
	book["ups"] = "Salah"

	fmt.Println(book)
	fmt.Println(len(book))
}
