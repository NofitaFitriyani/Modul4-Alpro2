package main

import (
	"fmt"
)

func faktorial(n int) int {
	if n == 0 {
		return 1
	}
	fact := 1
	for i := 1; i <= n; i++ {
		fact *= i
	}
	return fact
}

func permutasi(n int, r int) int {
	return faktorial(n) / faktorial(n-r)
}

func kombinasi(n int, r int) int {
	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}

func main() {
	var a, b, c, d int
	fmt.Println("Masukkan nilai a, b, c, d (dispasi): ")
	fmt.Scanf("%d %d %d %d", &a, &b, &c, &d)

	p_ac := permutasi(a, c)
	c_ac := kombinasi(a, c)

	p_bd := permutasi(b, d)
	c_bd := kombinasi(b, d)

	fmt.Println(p_ac, c_ac)
	fmt.Println(p_bd, c_bd)
}
