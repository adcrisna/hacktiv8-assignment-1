package main

import (
	"fmt"
	"os"
	"strconv"
)

type Mahasiswa struct {
	NoAbsen   int
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

var mhs = []Mahasiswa{
	{NoAbsen: 1, Nama: "Muhamad Ade Crisna", Alamat: "Cirebon", Pekerjaan: "Back-End Engineer I", Alasan: "Begitu Syulitt"},
	{NoAbsen: 2, Nama: "Nurul R", Alamat: "Jakarta", Pekerjaan: "FrontEnd", Alasan: "Belajar Golang"},
	{NoAbsen: 3, Nama: "Nadya P", Alamat: "Bekasi", Pekerjaan: "UI/UX", Alasan: "Apa Lagi Golang"},
	{NoAbsen: 4, Nama: "Putri G", Alamat: "Surabaya", Pekerjaan: "Quality Assurance", Alasan: "Iyakah Dik?"},
	{NoAbsen: 5, Nama: "Reyhan Begitu Syulitt", Alamat: "Jakarta", Pekerjaan: "System Analist", Alasan: "Ya Ndak Tau Ko Tanya Saya"},
}

func (s Mahasiswa) printData() {
	fmt.Println("Nama :", s.Nama)
	fmt.Println("Alamat :", s.Alamat)
	fmt.Println("Pekerjaan :", s.Pekerjaan)
	fmt.Println("Alasan :", s.Alasan)
}

func cariMahasiswa(noAbsen string) {

	var temp int

	absen, _ := strconv.Atoi(noAbsen)

	for index, value := range mhs {
		if absen == value.NoAbsen {
			temp = index + 1
		}
	}

	if temp > 0 {
		mhs[temp-1].printData()
	} else {
		fmt.Println("Data tidak ada")
	}
}

func main() {

	input := os.Args[1]

	cariMahasiswa(input)
}
