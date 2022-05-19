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

	fmt.Print("Masukkan jumlah data yang akan diproses: ")
	scanner.Scan()
	inputJumlahData := scanner.Text()
	jumlahData, err := strconv.Atoi(inputJumlahData)
	if err != nil {
		fmt.Println("Inputan jumlah data tidak valid. Error:", err.Error())
		return
	} else if jumlahData <= 0 {
		fmt.Println("Inputan jumlah data minimal harus bernilai 1 ya.")
		return
	}

	listData := make([]string, jumlahData)
	for index := 0; index < jumlahData; index++ {
		fmt.Printf("Data ke-%d\n", index+1)
		fmt.Print("Masukkan kode rumah piihan[A/B/C]: ")
		scanner.Scan()
		inputKodeRumah := scanner.Text()
		listData[index] = inputKodeRumah
		fmt.Println()
	}

	fmt.Println("DAFTAR HARGA RUMAH")
	fmt.Println("==================")
	for index, value := range listData {
		fmt.Printf("No\t\t\t: %d\n", index+1)
		fmt.Printf("Kode rumah\t\t: %s\n", strings.ToUpper(value))

		var tipeRumah string
		var uangMuka, bunga, sisaAngsuran, harga, total int
		switch strings.ToLower(value) {
		case "a":
			{
				tipeRumah = "RSS"
				uangMuka = 800000
				harga = 20000000
			}
		case "b":
			{
				tipeRumah = "RS"
				uangMuka = 100000
				harga = 25000000
			}
		case "c":
			{
				tipeRumah = "MEWAH"
				uangMuka = 12000000
				harga = 300000000
			}
		default:
			fmt.Print("Kode rumah tidak valid")
			checkIsLastIndex(index, len(listData)-1)
			continue
		}

		bunga = int(float64(harga) * 0.05)
		sisaAngsuran = harga - uangMuka
		total = harga + bunga

		fmt.Printf("Tipe rumah\t\t: %s\n", tipeRumah)
		fmt.Printf("Uang muka\t\t: Rp.%d\n", uangMuka)
		fmt.Printf("Bunga\t\t\t: Rp.%d\n", bunga)
		fmt.Printf("Sisa angsuran\t\t: Rp.%d\n", sisaAngsuran)
		fmt.Printf("Harga\t\t\t: Rp.%d\n", harga)
		fmt.Printf("Total Pembayaran\t: Rp.%d", total)
		checkIsLastIndex(index, len(listData)-1)
	}
}

func checkIsLastIndex(index int, length int) {
	if index != length {
		fmt.Print("\n\n")
	}
}
