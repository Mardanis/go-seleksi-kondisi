package main

import (
	"fmt"
	"strings"
)

func IfElseOr() {
	var colorBlack, colorBlue string = "BLACK", "BLUE"
	var color string

	fmt.Print("Input Warna Sepatu : ")
	fmt.Scan(&color)
	color = strings.Title(strings.ToUpper(color)) // to UPPERCASE

	if color == colorBlack || color == colorBlue { // jika warna itu hitam atau biru
		fmt.Println("Alhamdulillah jadi beli sepatur warna ", color)
	} else {
		fmt.Println("Maaf Tidak jadi beli karna hanya ada warna ", color)
	}
}
