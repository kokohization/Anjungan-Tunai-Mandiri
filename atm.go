package main

import "fmt"

//deklarasi variable buat kredensial
var username string = "fazle"
var password string = "2406424335"

//deklarasi variable buat array saldo and riwayat
var saldo = [1]int{3500000}
var riwayat = []string{}

//function main
func main() {

	var pilihan, jumlah int
	var inputusername, inputpassword string

	fmt.Println(".______________________________________________________.")
	fmt.Println("-----------Selamat Datang di Anjungan Tunai!------------")
	fmt.Println(".______________________________________________________.")

	//ini buat password & validasi kredensial
	for {
		fmt.Print("Masukkan Username anda-> ")
		fmt.Scan(&inputusername)
		fmt.Print("Masukkan Password anda-> ")
		fmt.Scan(&inputpassword)

		if !validasiakun(inputusername, inputpassword) {
			fmt.Println("Username atau Password salah!")
			fmt.Println(".______________________________________________________.")
			continue

		} else {
			fmt.Println("Username & Password Benar, Selamat Bertransaksi!")

			break
		}

	}

	fmt.Println("--------------------------------------------------------")

	//menu transaksi di atm
	for {

		fmt.Println(".________________________Menu__________________________.")
		fmt.Println("Pilih Menu (ketikkan angka)")
		fmt.Println("1. Lihat Informasi Akun")
		fmt.Println("2. Lihat Saldo")
		fmt.Println("3. Deposit")
		fmt.Println("4. Tarik Saldo")
		fmt.Println("5. Riwayat Transaksi")
		fmt.Println("6. Keluar")
		fmt.Print("Pilih Opsi-> ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			lihatinformasi()
			fmt.Println(".______________________________________________________.")
		case 2:
			lihatsaldo()
			fmt.Println(".______________________________________________________.")
		case 3:
			tambahsaldo(jumlah)
			fmt.Println(".______________________________________________________.")
		case 4:
			tariksaldo(jumlah)
			fmt.Println(".______________________________________________________.")
		case 5:
			lihatriwayat()
			fmt.Println(".______________________________________________________.")
		case 6:
			fmt.Println("Terimakasih Telah Menggunakan Anjungan Tunai", username)
			fmt.Println(".________________________xoxo__________________________.")
			return
		default:
			fmt.Println("Input Tidak Valid, Silahkan Coba Lagi!")
			fmt.Println(".______________________________________________________.")

		}
	}
}

//func buat validasi kredensial
func validasiakun(inputusername, inputpassword string) bool {
	return inputusername == username && inputpassword == password
}

//func buat lihat saldo
func lihatsaldo() {
	fmt.Println("._______________________Saldo__________________________.")
	fmt.Printf("Saldo akun %s: Rp%d\n", username, saldo[0])

}

//func buat lihat informasi akun
func lihatinformasi() {
	fmt.Println(".__________Ini Adalah Informasi Akun Anda!______________")
	fmt.Println("Username:", username)
	fmt.Println("Password:", password)

}

//func buat deposit
func tambahsaldo(jumlah int) {

	fmt.Println(".___________________Deposit Saldo______________________.")

	for {
		fmt.Print("Masukkan Jumlah Deposit: ")
		fmt.Scan(&jumlah)

		if jumlah > 0 {
			saldo[0] += jumlah
			riwayat = append(riwayat, fmt.Sprintf("Deposit Rp%d", jumlah))
			fmt.Printf("Berhasil Deposit Sejumlah Rp%d ke Akun %s", jumlah, username)

			break
		} else {
			fmt.Println("Jumlah Harus Lebih Dari 0!")
			fmt.Println(".______________________________________________________.")
			continue
		}
	}
}

//func buat lihat riwayat transaksi
func lihatriwayat() {
	fmt.Println("._____________________Riwayat__________________________.")
	fmt.Println("Riwayat Transaksi->")
	if len(riwayat) == 0 {
		fmt.Println("Belum Ada Riwayat Transaksi")
	} else {
		for _, transaksi := range riwayat {
			fmt.Println(transaksi)
		}
	}
}

//func buat penarikan saldo
func tariksaldo(jumlah int) {

	fmt.Println(".___________________Tarik Saldo________________________.")

	for {
		fmt.Println("Masukkan Jumlah Yang Akan Ditarik: ")
		fmt.Scan(&jumlah)

		if jumlah <= 0 {
			fmt.Println("Jumlah Harus Lebih Besar Dari 0")
			fmt.Println(".______________________________________________________.")
			continue
		}

		if saldo[0] < jumlah {
			fmt.Println("Saldo Anda Tidak Cukup!")
			fmt.Println(".______________________________________________________.")
			continue
		}

		saldo[0] -= jumlah
		riwayat = append(riwayat, fmt.Sprintf("Penarikan Rp%d", jumlah))
		fmt.Printf("Berhasil Penarikan Sejumlah Rp%d Dari Akun %s", jumlah, username)

		break
	}
}
