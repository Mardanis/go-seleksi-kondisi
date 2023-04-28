package main

import "fmt"

func IfElseTemporary() {
	var nilai float32 // nilai default dari float 0.0 / 0
	// var angka int     // 0
	// var huruf string // ""
	// var isActive bool

	fmt.Print("Input Nilai : ")
	fmt.Scan(&nilai)

	// hasil itu adalah temporary variable di block if else
	if hasil := nilai / 2; hasil >= 5 {
		fmt.Println("Lulus dengan hasil ", hasil)
	} else if hasil >= 1 && hasil < 5 {
		fmt.Println("Mengulang dengan hasil", hasil)
	} else {
		fmt.Println("DO ", hasil)
	}

	// fmt.Println("Hasil", isActive)
}
