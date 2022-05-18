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
		fmt.Println("Inputan angka minimal harus bernilai 1 ya.")
		return
	}

	printOutEvenNumber(number)
	fmt.Println()
	printOutOddNumber(number)
}

func printOutOddNumber(limitCounter int) {
	/**
	1=1
	1*3=3
	1*3*5=15
	1*3*5*7=105
	1*3*5*7*9=945
	*/
	number := limitCounter * 2
	for x := 1; x < number; x += 2 {
		total := 1
		for y := 1; y <= x; y += 2 {
			total *= y
			// Pengkondisian dibawah ini bisa dibuat menjadi seperti ini juga
			// if y == x {
			// ...
			// } else {
			// ...
			// }
			if isLastIndex := y == x; isLastIndex {
				fmt.Printf("%d=%d", y, total)
			} else {
				fmt.Printf("%d*", y)
			}
		}
		fmt.Println()
	}
}

func printOutEvenNumber(limitCounter int) {
	/**
	10+8+6+4+2=30
	10+8+6+4=28
	10+8+6=24
	10+8=18
	10=10
	*/
	number := 2 * limitCounter
	for x := 2; x <= number; x += 2 {
		total := 0
		for y := number; y >= x; y -= 2 {
			total += y
			if isLastIndex := y == x; isLastIndex {
				fmt.Printf("%d=%d", y, total)
			} else {
				fmt.Printf("%d+", y)
			}
		}
		fmt.Println()
	}
}
