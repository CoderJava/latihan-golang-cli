package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	daftarHargaBuku := [7]Buku{
		{Nama: "Sukses Belajar Borland C++", Harga: 50000},
		{Nama: "Kunci Pribadi Sukses", Harga: 35000},
		{Nama: "Mencari Mutiara Di Dasar Hati", Harga: 45000},
		{Nama: "Belajar Java Dasar", Harga: 80000},
		{Nama: "Belajar Golang Dasar", Harga: 120000},
		{Nama: "Belajar Dart Dasar", Harga: 90000},
		{Nama: "Belajar Kotlin Dasar", Harga: 95000},
	}

	fmt.Println("           TOKO BUKU SEJAHTERA           ")
	fmt.Println("            Jl.Keadilan No.16            ")
	fmt.Println("          Telp.7236573-72365741          ")
	fmt.Println("=========================================")
	fmt.Println("Daftar Buku Yang Tersedia")
	for index, buku := range daftarHargaBuku {
		fmt.Printf("[%d] %s\n", index+1, buku.Nama)
	}
	fmt.Print("Masukkan kode buku\t: ")
	scanner.Scan()
	inputKodeBukuBeli := scanner.Text()
	var listKodeBukuBeli = strings.Split(inputKodeBukuBeli, ",")
	var listBukuBeli []Buku
	for _, strItemKodeBukuBeli := range listKodeBukuBeli {
		kodeBukuBeli, err := strconv.Atoi(strItemKodeBukuBeli)
		if err != nil {
			fmt.Println("Inputan kode buku tidak valid. Error:", err.Error())
			return
		} else if kodeBukuBeli < 0 || kodeBukuBeli > len(daftarHargaBuku) {
			fmt.Println("Inputan kode buku tidak tersedia.")
			return
		}
		listBukuBeli = append(listBukuBeli, daftarHargaBuku[kodeBukuBeli-1])
	}
	fmt.Println()
	fmt.Println("Daftar Buku Yang Dibeli")
	totalHargaBeliBuku := 0
	for index, bukuBeli := range listBukuBeli {
		hargaBuku := bukuBeli.Harga
		fmt.Printf("%d. %s (Rp.%d)\n", index+1, bukuBeli.Nama, hargaBuku)
		totalHargaBeliBuku += hargaBuku
	}
	fmt.Println()
	fmt.Printf("Total\t\t: Rp.%d\n", totalHargaBeliBuku)

	if len(listBukuBeli) >= 3 {
		diskon := int(float64(totalHargaBeliBuku) * 0.1)
		totalHargaBeliBuku -= diskon
		fmt.Printf("Diskon\t\t: Rp.%d\n", diskon)
		fmt.Println("Bonus\t\t: Buku Saku")
	}

	ppn := int(float64(totalHargaBeliBuku) * 0.2)
	totalHargaBeliBuku += ppn
	fmt.Printf("PPN 2%%\t\t: Rp.%d\n", ppn)
	fmt.Printf("Jumlah Bayar\t: Rp.%d\n", totalHargaBeliBuku)

	fmt.Print("Bayar\t\t: Rp.")
	scanner.Scan()
	inputBayar := scanner.Text()
	bayar, err := strconv.Atoi(inputBayar)
	if err != nil {
		fmt.Println("Inputan bayar tidak valid. Error:", err.Error())
		return
	} else if bayar < totalHargaBeliBuku {
		fmt.Println("Uang pembayaran tidak cukup.")
		return
	}

	kembali := bayar - totalHargaBeliBuku
	fmt.Printf("Kembali\t\t: Rp.%d", kembali)
}

type Buku struct {
	Nama  string
	Harga int
}
