package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("                Studio Al Izzah                 \n")
	fmt.Print("================================================\n")
	fmt.Print("Isi Daftar Penyewa Berikut\n")
	fmt.Print("Nama Group\t\t: ")
	scanner.Scan()

	fmt.Print("Alamat\t\t\t: ")
	scanner.Scan()

	fmt.Print("No.Telp\t\t\t: ")
	scanner.Scan()

	fmt.Print("================================================\n")
	fmt.Print("Lama Rental (jam)\t: ")
	scanner.Scan()
	strLamaRental := scanner.Text()
	lamaRental, err := strconv.Atoi(strLamaRental)
	if err != nil {
		fmt.Println("Inputan lama rental tidak valid.")
		return
	} else if lamaRental < 1 {
		fmt.Println("Inputan lama rental harus minimal 1 jam ya.")
		return
	}

	const tarifPerJam = 200000
	var totalBayar int
	if lamaRental == 1 {
		totalBayar = tarifPerJam
	} else {
		biayaTambahan := (lamaRental - 1) * tarifPerJam
		biayaTambahan = int(float64(biayaTambahan) * 0.25)
		totalBayar = tarifPerJam + biayaTambahan
	}
	fmt.Print("Total Bayar\t\t: Rp.", totalBayar)
}
