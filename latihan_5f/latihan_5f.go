package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Masukkan angka: ")
	scanner.Scan()
	inputNumber := scanner.Text()
	number, err := strconv.Atoi(inputNumber)
	if err != nil {
		fmt.Println("Inputan angka tidak valid. Error:", err.Error())
		return
	} else if number <= 0 {
		fmt.Println("Inputan angka minimal harus bernilai 1 ya")
		return
	}

	fmt.Println()
	totalInColumn := 0
	for x := 1; x <= number; x++ {
		totalInRow := 0
		for y := 1; y <= number; y++ {
			if y <= x {
				totalInRow += y
				fmt.Printf("%d", y)
			} else {
				strY := strconv.Itoa(y)
				for index := 1; index <= len(strY); index++ {
					fmt.Print(" ")
				}
			}

			if y == number {
				fmt.Printf("\t=%d", totalInRow)
			}
		}
		totalInColumn += totalInRow
		fmt.Println()
	}
	fmt.Printf("Total=%d", totalInColumn)
}
