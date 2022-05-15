package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("       Toko Beras Murah         ")
	fmt.Println("================================")
	fmt.Print("Jumlah beli beras (kg)\t: ")
	scanner.Scan()
	strJumlahBeliBeras := scanner.Text()
	jumlahBeliBeras, err := strconv.Atoi(strJumlahBeliBeras)
	if err != nil {
		fmt.Println("Inputan jumlah beli beras tidak valid.")
		return
	} else if jumlahBeliBeras <= 0 {
		fmt.Println("Tidak boleh 0 atau minus ya.")
		return
	}

	const hargaPerKilo = 4500
	totalBayar := hargaPerKilo * jumlahBeliBeras
	fmt.Printf("Total belanja Anda\t: Rp.%d\n", totalBayar)

	if jumlahBeliBeras >= 23 {
		fmt.Println("Bonus 1 liter ice cream chocolate")
	}
	fmt.Print("Terimakasih atas kunjungannya")
}
