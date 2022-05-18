package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Masukkan nilai baris: ")
	scanner.Scan()
	inputBaris := scanner.Text()
	baris, err := strconv.Atoi(inputBaris)
	if err != nil {
		fmt.Println("Input baris tidak valid. Error:", err.Error())
		return
	} else if baris <= 0 {
		fmt.Println("Inputan baris minimal bernilai 1 ya.")
		return
	}

	printOutIncrementStyle(baris)
	fmt.Println()
	printOutDecrementStyle(baris)
}

func printOutIncrementStyle(jumlahBaris int) {
	/**
	11111
	2222
	333
	44
	5
	*/

	for x := 1; x <= jumlahBaris; x++ {
		for y := jumlahBaris; y >= x; y-- {
			fmt.Printf("%d", x)
		}
		fmt.Println()
	}
}

func printOutDecrementStyle(jumlahBaris int) {
	/**
	55555
	4444
	333
	22
	1
	*/

	for x := jumlahBaris; x >= 1; x-- {
		for y := 1; y <= x; y++ {
			fmt.Printf("%d", x)
		}
		fmt.Println()
	}
}
