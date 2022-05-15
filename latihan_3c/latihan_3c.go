package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Masukkan nilai x: ")
	scanner.Scan()
	strX := scanner.Text()
	x, err := strconv.Atoi(strX)
	if err != nil {
		fmt.Println("Inputan nilai x tidak valid. Harus angka semua ya.")
		return
	}

	fmt.Print("Masukkan nilai y: ")
	scanner.Scan()
	strY := scanner.Text()
	y, err := strconv.Atoi(strY)
	if err != nil {
		fmt.Println("Inputan nilai y tidak valid. Harus angka semua ya.")
		return
	}

	fmt.Println()
	fmt.Println("Semua akan menghasilkan nilai true dan false ya")
	fmt.Println("===============================================")

	fmt.Println("Operator Relasi")
	resultRelation1 := x == y
	resultRelation2 := x != y
	fmt.Printf("Hasil dari %d == %d = %t\n", x, y, resultRelation1)
	fmt.Printf("Hasil dari %d != %d = %t\n", x, y, resultRelation2)
	fmt.Println()

	fmt.Println("Operator Logika")
	fmt.Printf("Hasil dari %t && %t = %t\n", resultRelation1, resultRelation2, resultRelation1 && resultRelation2)
	fmt.Printf("Hasil dari %t || %t = %t\n", resultRelation1, resultRelation2, resultRelation1 || resultRelation2)
	fmt.Println()

	fmt.Println("===============================================")
	fmt.Println("Operator Relasi")
	resultRelation1 = x > y
	resultRelation2 = x < y
	fmt.Printf("Hasil dari %d > %d = %t\n", x, y, resultRelation1)
	fmt.Printf("Hasil dari %d < %d = %t\n", x, y, resultRelation2)
	fmt.Println()
	
	fmt.Println("Operator Logika")
	fmt.Printf("Hasil dari %t && %t = %t\n", resultRelation1, resultRelation2, resultRelation1 && resultRelation2)
	fmt.Printf("Hasil dari %t || %t = %t\n", resultRelation1, resultRelation2, resultRelation1 || resultRelation2)
	fmt.Println()

	fmt.Println("===============================================")
	fmt.Println("Operator Relasi")
	resultRelation1 = x >= y
	resultRelation2 = x <= y
	fmt.Printf("Hasil dari %d >= %d = %t\n", x, y, resultRelation1)
	fmt.Printf("Hasil dari %d <= %d = %t\n", x, y, resultRelation2)
	fmt.Println()

	fmt.Println("Operator Logika")
	fmt.Printf("Hasil dari %t && %t = %t\n", resultRelation1, resultRelation2, resultRelation1 && resultRelation2)
	fmt.Printf("Hasil dari %t || %t = %t\n", resultRelation1, resultRelation2, resultRelation1 || resultRelation2)

}
