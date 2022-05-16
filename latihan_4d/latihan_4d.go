package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Kode Paket [1..3]\t: ")
	scanner.Scan()
	kodePaket := scanner.Text()
	var namaPaket string
	var hargaPaket uint16
	if kodePaket == "1" {
		namaPaket = "PAKET HEMAT"
		hargaPaket = 7500
	} else if kodePaket == "2" {
		namaPaket = "PAKET NASI"
		hargaPaket = 10000
	} else if kodePaket == "3" {
		namaPaket = "PAKET SPESIAL"
		hargaPaket = 15000
	} else {
		fmt.Println("Inputan kode paket tidak valid.")
		return
	}

	fmt.Print("Jumlah Beli\t\t: ")
	scanner.Scan()
	strJumlahBeli := scanner.Text()
	jumlahBeli, err := strconv.Atoi(strJumlahBeli)
	if err != nil {
		fmt.Println("Inputan jumlah beli tidak valid.")
		return
	} else if jumlahBeli <= 0 {
		fmt.Println("Jumlah beli minimal harus 1 ya.")
		return
	}

	fmt.Print("Kode Kasir\t\t: ")
	scanner.Scan()
	kodeKasir := scanner.Text()

	fmt.Print("Nama Kasir\t\t: ")
	scanner.Scan()
	namaKasir := scanner.Text()

	fmt.Println("====================================")
	fmt.Println("           SEJAHTERA CAFE           ")
	fmt.Println("          Jl. Juang No.167          ")
	fmt.Println("        Telp.7236574-7236574        ")
	fmt.Println("====================================")
	fmt.Printf("\t%s\t%s\n", namaKasir, kodeKasir)
	fmt.Printf("\t%s\n", namaPaket)
	fmt.Printf("\t%d x %d\n", jumlahBeli, hargaPaket)

	total := jumlahBeli * int(hargaPaket)
	fmt.Printf("Total\t\t: Rp.%d\n", total)

	ppn := int(float64(total) * 0.1)
	fmt.Printf("PPN 10%%\t\t: Rp.%d\n", ppn)

	jumlahBayar := total + ppn
	fmt.Printf("Jumlah Bayar\t: Rp.%d\n", jumlahBayar)

	fmt.Printf("Bayar\t\t: Rp.")
	scanner.Scan()
	strBayar := scanner.Text()
	bayar, err := strconv.Atoi(strBayar)
	if err != nil {
		fmt.Println("Inputan bayar tidak valid.")
		return
	} else if bayar < jumlahBayar {
		fmt.Println("Uang pembayaran tidak mencukupi.")
		return
	}

	kembali := bayar - jumlahBayar
	fmt.Printf("Kembali\t\t: Rp.%d\n", kembali)
	fmt.Println("====================================")
	fmt.Println("         SELAMAT MENIKMATI          ")
	fmt.Println("            TERIMA KASIH            ")
}
