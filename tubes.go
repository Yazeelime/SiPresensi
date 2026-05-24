package main

import "fmt"

type Mahasiswa struct {
	Nim   string
	Nama  string
	Kelas string
}

var daftarMahasiswa []Mahasiswa

func main() {
	for {
		fmt.Println()
		fmt.Println("=== SiPresensi ===")
		fmt.Println("1. Tambah Mahasiswa")
		fmt.Println("2. Tampilkan Mahasiswa")
		fmt.Println("3. Ubah Mahasiswa")
		fmt.Println("4. Hapus Mahasiswa")
		fmt.Println("0. Keluar")

		pilihan := bacaAngka("Pilih menu: ")

		switch pilihan {
		case 1:
			tambahMahasiswa()
		case 2:
			tampilkanMahasiswa()
		case 3:
			ubahMahasiswa()
		case 4:
			hapusMahasiswa()
		case 0:
			fmt.Println("Program selesai.")
			return
		default:
			fmt.Println("Menu tidak tersedia.")
		}
	}
}

func tambahMahasiswa() {
	var mahasiswaBaru Mahasiswa

	mahasiswaBaru.Nim = bacaTeks("NIM: ")

	if cariIndexMahasiswa(mahasiswaBaru.Nim) != -1 {
		fmt.Println("NIM sudah terdaftar.")
		return
	}

	mahasiswaBaru.Nama = bacaTeks("Nama tanpa spasi: ")
	mahasiswaBaru.Kelas = bacaTeks("Kelas: ")

	daftarMahasiswa = append(daftarMahasiswa, mahasiswaBaru)
	fmt.Println("Data mahasiswa berhasil ditambahkan.")
}

func tampilkanMahasiswa() {
	if len(daftarMahasiswa) == 0 {
		fmt.Println("Belum ada data mahasiswa.")
		return
	}

	fmt.Println()
	fmt.Printf("%-14s %-25s %-14s\n", "NIM", "Nama", "Kelas")

	for _, mahasiswa := range daftarMahasiswa {
		fmt.Printf("%-14s %-25.25s %-14s\n",
			mahasiswa.Nim,
			mahasiswa.Nama,
			mahasiswa.Kelas,
		)
	}
}

func ubahMahasiswa() {
	nimDicari := bacaTeks("Masukkan NIM yang ingin diubah: ")
	posisi := cariIndexMahasiswa(nimDicari)

	if posisi == -1 {
		fmt.Println("Mahasiswa tidak ditemukan.")
		return
	}

	daftarMahasiswa[posisi].Nama = bacaTeks("Nama baru tanpa spasi: ")
	daftarMahasiswa[posisi].Kelas = bacaTeks("Kelas baru: ")

	fmt.Println("Data mahasiswa berhasil diubah.")
}

func hapusMahasiswa() {
	nimDicari := bacaTeks("Masukkan NIM yang ingin dihapus: ")
	posisi := cariIndexMahasiswa(nimDicari)

	if posisi == -1 {
		fmt.Println("Mahasiswa tidak ditemukan.")
		return
	}

	daftarMahasiswa = append(daftarMahasiswa[:posisi], daftarMahasiswa[posisi+1:]...)
	fmt.Println("Data mahasiswa berhasil dihapus.")
}

func cariIndexMahasiswa(nim string) int {
	for i, mahasiswa := range daftarMahasiswa {
		if mahasiswa.Nim == nim {
			return i
		}
	}

	return -1
}

func bacaAngka(pesan string) int {
	for {
		teks := bacaTeks(pesan)
		angka, valid := ubahKeAngka(teks)

		if valid {
			return angka
		}

		fmt.Println("Input harus berupa angka.")
	}
}

func bacaTeks(pesan string) string {
	var teks string

	for {
		fmt.Print(pesan)
		_, err := fmt.Scanln(&teks)

		if err == nil && teks != "" {
			return teks
		}

		fmt.Println("Input tidak boleh kosong atau mengandung spasi.")
	}
}

func ubahKeAngka(teks string) (int, bool) {
	if teks == "" {
		return 0, false
	}

	angka := 0

	for i := 0; i < len(teks); i++ {
		if teks[i] < '0' || teks[i] > '9' {
			return 0, false
		}

		angka = angka*10 + int(teks[i]-'0')
	}

	return angka, true
}