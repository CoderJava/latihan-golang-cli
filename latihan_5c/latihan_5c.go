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
	inputAngka := scanner.Text()
	angka, err := strconv.Atoi(inputAngka)
	if err != nil {
		fmt.Println("Inputan angka tidak valid: Error:", err.Error())
		return
	} else if angka <= 0 {
		fmt.Println("Inputan angka minimal 1 ya.")
		return
	}

	total, counter, index := 0, 0, 0
	for counter < angka {
		index += 1
		if index%2 != 0 {
			continue
		}
		total += index
		counter += 1
		if isLastIndex := counter == angka; isLastIndex {
			fmt.Printf("%d=%d", index, total)
		} else {
			fmt.Printf("%d+", index)
		}
	}
}
