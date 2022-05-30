package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	listPaketUlangTahun := []paketUlangTahun{
		{
			kode:  1,
			nama:  "Mewah",
			harga: 2000000,
		},
		{
			kode:  2,
			nama:  "Sedang",
			harga: 1500000,
		},
		{
			kode:  3,
			nama:  "Sederhana",
			harga: 1000000,
		},
	}
	listAtraksiUlangTahun := []atraksiUlangTahun{
		{
			kode:  1,
			nama:  "Badut",
			harga: 300000,
		},
		{
			kode:  2,
			nama:  "Sulap",
			harga: 500000,
		},
		{
			kode:  3,
			nama:  "Penyanyi Cilik",
			harga: 600000,
		},
	}

	fmt.Println("PAKET ULANG TAHUN")
	fmt.Println("=================")
	fmt.Print("Nama pemesan: ")
	scanner.Scan()
	inputNamaPemesan := scanner.Text()

	fmt.Println()
	fmt.Println("Pilihan Paket Yang Tersedia:")
	for _, itemPaketUlangTahun := range listPaketUlangTahun {
		fmt.Printf("%d. %s (Rp.%d)\n", itemPaketUlangTahun.kode, itemPaketUlangTahun.nama, itemPaketUlangTahun.harga)
	}
	fmt.Print("Pilih paket [1/2/3]: ")
	scanner.Scan()
	inputKodePaket := scanner.Text()
	kodePaket, err := strconv.Atoi(inputKodePaket)
	if err != nil {
		fmt.Println("Pilihan paket tidak valid. Error:", err.Error())
		return
	}
	var selectedPaketUlangTahun paketUlangTahun
	for _, itemPaketUlangTahun := range listPaketUlangTahun {
		if itemPaketUlangTahun.kode == kodePaket {
			selectedPaketUlangTahun = itemPaketUlangTahun
			break
		}
	}
	if selectedPaketUlangTahun.nama == "" {
		fmt.Println("Paket ulang tahun yang dipilih tidak ditemukan.")
		return
	}

	fmt.Println()
	fmt.Println("Pilihan Atraksi Tambahan:")
	for _, itemAtraksi := range listAtraksiUlangTahun {
		fmt.Printf("%d. %s (Rp.%d)\n", itemAtraksi.kode, itemAtraksi.nama, itemAtraksi.harga)
	}
	fmt.Print("Pilih atraksi [1/2/3]: ")
	scanner.Scan()
	inputKodeAtraksi := scanner.Text()
	kodeAtraksi, err := strconv.Atoi(inputKodeAtraksi)
	if err != nil {
		fmt.Println("Pilihan atraksi tidak valid. Error:", err.Error())
		return
	}
	var selectedAtraksi atraksiUlangTahun
	for _, itemAtraksi := range listAtraksiUlangTahun {
		if itemAtraksi.kode == kodeAtraksi {
			selectedAtraksi = itemAtraksi
			break
		}
	}
	if selectedAtraksi.nama == "" {
		fmt.Println("Atraksi yang dipilih tidak ditemukan.")
		return
	}

	fmt.Println()
	fmt.Println("Detail Biaya Paket Ulang Tahun")
	fmt.Printf("Nama pemesan: %s\n", inputNamaPemesan)
	fmt.Println("==============================")
	fmt.Printf("Paket ulang tahun yang dipilih\t: %d\n", selectedPaketUlangTahun.kode)
	fmt.Printf("Nama paket ulang tahun\t\t: %s\n", selectedPaketUlangTahun.nama)
	fmt.Printf("Harga paket ulang tahun\t\t: Rp.%d\n", selectedPaketUlangTahun.harga)
	fmt.Printf("Atraksi tambahan\t\t: %s\n", selectedAtraksi.nama)
	fmt.Printf("Harga atraksi\t\t\t: Rp.%d\n", selectedAtraksi.harga)
	fmt.Println("==============================")

	totalHarga := selectedPaketUlangTahun.harga + selectedAtraksi.harga
	potonganHarga := 0
	if totalHarga >= 2000000 {
		fmt.Println("Bonus => Black Forest")
		potonganHarga = int(float64(totalHarga) * 0.1)
	} else {
		fmt.Println("Maaf ya tidak dapat bonus")
	}

	fmt.Println("==============================")
	fmt.Printf("Total harga\t\t\t: Rp.%d\n", totalHarga)
	fmt.Printf("Potongan harga\t\t\t: Rp.%d\n", potonganHarga)

	totalYangPerluDibayar := totalHarga - potonganHarga
	fmt.Printf("Total yang perlu dibayar\t: Rp.%d\n", totalYangPerluDibayar)
	fmt.Print("Uang bayar\t\t\t: Rp.")
	scanner.Scan()
	inputUangBayar := scanner.Text()
	uangBayar, err := strconv.Atoi(inputUangBayar)
	if err != nil {
		fmt.Println("Input uang bayar tidak valid. Error:", err.Error())
		return
	} else if uangBayar < totalYangPerluDibayar {
		fmt.Println("Uang pembayaran tidak cukup nih.")
		return
	}

	uangKembalian := uangBayar - totalYangPerluDibayar
	fmt.Printf("Uang kembali\t\t\t: Rp.%d\n", uangKembalian)
	fmt.Println("==============================")
}

type paketUlangTahun struct {
	kode  int
	nama  string
	harga int
}

type atraksiUlangTahun struct {
	kode  int
	nama  string
	harga int
}
