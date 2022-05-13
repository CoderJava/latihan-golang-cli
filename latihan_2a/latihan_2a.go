package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("BIODATA PRIBADY SAYA")
	fmt.Println("====================")
	fmt.Print("Nama Lengkap\t: ")
	scanner.Scan()
	namaLengkap := scanner.Text()

	fmt.Print("Kota Kelahiran\t: ")
	scanner.Scan()
	kotaKelahiran := scanner.Text()

	fmt.Print("Tanggal Lahir\t: ")
	scanner.Scan()
	tanggalLahir := scanner.Text()

	fmt.Print("Alamat\t\t: ")
	scanner.Scan()
	alamat := scanner.Text()

	fmt.Print("No.Telp\t\t: ")
	scanner.Scan()
	noTelp := scanner.Text()

	fmt.Println()
	fmt.Printf("Perkenalkan nama saya %s. Saya lahir di %s. Saya lahir pada tanggal %s. Alamat saya %s. Teman-teman boleh menghubungi saya di No.Telp ini--> %s", namaLengkap, kotaKelahiran, tanggalLahir, alamat, noTelp)
	fmt.Print("\n\nSalam Persahabatan...")
}
