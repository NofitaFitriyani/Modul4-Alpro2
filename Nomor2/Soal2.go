package main

import "fmt"

func hitungSkor(soal [8]int, skor *int, waktu *int) {
	*skor = 0
	*waktu = 0

	for i := 0; i < 8; i++ {
		if soal[i] < 301 {
			*skor++
			*waktu += soal[i]
		}
	}
}

func main() {
	nama1 := "Astuti"
	soal1 := [8]int{20, 50, 301, 301, 61, 71, 75, 10}

	nama2 := "Bertha"
	soal2 := [8]int{25, 47, 301, 26, 50, 60, 65, 21}

	var skor1, waktu1, skor2, waktu2 int

	hitungSkor(soal1, &skor1, &waktu1)
	hitungSkor(soal2, &skor2, &waktu2)

	pemenang := ""
	if skor2 > skor1 || (skor2 == skor1 && waktu2 < waktu1) {
		pemenang = nama2
	} else {
		pemenang = nama1
	}

	fmt.Printf("%s %d %d\n", nama1, skor1, waktu1)
	fmt.Printf("%s %d %d\n", nama2, skor2, waktu2)
	fmt.Printf("Pemenang: %s\n", pemenang)
}
