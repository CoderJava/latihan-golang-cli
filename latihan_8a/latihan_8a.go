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

	const kodeKelasR = "R"
	const kodeKelasP = "P"
	const kodeKelasI = "I"

	const hargaKelasR = 800000
	const hargaKelasP = 1000000
	const hargaKelasI = 500000

	const diskonRanking1 = 0.5
	const diskonRanking2 = 0.2
	const diskonRanking3 = 0.1

	fmt.Println("Bentuk Data Masukan")
	fmt.Print("Jumlah Data: ")
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

	listDataSiswa := make([]Siswa, jumlahData)
	for index := 0; index < jumlahData; index++ {
		fmt.Printf("Data ke-%d\n", index+1)
		fmt.Print("Nama Siswa: ")
		scanner.Scan()
		inputNamaSiswa := scanner.Text()

		fmt.Printf("Kode Kelas[%s/%s/%s]: ", kodeKelasR, kodeKelasP, kodeKelasI)
		scanner.Scan()
		inputKodeKelas := scanner.Text()
		inputKodeKelas = strings.ToUpper(inputKodeKelas)
		if inputKodeKelas != kodeKelasR && inputKodeKelas != kodeKelasP && inputKodeKelas != kodeKelasI {
			fmt.Printf("Input kode kelas tidak valid. Harus bernilai %s atau %s atau %s.", kodeKelasR, kodeKelasP, kodeKelasI)
			return
		}

		fmt.Print("Ranking: ")
		scanner.Scan()
		inputRanking := scanner.Text()
		ranking, err := strconv.Atoi(inputRanking)
		if err != nil {
			fmt.Println("Input ranking tidak valid. Error:", err.Error())
			return
		} else if ranking <= 0 {
			fmt.Println("Input rangking tidak boleh bernilai minus atau nol.")
			return
		}

		var harga int
		switch inputKodeKelas {
		case kodeKelasR:
			harga = hargaKelasR
		case kodeKelasP:
			harga = hargaKelasP
		case kodeKelasI:
			harga = hargaKelasI
		}

		var potongan int
		switch ranking {
		case 1:
			potongan = int(float64(harga) * diskonRanking1)
		case 2:
			potongan = int(float64(harga) * diskonRanking2)
		case 3:
			potongan = int(float64(harga) * diskonRanking3)
		default:
			potongan = 0
		}

		siswa := Siswa{
			nama:      inputNamaSiswa,
			kelas:     inputKodeKelas,
			peringkat: ranking,
			harga:     harga,
			potongan:  potongan,
			total:     harga - potongan,
		}
		listDataSiswa[index] = siswa
		fmt.Println()
	}

	fmt.Println()
	fmt.Println("Bimbingan Belajar Bina Akhlak")
	fmt.Println("=============================")
	for index, siswa := range listDataSiswa {
		fmt.Printf("No\t\t: %d\n", index+1)
		fmt.Printf("Nama Siswa\t: %s\n", siswa.nama)
		fmt.Printf("Kelas\t\t: %s\n", siswa.kelas)
		fmt.Printf("Peringkat\t: %d\n", siswa.peringkat)
		fmt.Printf("Harga\t\t: Rp.%d\n", siswa.harga)
		fmt.Printf("Potongan\t: Rp.%d\n", siswa.potongan)
		fmt.Printf("Total\t\t: Rp.%d\n\n", siswa.total)
	}
}

type Siswa struct {
	nama      string
	kelas     string
	peringkat int
	harga     int
	potongan  int
	total     int
}
