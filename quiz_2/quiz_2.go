package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	listFilm := []film{
		{
			kode:      1,
			jenis:     "Indonesia",
			hargaSewa: 2000,
		},
		{
			kode:      2,
			jenis:     "Barat",
			hargaSewa: 3000,
		},
		{
			kode:      3,
			jenis:     "Korea",
			hargaSewa: 2500,
		},
	}
	listKategoriPenyewa := []kategoriPenyewa{
		{
			kode:   "P",
			nama:   "Pelanggan",
			diskon: 0.02,
		},
		{
			kode:   "U",
			nama:   "Umum",
			diskon: 0,
		},
	}

	// Input jumlah data
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

	// Input nama peminjam
	fmt.Print("Nama peminjam\t: ")
	scanner.Scan()
	namaPeminjam := scanner.Text()

	// Input alamat
	fmt.Print("Alamat\t\t: ")
	scanner.Scan()
	alamat := scanner.Text()

	listRentalFilm := make([]rentalFilm, jumlahData)
	for index := 0; index < jumlahData; index++ {
		// Input lama pinjam
		fmt.Println()
		fmt.Print("Lama pinjam (hari)\t: ")
		scanner.Scan()
		inputLamaPinjam := scanner.Text()
		lamaPinjam, err := strconv.Atoi(inputLamaPinjam)
		if err != nil {
			fmt.Println("Input lama pinjam tidak valid. Error:", err.Error())
			return
		} else if lamaPinjam <= 0 {
			fmt.Println("Input lama pinjam minimal harus 1 hari ya.")
			return
		}

		// Tampilkan list kode film yang tersedia
		strKodeFilm := ""
		fmt.Print("Kode film yang tersedia\n")
		lastItemFilm := listFilm[len(listFilm)-1]
		for _, itemFilm := range listFilm {
			kodeFilm := itemFilm.kode
			strKodeFilm += strconv.Itoa(kodeFilm)
			fmt.Printf("%d. %s (Rp.%d)\n", kodeFilm, itemFilm.jenis, itemFilm.hargaSewa)
			if itemFilm != lastItemFilm {
				strKodeFilm += "/"
			}
		}
		fmt.Println()

		// Input kode film
		fmt.Printf("Pilih kode film [%s]: ", strKodeFilm)
		scanner.Scan()
		inputKodeFilm := scanner.Text()
		selectedKodeFilm, err := strconv.Atoi(inputKodeFilm)
		if err != nil {
			fmt.Println("Input kode film tidak valid. Error:", err.Error())
			return
		}
		var selectedItemFilm film
		for _, itemFilm := range listFilm {
			if selectedKodeFilm == itemFilm.kode {
				selectedItemFilm = itemFilm
				break
			}
		}
		if selectedItemFilm.kode == 0 {
			fmt.Println("Kode film yang dipilih tidak ditemukan.")
			return
		}

		// Tampilkan list kategori penyewa yang tersedia
		strKodePenyewa := ""
		lastItemKategoriPenyewa := listKategoriPenyewa[len(listKategoriPenyewa)-1]
		fmt.Print("Kode penyewa yang tersedia\n")
		for _, itemKategoriPenyewa := range listKategoriPenyewa {
			kodePenyewa := itemKategoriPenyewa.kode
			strKodePenyewa += kodePenyewa
			fmt.Printf("%s. %s (%.1f%%)\n", kodePenyewa, itemKategoriPenyewa.nama, itemKategoriPenyewa.diskon*100)
			if kodePenyewa != lastItemKategoriPenyewa.kode {
				strKodePenyewa += "/"
			}
		}

		// Input kode penyewa
		fmt.Printf("Pilih kode penyewa [%s]: ", strKodePenyewa)
		scanner.Scan()
		inputKodePenyewa := scanner.Text()
		var selectedKategoriPenyewa kategoriPenyewa
		for _, itemKategoriPenyewa := range listKategoriPenyewa {
			if inputKodePenyewa == itemKategoriPenyewa.kode {
				selectedKategoriPenyewa = itemKategoriPenyewa
				break
			}
		}
		if selectedKategoriPenyewa.kode == "" {
			fmt.Println("Kode penyewa yang dipilih tidak ditemukan.")
			return
		}

		// Hitung total
		hargaSewa := selectedItemFilm.hargaSewa
		total := lamaPinjam * hargaSewa

		// Hitung denda
		denda := 0
		if lamaPinjam > 15 {
			switch selectedKodeFilm {
			case 1:
				denda = (lamaPinjam - 15) * 1500
			case 2:
				denda = (lamaPinjam - 15) * 2500
			default:
				denda = (lamaPinjam - 15) * 2500
			}
		}

		// Hitung diskon
		diskon := int(float64(total) * selectedKategoriPenyewa.diskon)

		// Hitung bayar
		bayar := total + denda - diskon

		itemRentalFilm := rentalFilm{
			kode:            selectedKodeFilm,
			jenis:           selectedItemFilm.jenis,
			kategoriPenyewa: selectedKategoriPenyewa.nama,
			hargaSewa:       hargaSewa,
			lamaPinjam:      lamaPinjam,
			total:           total,
			denda:           denda,
			diskon:          diskon,
			bayar:           bayar,
		}
		listRentalFilm[index] = itemRentalFilm
	}

	fmt.Println()
	fmt.Println("QATRUNADA CLUB")
	fmt.Println("==============")
	fmt.Printf("Nama peminjam\t\t: %s\n", namaPeminjam)
	fmt.Printf("Alamat\t\t\t: %s", alamat)
	fmt.Println()

	for index, itemRentalFilm := range listRentalFilm {
		fmt.Println()
		fmt.Printf("No\t\t\t: %d\n", index+1)
		fmt.Printf("Kode film\t\t: %d\n", itemRentalFilm.kode)
		fmt.Printf("Jenis film\t\t: %s\n", itemRentalFilm.jenis)
		fmt.Printf("Kategori penyewa\t: %s\n", itemRentalFilm.kategoriPenyewa)
		fmt.Printf("Harga sewa per hari\t: Rp.%d\n", itemRentalFilm.hargaSewa)
		fmt.Printf("Lama pinjam\t\t: %d hari\n", itemRentalFilm.lamaPinjam)
		fmt.Printf("Total\t\t\t: Rp.%d\n", itemRentalFilm.total)
		fmt.Printf("Denda\t\t\t: Rp.%d\n", itemRentalFilm.denda)
		fmt.Printf("Diskon\t\t\t: Rp.%d\n", itemRentalFilm.diskon)
		fmt.Printf("Bayar\t\t\t: Rp.%d\n", itemRentalFilm.bayar)
	}
}

type rentalFilm struct {
	kode            int
	jenis           string
	kategoriPenyewa string
	hargaSewa       int
	lamaPinjam      int
	total           int
	denda           int
	diskon          int
	bayar           int
}

type film struct {
	kode      int
	jenis     string
	hargaSewa int
}

type kategoriPenyewa struct {
	kode   string
	nama   string
	diskon float64
}
