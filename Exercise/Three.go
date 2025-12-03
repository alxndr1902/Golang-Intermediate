package main

import "fmt"

func main() {
	var hargaBarang int
	var diskon int
	var disc float64
	var totalHarga float64

	fmt.Println("Masukkan nilai harga barang:")
	fmt.Scanln(&hargaBarang)

	fmt.Println("Masukkan persentase diskon:")
	fmt.Scanln(&diskon)

	disc = float64(diskon) / 100
	totalHarga = float64(hargaBarang) * (1 - disc)

	fmt.Println("Harga barang adalah", hargaBarang)
	fmt.Println("Diskon sebesar", disc)
	fmt.Println("Total Harga barang adalah", totalHarga)
}
