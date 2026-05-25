package main

import "fmt"

type Mahasiswa struct {
	Nim   string
	Nama  string
	Kelas string
}

var dataMahasiswa [100]Mahasiswa

func main() {
	jumlahMahasiswa := 0
	selesai := false

	for !selesai {
		tampilkanMenuUtama()
		pilihan := bacaAngka("Pilih menu: ")

		switch pilihan {
		case 1:
			tambahMahasiswa(&dataMahasiswa, &jumlahMahasiswa)
		case 2:
			tampilkanMahasiswa(dataMahasiswa, jumlahMahasiswa)
		case 3:
			ubahMahasiswa(&dataMahasiswa, jumlahMahasiswa)
		case 4:
			hapusMahasiswa(&dataMahasiswa, &jumlahMahasiswa)
		case 0:
			fmt.Println("Program selesai.")
			selesai = true
		default:
			fmt.Println("Menu tidak tersedia.")
		}
	}
}

func tampilkanMenuUtama() {
	fmt.Println()
	fmt.Println("=== SiPresensi ===")
	fmt.Println("1. Tambah Mahasiswa")
	fmt.Println("2. Tampilkan Mahasiswa")
	fmt.Println("3. Ubah Mahasiswa")
	fmt.Println("4. Hapus Mahasiswa")
	fmt.Println("0. Keluar")
}

func tambahMahasiswa(daftar *[100]Mahasiswa, jumlah *int) {
	var mahasiswaBaru Mahasiswa

	if *jumlah >= 100 {
		fmt.Println("Data mahasiswa sudah penuh.")
	} else {
		mahasiswaBaru.Nim = bacaTeks("NIM: ")

		if cariMahasiswaSequential(*daftar, *jumlah, mahasiswaBaru.Nim) != -1 {
			fmt.Println("NIM sudah terdaftar.")
		} else {
			mahasiswaBaru.Nama = bacaTeks("Nama tanpa spasi: ")
			mahasiswaBaru.Kelas = bacaTeks("Kelas: ")

			daftar[*jumlah] = mahasiswaBaru
			*jumlah = *jumlah + 1

			urutkanNimInsertion(daftar, *jumlah)
			fmt.Println("Data mahasiswa berhasil ditambahkan.")
		}
	}
}

func tampilkanMahasiswa(daftar [100]Mahasiswa, jumlah int) {
	if jumlah == 0 {
		fmt.Println("Belum ada data mahasiswa.")
	} else {
		fmt.Println()
		fmt.Printf("%-14s %-25s %-14s\n", "NIM", "Nama", "Kelas")

		for i := 0; i < jumlah; i++ {
			fmt.Printf("%-14s %-25.25s %-14s\n",
				daftar[i].Nim,
				daftar[i].Nama,
				daftar[i].Kelas,
			)
		}
	}
}

func ubahMahasiswa(daftar *[100]Mahasiswa, jumlah int) {
	nimDicari := bacaTeks("Masukkan NIM yang ingin diubah: ")
	posisi := cariMahasiswaBinary(*daftar, jumlah, nimDicari)

	if posisi == -1 {
		fmt.Println("Mahasiswa tidak ditemukan.")
	} else {
		daftar[posisi].Nama = bacaTeks("Nama baru tanpa spasi: ")
		daftar[posisi].Kelas = bacaTeks("Kelas baru: ")

		fmt.Println("Data mahasiswa berhasil diubah.")
	}
}

func hapusMahasiswa(daftar *[100]Mahasiswa, jumlah *int) {
	nimDicari := bacaTeks("Masukkan NIM yang ingin dihapus: ")
	posisi := cariMahasiswaBinary(*daftar, *jumlah, nimDicari)

	if posisi == -1 {
		fmt.Println("Mahasiswa tidak ditemukan.")
	} else {
		for i := posisi; i < *jumlah-1; i++ {
			daftar[i] = daftar[i+1]
		}

		daftar[*jumlah-1] = Mahasiswa{}
		*jumlah = *jumlah - 1

		fmt.Println("Data mahasiswa berhasil dihapus.")
	}
}

func cariMahasiswaSequential(daftar [100]Mahasiswa, jumlah int, nim string) int {
	posisi := -1
	i := 0

	for i < jumlah && posisi == -1 {
		if daftar[i].Nim == nim {
			posisi = i
		}

		i = i + 1
	}

	return posisi
}

func cariMahasiswaBinary(daftar [100]Mahasiswa, jumlah int, nim string) int {
	kiri := 0
	kanan := jumlah - 1
	posisi := -1

	for kiri <= kanan && posisi == -1 {
		tengah := (kiri + kanan) / 2

		if daftar[tengah].Nim == nim {
			posisi = tengah
		} else if daftar[tengah].Nim < nim {
			kiri = tengah + 1
		} else {
			kanan = tengah - 1
		}
	}

	return posisi
}

func urutkanNimInsertion(daftar *[100]Mahasiswa, jumlah int) {
	for i := 1; i < jumlah; i++ {
		mahasiswaDipilih := daftar[i]
		j := i - 1

		for j >= 0 && daftar[j].Nim > mahasiswaDipilih.Nim {
			daftar[j+1] = daftar[j]
			j = j - 1
		}

		daftar[j+1] = mahasiswaDipilih
	}
}

func bacaAngka(pesan string) int {
	angka := 0
	valid := false

	for !valid {
		teks := bacaTeks(pesan)
		hasil, berhasil := ubahKeAngka(teks)

		if berhasil {
			angka = hasil
			valid = true
		} else {
			fmt.Println("Input harus berupa angka.")
		}
	}

	return angka
}

func bacaTeks(pesan string) string {
	var teks string
	valid := false

	for !valid {
		fmt.Print(pesan)
		_, err := fmt.Scanln(&teks)

		if err == nil && teks != "" {
			valid = true
		} else {
			fmt.Println("Input tidak boleh kosong atau mengandung spasi.")
		}
	}

	return teks
}

func ubahKeAngka(teks string) (int, bool) {
	angka := 0
	valid := true
	i := 0

	if teks == "" {
		valid = false
	}

	for i < len(teks) && valid {
		if teks[i] < '0' || teks[i] > '9' {
			valid = false
		} else {
			angka = angka*10 + int(teks[i]-'0')
		}

		i = i + 1
	}

	return angka, valid
}
