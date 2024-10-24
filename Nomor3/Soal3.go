package main

import "fmt"

func cetakDeret(n int) {
	fmt.Print(n)

	for n != 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
		fmt.Print(" ", n)
	}
	fmt.Println()
}

func main() {
	var n int

	for {
		fmt.Print("Masukkan bilangan: ")
		fmt.Scan(&n)

		if n > 0 && n < 1000000 {
			break
		} else {
			fmt.Println("Input tidak valid, coba lagi.")
		}
	}

	cetakDeret(n)
}
