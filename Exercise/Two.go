package main

import (
	"fmt"
	"time"
)

func main() {
	var date int
	var month int
	var year int
	fmt.Println("Masukkan tanggal ulang tahun Anda")
	fmt.Scanln(&date)

	fmt.Println("Masukkan bulan ulang tahun Anda")
	fmt.Scanln(&month)

	fmt.Println("Masukkan tahun ulang tahun Anda")
	fmt.Scanln(&year)

	var birthDate time.Time
	birthDate = time.Date(year, time.Month(month), date, 0, 0, 0, 0, time.Local)
	fmt.Println(birthDate)

	var name string
	fmt.Println("Masukkan nama anda:")
	fmt.Scanln(&name)

	fmt.Println("Menghitung umur", name, "...")
	now := time.Now()
	age := now.Year() - birthDate.Year()
	if now.YearDay() < birthDate.YearDay() {
		age--
	}

	fmt.Println("Umur anda adalah", age)
}
