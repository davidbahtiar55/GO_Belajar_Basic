package main

import "fmt"

func main() {
	var (
		ujian   = 80
		absensi = 90

		lulusUjian   = ujian >= 80
		lulusAbsensi = absensi >= 90
		lulus        = lulusUjian && lulusAbsensi
	)

	fmt.Println(lulus)
	fmt.Println(lulusUjian)
	fmt.Println(lulusAbsensi)
}
