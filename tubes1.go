package main

import (
	"fmt"
	"strings"
)

const NMAX int = 10

type comment struct {
	pesan string
	star  float64
}

type tabCmnt [NMAX]comment

func main() {
	var data tabCmnt
	var nData int
	var x1 string
	fmt.Print("masukan jumlah komentar: ")
	fmt.Scan(&nData)
	fmt.Print("masukkan kata yang akan dicari: ")
	fmt.Scan(&x1)
	bacaKomentar(&data, nData)
	fmt.Println("=========================================================")
	cetakKomentar(data, nData)
	if sequentialSearch(data, nData, x1) {
		freq := frekuensiKata(data, nData, x1)
		fmt.Printf("ditemukan sebanyak %d kata %s", freq, x1)
	} else {
		fmt.Printf("kata %s tidak ditemukan", x1)
	}
	fmt.Println()
	fmt.Printf("hasil analisis dari komentar-komentar tersebut: %s ", menentukanNilaiProduk(data, nData, x1))
	fmt.Println()
	fmt.Println("=========================================================")
}

func bacaKomentar(A *tabCmnt, n int) {
	fmt.Println("Masukkan komentar dan bintang: ")
	for i := 0; i < n; i++ {
		fmt.Scan(&A[i].pesan)
		fmt.Scan(&A[i].star)
	}
}

func cetakKomentar(A tabCmnt, n int) {
	fmt.Printf("rata-rata bintang : %.1f ★", float32(rataArr(A, n)))
	fmt.Println()
	fmt.Println("komentar: ")
	for i := 0; i < n; i++ {
		fmt.Printf("%d. %s ", i+1, A[i].pesan)
		fmt.Println()
	}
}

func frekuensiKata(A tabCmnt, n int, x string) int {
	freq := 0
	for i := 0; i < n; i++ {
		if A[i].pesan == x {
			freq++
		}
	}
	return freq
}

func sequentialSearch(A tabCmnt, n int, x string) bool {
	var ketemu bool
	var i int
	ketemu = false
	i = 0

	for !ketemu && i < n {
		ketemu = A[i].pesan == x
		i++
	}
	return ketemu
}

func contains(s, substr string) bool {
	return strings.Contains(s, substr)
}

func analisisSentimen(A tabCmnt, n int, x string) bool {
	for i := 0; i < n; i++ {
		pesan := A[i].pesan

		if contains(pesan, "bagus") || contains(pesan, "keren") || contains(pesan, "mantap") || contains(pesan, "berfungsi") || contains(pesan, "lucu") {
			return true
		}

		if contains(pesan, "jelek") || contains(pesan, "kurang") || contains(pesan, "aneh") || contains(pesan, "cacat") {
			return false
		}
	}
	return true
}

func menentukanNilaiProduk(A tabCmnt, n int, x string) string {
	kesan := analisisSentimen(A, n, x)
	for i := 0; i < n; i++ {
		if kesan == true && A[i].star > 3 {
			return "positif"
		}
		if kesan == false && A[i].star <= 3 {
			return "negatif"
		}
	}
	return "netral"
}

func penjumlahanArr(A tabCmnt, n int) float64 {
	var jumlah float64
	jumlah = 0
	for i := 0; i < n; i++ {
		jumlah += A[i].star
	}
	return jumlah
}

func rataArr(A tabCmnt, n int) float64 {
	var rata float64
	rata = float64(penjumlahanArr(A, n)) / float64(n)
	return rata
}