package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

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

	fmt.Print("Input bulan: ")
	scanner.Scan()
	inputNamaBulan := scanner.Text()

	fmt.Print("Input tahun: ")
	scanner.Scan()
	inputTahun := scanner.Text()
	tahun, err := strconv.Atoi(inputTahun)
	if err != nil {
		fmt.Println("Input tahun tidak valid. Error:", err.Error())
		return
	} else if tahun <= 0 {
		fmt.Println("Input tahun minimal harus bernilai 1 ya.")
		return
	}

	listDataPasien := make([]pasien, jumlahData)
	for index := 0; index < jumlahData; index++ {
		fmt.Println("==========")
		fmt.Printf("Data ke-%d\n", index+1)

		fmt.Print("Nama pasien\t\t: ")
		scanner.Scan()
		inputNamaPasien := scanner.Text()

		fmt.Print("Jenis penyakit\t\t: ")
		scanner.Scan()
		inputJenisPenyakit := scanner.Text()

		fmt.Print("Lama perawatan (hari)\t: ")
		scanner.Scan()
		inputLamaPerawatan := scanner.Text()
		lamaPerawatan, err := strconv.Atoi(inputLamaPerawatan)
		if err != nil {
			fmt.Println("Input lama perawatan tidak valid. Error:", err.Error())
			return
		} else if lamaPerawatan <= 0 {
			fmt.Println("Input lama perawatan minimal harus 1 hari ya.")
			return
		}

		var biayaKamar int
		var biayaDokter int
		if lamaPerawatan <= 7 {
			biayaKamar = 150000
			biayaDokter = 300000
		} else if lamaPerawatan <= 15 {
			biayaKamar = 400000
			biayaDokter = 600000
		} else {
			biayaKamar = 700000
			biayaDokter = 1500000
		}
		total := biayaKamar + biayaDokter

		itemPasien := pasien{
			nama:          inputNamaPasien,
			jenisPenyakit: inputJenisPenyakit,
			lamaPerawatan: lamaPerawatan,
			biayaKamar:    biayaKamar,
			biayaDokter:   biayaDokter,
			total:         total,
		}
		listDataPasien[index] = itemPasien
	}
	fmt.Print("\n\n")
	fmt.Println("DATA PASIEN IBNU SINA HOSPITAL")
	fmt.Printf("Bulan\t: %s\n", inputNamaBulan)
	fmt.Printf("Tahun\t: %s\n", inputTahun)
	fmt.Printf("===============================\n")
	for index, itemPasien := range listDataPasien {
		fmt.Printf("No\t\t: %d\n", index+1)
		fmt.Printf("Nama\t\t: %s\n", itemPasien.nama)
		fmt.Printf("Jenis penyakit\t: %s\n", itemPasien.jenisPenyakit)
		fmt.Printf("Lama perawatan\t: %d hari\n", itemPasien.lamaPerawatan)
		fmt.Printf("Biaya kamar\t: Rp.%d\n", itemPasien.biayaKamar)
		fmt.Printf("Biaya dokter\t: Rp.%d\n", itemPasien.biayaDokter)
		fmt.Printf("Total\t\t: Rp.%d", itemPasien.total)
		fmt.Print("\n\n")
	}
}

type pasien struct {
	nama          string
	jenisPenyakit string
	lamaPerawatan int
	biayaKamar    int
	biayaDokter   int
	total         int
}
