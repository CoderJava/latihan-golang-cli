package main

import "fmt"

func main() {
	a, b := 22, 26
	fmt.Println("Nilai sebelum pemanggilan fungsi")
	fmt.Printf("a=%d b=%d\n\n", a, b)

	tukar(&a, &b)

	fmt.Println("Nilai setelah pemanggilan fungsi")
	fmt.Printf("a=%d b=%d", a, b)
}

func tukar(a *int, b *int) {
	// a, b = b, a // Ini yang ditukar ialah nilai pointer-nya bukan value-nya.
	*a, *b = *b, *a
	fmt.Println("Nilai di dalam fungsi tukar")
	fmt.Printf("a=%d b=%d\n\n", *a, *b)
}
