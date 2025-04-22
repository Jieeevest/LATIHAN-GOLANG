package main

import "fmt"

func main() {
	/* 
		Soal 1: Tampilkan Angka Genap
		Buat program yang mencetak angka genap dari 1 sampai 20.

		Output:
		2
		4
		6
		...
		20
	*/

	for i := 2; i <= 20; i += 2 {
		fmt.Println(i)
	}

	/* 
		Soal 2: Jumlahkan Semua Elemen Slice
		Diberi slice:

		numbers := []int{2, 5, 3, 7, 1}
		Hitung dan tampilkan jumlah semua angkanya.
		Output: Jumlah: 18
	*/

	numbers := []int{2, 5, 3, 7, 1}
	sum := 0
	for _, value := range numbers {
		sum += value
	}
	fmt.Println("Jumlah: ", sum)

	/* 
		Soal 3: Tebak Password (While-style)
		Tanya user untuk memasukkan password (misal: "golang123") sampai benar, baru keluar dari loop dan cetak “Akses diberikan”.
	*/

	password := "golang123"
	var input string
	for input != password {
		fmt.Print("Masukkan password: ")
		fmt.Scanln(&input)
	}
	fmt.Println("Akses diberikan")
}