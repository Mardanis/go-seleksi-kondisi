package main

import (
	"fmt"
)

func IfElseAnd() {
	var nilai int
	var grade, nama string

	fmt.Println("Program Grade Mahasiswa")
	fmt.Print("Input Nama : ")
	fmt.Scan(&nama)
	fmt.Print("Input Nilai : ")
	fmt.Scan(&nilai)

	/*
		A = nilai lebih dari atau sama dengan 90
		B = nilai lebih dari atau sama dengan 80 dan kurang dari 90
		C = nilai lebih dari atau sama dengan 65 dan kurang dari 80
		D = nilai lebih dari atau sama dengan 50 dan kurang dari 65
		E = kurang dari 50
	*/

	if nilai >= 90 { // 90, 91, ....
		grade = "A"
	} else if nilai >= 80 && nilai < 90 {
		grade = "B"
	} else if nilai >= 65 && nilai < 80 {
		grade = "C"
	} else if nilai >= 50 && nilai < 65 {
		grade = "D"
	} else {
		grade = "E"
	}

	fmt.Printf("%s mendapatkan nilai %s", nama, grade)
}
