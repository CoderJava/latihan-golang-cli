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

	fmt.Println("                  WARNET                  ")
	fmt.Println("==========================================")
	fmt.Print("Nama Pengunjung\t\t: ")
	scanner.Scan()

	fmt.Print("Keterangan\t\t: ")
	scanner.Scan()
	keterangan := strings.ToLower(scanner.Text())
	if keterangan != "p" && keterangan != "umum" {
		fmt.Println("Inputan keterangan tidak valid.")
		return
	}

	fmt.Print("Durasi (jam)\t\t: ")
	scanner.Scan()
	strDurasi := scanner.Text()
	durasi, err := strconv.Atoi(strDurasi)
	if err != nil {
		fmt.Println("Inputan durasi tidak valid.")
		return
	} else if durasi <= 0 {
		fmt.Println("Inputan durasi minimal 1 jam ya.")
		return
	}

	var totalHargaSewa, totalPembayaran, diskon int
	if keterangan == "p" {
		const hargaSewaPerJam = 4000
		totalHargaSewa = hargaSewaPerJam * durasi
		if durasi >= 5 {
			// Diskon 50%
			diskon = int(float64(totalHargaSewa) * 0.5)
			totalPembayaran = totalHargaSewa - diskon
		} else if durasi >= 3 {
			// Diskon 30%
			diskon = int(float64(totalHargaSewa) * 0.3)
			totalPembayaran = totalHargaSewa - diskon
		}
	} else if keterangan == "umum" {
		// Diskon 10%
		const hargaSewaPerJam = 5000
		totalHargaSewa = hargaSewaPerJam * durasi
		diskon = int(float64(totalHargaSewa) * 0.1)
		totalPembayaran = totalHargaSewa - diskon
	}

	fmt.Printf("Total Harga\t\t: Rp.%d\n", totalHargaSewa)
	fmt.Printf("Diskon Yang Diperoleh\t: Rp.%d\n", diskon)
	fmt.Printf("Total Pembayaran\t: Rp.%d\n", totalPembayaran)
	fmt.Println("==========================================")
	fmt.Print("Uang Bayar\t\t: Rp.")
	scanner.Scan()
	strUangBayar := scanner.Text()
	uangBayar, err := strconv.Atoi(strUangBayar)
	if err != nil {
		fmt.Println("Inputan uang bayar tidak valid.")
		return
	} else if uangBayar < totalPembayaran {
		fmt.Println("Uang bayar tidak mencukupi.")
		return
	}

	uangKembali := uangBayar - totalPembayaran
	fmt.Printf("Uang Kembali\t\t: Rp.%d\n", uangKembali)
	fmt.Println("==========================================")
	fmt.Println("TERIMAKASIH ATAS KUNJUNGAN ANDA")
}
