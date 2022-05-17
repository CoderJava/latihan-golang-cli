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
		fmt.Println("Inputan angka tidak valid. Error:", err.Error())
		return
	} else if angka <= 0 {
		fmt.Println("Inputan angka minimal 1 ya.")
		return
	}
	total := 0
	for i := 1; i <= angka; i++ {
		fmt.Print(i)
		total += i
	}
	fmt.Printf("=%d", total)
}
