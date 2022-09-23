package main

import (
	"fmt"
	"os"
)

type Biodata struct {
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

func main() {
	cariData()
}

func cariData() {
	data := os.Args[1]

	if data == "1" {
		b := Biodata{
			Nama:      "Muhamad Ade Crisna",
			Alamat:    "Kesambi, Kota Cirebon",
			Pekerjaan: "Mahasiswa",
			Alasan:    "Menambah Skill Dalam Menjadi Seorang Back-End Developer",
		}
		fmt.Println("Nama :", b.Nama)
		fmt.Println("Alamat :", b.Alamat)
		fmt.Println("Pekerjaan :", b.Pekerjaan)
		fmt.Println("Alasan :", b.Alasan)
	} else if data == "2" {
		b := Biodata{
			Nama:      "Arfan Maulana",
			Alamat:    "Kesambi, Kota Cirebon",
			Pekerjaan: "Staff Administrasi",
			Alasan:    "Ingin Menjadi Golang Developer",
		}
		fmt.Println("Nama :", b.Nama)
		fmt.Println("Alamat :", b.Alamat)
		fmt.Println("Pekerjaan :", b.Pekerjaan)
		fmt.Println("Alasan :", b.Alasan)
	} else if data == "3" {
		b := Biodata{
			Nama:      "Alex Martin",
			Alamat:    "Pekalipan, Kota Cirebon",
			Pekerjaan: "Staff Gudang",
			Alasan:    "Ingin Menjadi Back-end Developer",
		}
		fmt.Println("Nama :", b.Nama)
		fmt.Println("Alamat :", b.Alamat)
		fmt.Println("Pekerjaan :", b.Pekerjaan)
		fmt.Println("Alasan :", b.Alasan)
	} else if data == "4" {
		b := Biodata{
			Nama:      "Iqbal Pratama",
			Alamat:    "Ciledug, Kabupaten Cirebon",
			Pekerjaan: "Staff Gudang",
			Alasan:    "Ingin Menjadi Back-end Developer",
		}
		fmt.Println("Nama :", b.Nama)
		fmt.Println("Alamat :", b.Alamat)
		fmt.Println("Pekerjaan :", b.Pekerjaan)
		fmt.Println("Alasan :", b.Alasan)
	} else if data == "5" {
		b := Biodata{
			Nama:      "Selvi Felyanti",
			Alamat:    "Harjamukti, Kota Cirebon",
			Pekerjaan: "Staff Marketing",
			Alasan:    "Ingin Menjadi Back-end Developer",
		}
		fmt.Println("Nama :", b.Nama)
		fmt.Println("Alamat :", b.Alamat)
		fmt.Println("Pekerjaan :", b.Pekerjaan)
		fmt.Println("Alasan :", b.Alasan)
	} else if data == "6" {
		b := Biodata{
			Nama:      "Afifah Nur",
			Alamat:    "Harjamukti, Kota Cirebon",
			Pekerjaan: "Sales",
			Alasan:    "Ingin Menjadi Back-end Developer",
		}
		fmt.Println("Nama :", b.Nama)
		fmt.Println("Alamat :", b.Alamat)
		fmt.Println("Pekerjaan :", b.Pekerjaan)
		fmt.Println("Alasan :", b.Alasan)
	} else if data == "7" {
		b := Biodata{
			Nama:      "Nurul Rahayu",
			Alamat:    "Ciperna, Kabupaten Cirebon",
			Pekerjaan: "Mahasiswa",
			Alasan:    "Ingin Menjadi Back-end Developer",
		}
		fmt.Println("Nama :", b.Nama)
		fmt.Println("Alamat :", b.Alamat)
		fmt.Println("Pekerjaan :", b.Pekerjaan)
		fmt.Println("Alasan :", b.Alasan)
	} else if data == "8" {
		b := Biodata{
			Nama:      "Reyhan Maulana",
			Alamat:    "Harjamukti, Kota Cirebon",
			Pekerjaan: "Content Creater",
			Alasan:    "Ingin Menjadi Back-end Developer",
		}
		fmt.Println("Nama :", b.Nama)
		fmt.Println("Alamat :", b.Alamat)
		fmt.Println("Pekerjaan :", b.Pekerjaan)
		fmt.Println("Alasan :", b.Alasan)
	} else if data == "9" {
		b := Biodata{
			Nama:      "Rivaldo Kusuma",
			Alamat:    "Pekalipan Kota Cirebon",
			Pekerjaan: "PNS",
			Alasan:    "Ingin Menjadi Back-end Developer",
		}
		fmt.Println("Nama :", b.Nama)
		fmt.Println("Alamat :", b.Alamat)
		fmt.Println("Pekerjaan :", b.Pekerjaan)
		fmt.Println("Alasan :", b.Alasan)
	} else if data == "10" {
		b := Biodata{
			Nama:      "Wahyu Setiawan",
			Alamat:    "Perum, Kota Cirebon",
			Pekerjaan: "UI/UX Design",
			Alasan:    "Ingin Belajar Menjadi Back-end Developer",
		}
		fmt.Println("Nama :", b.Nama)
		fmt.Println("Alamat :", b.Alamat)
		fmt.Println("Pekerjaan :", b.Pekerjaan)
		fmt.Println("Alasan :", b.Alasan)
	} else {
		fmt.Println("Data Yang Anda Cari Tidak Ada!!")
	}
}
