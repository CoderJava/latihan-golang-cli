package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("            Toko Elektronik Sejahtera             \n")
	fmt.Print("==================================================\n")
	fmt.Print("Nama Barang\t: ")
	scanner.Scan()

	fmt.Print("Harga Barang\t: ")
	scanner.Scan()
	strHargaBarang := scanner.Text()
	hargaBarang, err := strconv.Atoi(strHargaBarang)
	if err != nil {
		fmt.Println("Inputan harga barang tidak valid.")
		return
	} else if hargaBarang < 0 {
		fmt.Println("Nilai harga barang tidak boleh minus.")
		return
	}

	fmt.Print("Jumlah Beli\t: ")
	scanner.Scan()
	strJumlahBeli := scanner.Text()
	jumlahBeli, err := strconv.Atoi(strJumlahBeli)
	if err != nil {
		fmt.Print("Inputan jumlah beli tidak valid.")
		return
	} else if jumlahBeli < 0 {
		fmt.Println("Nilai jumlah beli tidak boleh minus.")
		return
	}

	jumlahBayar := hargaBarang * jumlahBeli
	fmt.Printf("Jumlah Bayar\t: %d\n", jumlahBayar)

	ppn := int(float64(jumlahBayar) * 0.10)
	fmt.Printf("PPN (10%%)\t: %d\n", ppn)

	totalBayar := jumlahBayar + ppn
	fmt.Printf("Total Bayar\t: %d\n", totalBayar)
	fmt.Print("==================================================\n")
	fmt.Print("Uang Bayar\t: ")
	scanner.Scan()
	strUangBayar := scanner.Text()
	uangBayar, err := strconv.Atoi(strUangBayar)
	if err != nil {
		fmt.Println("Inputan uang bayar tidak valid.")
		return
	} else if uangBayar < totalBayar {
		fmt.Println("Uang bayar tidak mencukupi.")
		return
	}

	uangKembali := uangBayar - totalBayar
	fmt.Printf("Uang Kembali\t: %d\n", uangKembali)
	fmt.Print("==================================================\n")
	fmt.Println("         TERIMAKASIH ATAS KUNJUNGANNYA          ")
}
