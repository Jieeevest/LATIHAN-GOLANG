

package main

import (
	"fmt"
	"math/rand"
)


func main() {
	/*
	Buat program yang menerima angka dan mencetak:

	"Bilangan negatif" jika angka < 0

	"Nol" jika angka == 0

	"Bilangan positif" jika angka > 0

	*/
	angka := -1
	if angka < 0 {
		fmt.Println("Bilangan negatif")
	} else if angka == 0 {
		fmt.Println("Nol")
	} else {
		fmt.Println("Bilangan positif")
	}

	/*
	Buat program yang mengecek apakah seseorang bisa ikut ujian:

	Usia minimal 17 tahun

	Sudah bayar biaya ujian (sudahBayar := true/false)

	Tampilkan:

	"Boleh ikut ujian" jika dua-duanya memenuhi

	"Belum memenuhi syarat" jika salah satu tidak

	*/

	usia := rand.Intn(25)
	sudahBayar := true
	if usia > 17 && sudahBayar {
		fmt.Println("Boleh ikut ujian")
	} else {
		fmt.Println("Belum memenuhi syarat")
	}
}