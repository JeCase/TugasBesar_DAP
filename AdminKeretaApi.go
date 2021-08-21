package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

var clear map[string]func()

const N int = 100

type kereta struct {
	kode, nama_kereta, sts_awal, sts_akhir, jam string
	kursi, harga                                int
}

type daf [N]kereta

func main() {
	var (
		x         string
		jum_train int
		train     daf
	)

	for x != "7" {

		fmt.Println(":=======================================:")
		fmt.Println("|  APLIKASI PEMESANAN TIKET KERETA API  |")
		fmt.Println(":=======================================:")
		fmt.Println()
		fmt.Println("1. Petugas")
		fmt.Println("2. Pemesan")
		fmt.Println("3. Data Kereta")
		fmt.Println("4. Tiket Promo Rp.50000")
		fmt.Println("5. Jumlah Kursi Terisi")
		fmt.Println("6. Delete Data Kereta")
		fmt.Println("7. Keluar")
		fmt.Println()
		fmt.Print("Ketik Angka Yang dipilih : ")
		fmt.Scanln(&x)

		fmt.Println()

		if x == "1" {
			CallClear()
			create(&train, &jum_train)
		} else if x == "2" {
			CallClear()
			order(&train, jum_train)
		} else if x == "3" {
			CallClear()
			printCreate(train, jum_train)
		} else if x == "4" {
			CallClear()
			gocap(train, jum_train)
		} else if x == "5" {
			CallClear()
			urut(&train, jum_train)
		} else if x == "6" {
			CallClear()
			delete(&train, &jum_train)
		} else if x == "7" {
			x = "7"
			CallClear()
		}
	}
}

func create(train *daf, jum_train *int) {
	var (
		kode, nama, awal, akhir, jam string
		kursi, harga                 int
	)

	i := *jum_train
	kode = "OKE"

	fmt.Println(":***************************************:")
	fmt.Println("|             #MENU PETUGAS#            |")
	fmt.Println(":***************************************:")
	fmt.Println()
	for i < N && (kode != "DONE") {
		fmt.Println("*****************************************")
		fmt.Println("Ketik 'DONE' Pada Masukan Kode Jika Sudah")
		fmt.Println("*****************************************")
		fmt.Println()
		fmt.Print("Masukan Kode Kereta 		: ")
		fmt.Scanln(&kode)

		if kode != "DONE" {
			fmt.Print("Masukan Nama Kereta 		: ")
			fmt.Scanln(&nama)
			fmt.Print("Masukan Stasiun Awal Kereta 	: ")
			fmt.Scanln(&awal)
			fmt.Print("Masukan Stasiun Akhir Kereta 	: ")
			fmt.Scanln(&akhir)
			fmt.Print("Masukan Jumlah Kursi (Maks 100)	: ")
			fmt.Scanln(&kursi)
			fmt.Print("Masukan Jam Keberangkatan 	: ")
			fmt.Scanln(&jam)
			fmt.Print("Masukan Harga Tiket 		: ")
			fmt.Scanln(&harga)
			fmt.Println()

			if kursi <= 100 {
				train[i].kode = kode
				train[i].nama_kereta = nama
				train[i].sts_awal = awal
				train[i].sts_akhir = akhir
				train[i].kursi = kursi
				train[i].jam = jam
				train[i].harga = harga
				i++
			} else {

			}
		}
		fmt.Println()
	}
	*jum_train = i
	CallClear()
}

func printCreate(train daf, jum_train int) {
	var n string

	fmt.Println(":***************************************:")
	fmt.Println("|             #DATA KERETA#             |")
	fmt.Println(":***************************************:")
	fmt.Println()
	for i := 0; i < jum_train; i++ {
		fmt.Println("Kode Kereta 			: ", train[i].kode)
		fmt.Println("Nama Kereta 			: ", train[i].nama_kereta)
		fmt.Println("Stasiun Awal 			: ", train[i].sts_awal)
		fmt.Println("Stasiun Akhir 			: ", train[i].sts_akhir)
		fmt.Println("Jumlah Kursi Terisi 		: ", train[i].kursi)
		fmt.Println("Jam Keberangkatan 		: ", train[i].jam)
		fmt.Println("Harga Tiket 			: ", train[i].harga)
		fmt.Println()
	}

	fmt.Print("Ketik '0' Untuk Kembali Ke Menu Utama : ")
	fmt.Scanln(&n)
	fmt.Println()
	if n == "0" {
		CallClear()
	} else {
		fmt.Print("Ketik '0' Untuk Kembali Ke Menu Utama : ")
		fmt.Scanln(&n)
	}
}

func order(train *daf, jum_train int) {
	var (
		awal, akhir string
		tiket       int
		found       bool
	)

	found = false
	fmt.Println(":***************************************:")
	fmt.Println("|           #MENU PEMESANAN#            |")
	fmt.Println(":***************************************:")
	fmt.Println()
	fmt.Print("Stasiun Awal : ")
	fmt.Scanln(&awal)
	fmt.Print("Stasiun Akhir : ")
	fmt.Scanln(&akhir)
	fmt.Print("Jumlah Tiket Yang Dipesan (Maks 10 Tiket) : ")
	fmt.Scanln(&tiket)

	if tiket <= 10 {
		search(train, jum_train, awal, akhir, tiket, &found)
	} else {
		fmt.Println("MAKSIMAL PEMESANAN TIKET ADALAH 10")
	}
	fmt.Println()
}

func search(train *daf, jum_train int, awal, akhir string, tiket int, found *bool) {
	var (
		n, jam string
	)

	for i := 0; i < jum_train; i++ {
		if awal == train[i].sts_awal && akhir == train[i].sts_akhir {
			if (train[i].kursi + tiket) <= 100 {
				fmt.Println()
				fmt.Println("*", "Ada kereta", train[i].nama_kereta, "Jam Keberangkatan : ", train[i].jam)
				*found = true
			} else if (train[i].kursi + tiket) != 0 {
				fmt.Println()
				fmt.Println("*", "Maaf Kursi kereta ", train[i].nama_kereta, " Tersisa ", (100 - train[i].kursi))
				*found = true
			} else {
				fmt.Println()
				fmt.Println("Maaf Kursi Kereta Jurusan Tersebut Telah Penuh.")
				fmt.Println()
			}
		}
	}
	if *found != true {
		fmt.Println("Maaf Kereta Tidak Ditemukan")
		fmt.Println()
	} else if *found == false {
		fmt.Println("Maaf Kereta Tidak Ditemukan")
		fmt.Println()
	}
	for i := 0; i < jum_train; i++ {
		if awal == train[i].sts_awal && akhir == train[i].sts_akhir {
			if (train[i].kursi + tiket) <= 100 {
				fmt.Print("Pilih jam keberangkatan : ")
				fmt.Scanln(&jam)
				for i := 0; i < jum_train; i++ {
					if awal == train[i].sts_awal && akhir == train[i].sts_akhir {
						if jam == train[i].jam {
							fmt.Println()
							fmt.Println("Apakah anda ingin memesan tiket kereta", train[i].nama_kereta, "dengan harga", diskon(*train, jum_train, jam, tiket))
							fmt.Println()
							OHYESOHNO(train, jum_train, awal, akhir, jam, tiket)
						} else {
							fmt.Println("Maaf Jam Tidak Sesuai")
						}
					}
				}
			}
		}
	}
	if n == "0" {
		CallClear()
	} else {
		fmt.Print("Ketik '0' Untuk Kembali Ke Menu Utama : ")
		fmt.Scanln(&n)
		CallClear()
	}
}

func diskon(train daf, jum_train int, jam string, tiket int) int {
	var (
		diskon, tot_bayar int
	)
	for i := 0; i < jum_train; i++ {
		if jam == "02.00" {
			diskon = ((train[i].harga * tiket) * 10) / 100
			tot_bayar = train[i].harga*tiket - diskon
		} else if jam == "14.00" {
			diskon = ((train[i].harga * tiket) * 5) / 100
			tot_bayar = train[i].harga*tiket - diskon
		} else {
			tot_bayar = train[i].harga * tiket
		}
	}
	return tot_bayar
}

func OHYESOHNO(train *daf, jum_train int, awal, akhir, jam string, tiket int) {
	var (
		x, pemesan string
		i          int
	)

	fmt.Println("1. YES ")
	fmt.Println("2. NO ")
	fmt.Println()
	fmt.Print("Masukkan angka : ")
	fmt.Scanln(&x)
	fmt.Println()
	if x == "1" {
		fmt.Print("Masukkan Nama Anda : ")
		fmt.Scanln(&pemesan)
		for i = 0; i < jum_train; i++ {
			if awal == train[i].sts_awal && akhir == train[i].sts_akhir {
				if jam == train[i].jam {
					train[i].kursi = train[i].kursi + tiket
					fmt.Println()
					fmt.Println("Nama Pemesan 				: ", pemesan)
					fmt.Println("Jumlah Tiket yang di Pesan 		: ", tiket)
					fmt.Println("Kode Kereta  				: ", train[i].kode)
					fmt.Println("Nama Kereta  				: ", train[i].nama_kereta)
					fmt.Println("Total Yang Harus di Bayar 		: ", diskon(*train, jum_train, jam, tiket))
					fmt.Println("Kursi Kereta", train[i].nama_kereta, "terisi : ", train[i].kursi)
					fmt.Println()
				}
			}
		}
	} else if x == "2" {
		order(&*train, jum_train)
	}
	fmt.Println()
}

func gocap(train daf, jum_train int) {
	var n string
	fmt.Println(":***************************************:")
	fmt.Println("|         Rp. 50000 Untuk Kamu          |")
	fmt.Println(":***************************************:")
	fmt.Println()
	for i := 0; i < jum_train; i++ {
		if train[i].harga == 50000 {
			fmt.Println("Kode Kereta : ", train[i].kode)
		}
	}
	fmt.Println()
	fmt.Print("Ketik '0' Untuk Kembali Ke Menu Utama : ")
	fmt.Scanln(&n)

	if n == "0" {
		CallClear()
	} else {
		fmt.Print("Ketik '0' Untuk Kembali Ke Menu Utama : ")
		fmt.Scanln(&n)
	}
	fmt.Println()
}

func urut(train *daf, jum_train int) {
	var min, pass, i int

	pass = 0
	for pass < jum_train-1 {
		min = pass
		i = pass + 1
		for i < jum_train {
			if train[min].kursi > train[i].kursi {
				min = i
			}
			i++
		}
		temp := train[min]
		train[min] = train[pass]
		train[pass] = temp
		pass++
	}
	printUrut(*train, jum_train)
}

func printUrut(train daf, jum_train int) {
	var n string

	fmt.Println(":***************************************:")
	fmt.Println("|           Kursi Tersedia             |")
	fmt.Println(":***************************************:")
	fmt.Println()
	for i := 0; i < jum_train; i++ {
		if train[i].kursi < 100 {
			fmt.Println("Kode Kereta : ", train[i].kode)
			fmt.Println("Nama Kereta : ", train[i].nama_kereta)
			fmt.Println("Stasiun Awal : ", train[i].sts_awal)
			fmt.Println("Stasiun Akhir : ", train[i].sts_akhir)
			fmt.Println("Jumlah Kursi Terisi : ", train[i].kursi)
			fmt.Println("Jam Keberangkatan : ", train[i].jam)
			fmt.Println("Harga Tiket : ", train[i].harga)
			fmt.Println()
		}
	}
	fmt.Print("Ketik '0' Untuk Kembali Ke Menu Utama : ")
	fmt.Scanln(&n)

	if n == "0" {
		CallClear()
	} else {
		fmt.Print("Ketik '0' Untuk Kembali Ke Menu Utama : ")
		fmt.Scanln(&n)
	}
	fmt.Println()
}

func delete(train *daf, jum_train *int) {
	var (
		kode        string
		i, jum_kode int
		found       bool
	)
	fmt.Println(":***************************************:")
	fmt.Println("|          Delete Data Kereta           |")
	fmt.Println(":***************************************:")
	fmt.Println()
	fmt.Print("Masukan kode kereta yang akan dihapus : ")
	fmt.Scanln(&kode)

	i = 0
	j := 0
	found = false

	for i = 0; i < *jum_train; i++ {
		if kode == train[i].kode {
			for j < *jum_train && !(found) {
				if kode == train[i].kode {
					i++
					found = true
				}
				j++
			}
			jum_kode = i - 1

			*jum_train--
			for jum_kode <= *jum_train {
				train[jum_kode] = train[jum_kode+1]
				jum_kode++

			}
		} else {

		}
	}
	CallClear()
}

func init() {
	clear = make(map[string]func()) //Initialize it

	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func CallClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

/*
NAMA OF ANGGOTA :
Jerry CS (1301194022)
Muhammad Zahran Zuan (1301190358)
*/
