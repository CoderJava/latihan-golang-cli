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

	const kodeKeretaB = "B"
	const kodeKeretaL = "L"
	const kodeKeretaP = "P"

	const namaKeretaB = "Argo Bromo"
	const namaKeretaL = "Argo Lawu"
	const namaKeretaP = "Parahyangan"

	const kodeKelas1 = "1"
	const kodeKelas2 = "2"
	const kodeKelas3 = "3"

	const namaKelas1 = "Eksekutif"
	const namaKelas2 = "Bisnis"
	const namaKelas3 = "Ekonomi"

	const hargaKelas1 = 100000
	const hargaKelas2 = 80000
	const hargaKelas3 = 50000

	const kodeJurusanA = "A"
	const kodeJurusanB = "B"
	const kodeJurusanC = "C"
	const kodeJurusanD = "D"

	const namaJurusanA = "JKT-SBY"
	const namaJurusanB = "JKT-SMG"
	const namaJurusanC = "JKT-BND"
	const namaJurusanD = "JKT-BALI"

	const hargaJurusanA = 40000
	const hargaJurusanB = 30000
	const hargaJurusanC = 20000
	const hargaJurusanD = 50000

	fmt.Print("Input jumlah data: ")
	scanner.Scan()
	inputJumlahData := scanner.Text()
	jumlahData, err := strconv.Atoi(inputJumlahData)
	if err != nil {
		fmt.Println("Input jumlah data tidak valid. Error:", err.Error())
		return
	} else if jumlahData <= 0 {
		fmt.Println("Input jumlah data minimal harus bernilai 1 ya.")
		return
	}

	listDataKereta := make([]Kereta, jumlahData)
	for index := 0; index < jumlahData; index++ {
		fmt.Println("===========================")
		fmt.Printf("Data ke-%d\n", index+1)
		fmt.Print("Kode kereta [B/L/P]: ")
		scanner.Scan()
		inputKodeKereta := strings.ToUpper(scanner.Text())
		var namaKereta string
		switch inputKodeKereta {
		case kodeKeretaB:
			namaKereta = namaKeretaB
		case kodeKeretaL:
			namaKereta = namaKeretaL
		case kodeKeretaP:
			namaKereta = namaKeretaP
		default:
			fmt.Println("Input kode kereta tidiak valid.")
			return
		}

		fmt.Println("\nKelas pilihan:")
		fmt.Println("1. Eksekutif")
		fmt.Println("2. Bisnis")
		fmt.Println("3. Ekonomi")
		fmt.Print("Kode kelas [1/2/3]: ")
		scanner.Scan()
		inputKelas := scanner.Text()
		var kelas string
		var hargaKelas int
		switch inputKelas {
		case kodeKelas1:
			kelas = namaKelas1
			hargaKelas = hargaKelas1
		case kodeKelas2:
			kelas = namaKelas2
			hargaKelas = hargaKelas2
		case kodeKelas3:
			kelas = namaKelas3
			hargaKelas = hargaKelas3
		default:
			fmt.Println("Input kode kelas tidak valid.")
			return
		}

		fmt.Println("\nJurusan:")
		fmt.Println("A. JKT-SBY")
		fmt.Println("B. JKT-SMG")
		fmt.Println("C. JKT-BND")
		fmt.Println("D. JKT-BALI")
		fmt.Print("Tujuan Anda [A/B/C/D]: ")
		scanner.Scan()
		inputJurusan := strings.ToUpper(scanner.Text())
		var jurusan string
		var hargaJurusan int
		switch inputJurusan {
		case kodeJurusanA:
			jurusan = namaJurusanA
			hargaJurusan = hargaJurusanA
		case kodeJurusanB:
			jurusan = namaJurusanB
			hargaJurusan = hargaJurusanB
		case kodeJurusanC:
			jurusan = namaJurusanC
			hargaJurusan = hargaJurusanC
		case kodeJurusanD:
			jurusan = namaJurusanD
			hargaJurusan = hargaJurusanD
		default:
			fmt.Println("Input jurusan tidak valid.")
			return
		}

		fmt.Print("Jumlah beli tiket: ")
		scanner.Scan()
		inputJumlahBeliTiket := scanner.Text()
		jumlahBeliTiket, err := strconv.Atoi(inputJumlahBeliTiket)
		if err != nil {
			fmt.Println("Input jumlah beli tiket tidak valid. Error:", err.Error())
			return
		} else if jumlahBeliTiket <= 0 {
			fmt.Println("Input jumlah beli tiket minimal harus 1 ya.")
			return
		}

		harga := hargaKelas + hargaJurusan
		totalHarga := harga * jumlahBeliTiket
		diskon := 0
		if jumlahBeliTiket >= 10 {
			diskon = int(float64(totalHarga) * 0.25)
		} else if jumlahBeliTiket >= 5 {
			diskon = int(float64(totalHarga) * 0.10)
		}
		bayar := totalHarga - diskon

		kereta := Kereta{
			kode:       inputKodeKereta,
			nama:       namaKereta,
			kelas:      kelas,
			jurusan:    jurusan,
			harga:      harga,
			jumlahBeli: jumlahBeliTiket,
			totalHarga: totalHarga,
			diskon:     diskon,
			bayar:      bayar,
		}
		listDataKereta[index] = kereta
	}

	totalBayar := 0
	fmt.Println()
	fmt.Println("===========================")
	fmt.Println("BIRO TRAVEL INSAN SEJAHTERA")
	fmt.Println("===========================")
	for index, kereta := range listDataKereta {
		fmt.Printf("No\t\t: %d\n", index+1)
		fmt.Printf("Kode kereta\t: %s\n", kereta.kode)
		fmt.Printf("Nama kereta\t: %s\n", kereta.nama)
		fmt.Printf("Kelas\t\t: %s\n", kereta.kelas)
		fmt.Printf("Jurusan\t\t: %s\n", kereta.jurusan)
		fmt.Printf("Harga tiket\t: Rp.%d\n", kereta.harga)
		fmt.Printf("Jumlah tiket\t: %d\n", kereta.jumlahBeli)
		fmt.Printf("Total harga\t: Rp.%d\n", kereta.totalHarga)
		fmt.Printf("Diskon\t\t: Rp.%d\n", kereta.diskon)
		fmt.Printf("Bayar\t\t: Rp.%d\n", kereta.bayar)
		totalBayar += kereta.bayar
		if index != len(listDataKereta)-1 {
			fmt.Println()
		}
	}
	fmt.Println("===========================")
	fmt.Printf("Total bayar\t: %d", totalBayar)
}

type Kereta struct {
	kode       string
	nama       string
	kelas      string
	jurusan    string
	harga      int
	jumlahBeli int
	totalHarga int
	diskon     int
	bayar      int
}
