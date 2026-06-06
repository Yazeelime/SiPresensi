package main

import "fmt"

type Mahasiswa struct {
	Nim   string
	Nama  string
	Kelas string
}

type Presensi struct {
	Nim        string
	MataKuliah string
	Status     string
}

var dataPresensi [1000]Presensi

var dataMahasiswa [100]Mahasiswa

func main() {
	jumlahMahasiswa := 0
	jumlahPresensi := 0
	selesai := false

	for !selesai {
		tampilkanMenuUtama()
		pilihan := bacaAngka("Pilih menu: ")

		switch pilihan {
		case 1:
			menuDataMahasiswa(&dataMahasiswa, &jumlahMahasiswa)
		case 2:
			menuDataPresensi(
				&dataPresensi,
				&jumlahPresensi,
				dataMahasiswa,
				jumlahMahasiswa,
			)
		case 3:
			menuCariMahasiswa(
				&dataMahasiswa,
				jumlahMahasiswa,
			)
		case 4:
			menuStatistik(
				dataMahasiswa,
				jumlahMahasiswa,
				dataPresensi,
				jumlahPresensi,
	)
		case 0:
			fmt.Println("Terima kasih telah mengakses SiPresensi.")
			selesai = true
		default:
			fmt.Println("Menu tidak tersedia.")
		}
	}
}

func tampilkanMenuUtama() {
	fmt.Println()
	fmt.Println("=== SiPresensi ===")
	fmt.Println("1. Data Mahasiswa")
	fmt.Println("2. Data Presensi")
	fmt.Println("3. Cari Mahasiswa")
	fmt.Println("4. Statistik")
	fmt.Println("0. Keluar")
}
func menuDataMahasiswa(daftar *[100]Mahasiswa, jumlah *int) {
	kembali := false

	for !kembali {
		fmt.Println()
		fmt.Println("=== Data Mahasiswa ===")
		fmt.Println("1. Tambah Mahasiswa")
		fmt.Println("2. Tampilkan Mahasiswa")
		fmt.Println("3. Ubah Mahasiswa")
		fmt.Println("4. Hapus Mahasiswa")
		fmt.Println("0. Kembali")

		pilihan := bacaAngka("Pilih menu: ")

		switch pilihan {
		case 1:
			tambahMahasiswa(daftar, jumlah)
		case 2:
			tampilkanMahasiswa(*daftar, *jumlah)
		case 3:
			ubahMahasiswa(daftar, *jumlah)
		case 4:
			hapusMahasiswa(daftar, jumlah)
		case 0:
			kembali = true
		default:
			fmt.Println("Menu tidak tersedia.")
		}
	}
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

func menuDataPresensi(
	daftarPresensi *[1000]Presensi,
	jumlahPresensi *int,
	daftarMahasiswa [100]Mahasiswa,
	jumlahMahasiswa int,
) {
	kembali := false

	for !kembali {
		fmt.Println()
		fmt.Println("=== Data Presensi ===")
		fmt.Println("1. Tambah Presensi")
		fmt.Println("2. Tampilkan Presensi")
		fmt.Println("3. Ubah Presensi")
		fmt.Println("4. Hapus Presensi")
		fmt.Println("0. Kembali")

		pilihan := bacaAngka("Pilih menu: ")

		switch pilihan {
		case 1:
			tambahPresensi(
				daftarPresensi,
				jumlahPresensi,
				daftarMahasiswa,
				jumlahMahasiswa,
			)
		case 2:
			tampilkanPresensi(*daftarPresensi, *jumlahPresensi)
		case 3:
			ubahPresensi(daftarPresensi, *jumlahPresensi)
		case 4:
			hapusPresensi(daftarPresensi, jumlahPresensi)
		case 0:
			kembali = true
		default:
			fmt.Println("Menu tidak tersedia.")
		}
	}
}

func tambahPresensi(
	daftarPresensi *[1000]Presensi,
	jumlahPresensi *int,
	daftarMahasiswa [100]Mahasiswa,
	jumlahMahasiswa int,
) {
	var presensiBaru Presensi

	if *jumlahPresensi >= 1000 {
		fmt.Println("Data presensi sudah penuh.")
		return
	}

	presensiBaru.Nim = bacaTeks("NIM: ")

	if cariMahasiswaSequential(
		daftarMahasiswa,
		jumlahMahasiswa,
		presensiBaru.Nim,
	) == -1 {

		fmt.Println("Mahasiswa tidak ditemukan.")
		return
	}

	presensiBaru.MataKuliah = bacaTeks("Mata Kuliah: ")
	presensiBaru.Status = bacaStatus()

	daftarPresensi[*jumlahPresensi] = presensiBaru
	*jumlahPresensi++

	fmt.Println("Presensi berhasil ditambahkan.")
}

func tampilkanPresensi(
	daftar [1000]Presensi,
	jumlah int,
) {
	if jumlah == 0 {
		fmt.Println("Belum ada data presensi.")
		return
	}

	fmt.Println()
	fmt.Printf("%-15s %-25s %-10s\n",
		"NIM",
		"Mata Kuliah",
		"Status")

	for i := 0; i < jumlah; i++ {
		fmt.Printf("%-15s %-25s %-10s\n",
			daftar[i].Nim,
			daftar[i].MataKuliah,
			daftar[i].Status)
	}
}

func ubahPresensi(
	daftar *[1000]Presensi,
	jumlah int,
) {
	nim := bacaTeks("NIM: ")
	mataKuliah := bacaTeks("Mata Kuliah: ")

	posisi := cariPresensi(
		*daftar,
		jumlah,
		nim,
		mataKuliah,
	)

	if posisi == -1 {
		fmt.Println("Data presensi tidak ditemukan.")
		return
	}

	daftar[posisi].Status = bacaStatus()

	fmt.Println("Data presensi berhasil diubah.")
}

func hapusPresensi(
	daftar *[1000]Presensi,
	jumlah *int,
) {
	nim := bacaTeks("NIM: ")
	mataKuliah := bacaTeks("Mata Kuliah: ")

	posisi := cariPresensi(
		*daftar,
		*jumlah,
		nim,
		mataKuliah,
	)

	if posisi == -1 {
		fmt.Println("Data presensi tidak ditemukan.")
		return
	}

	for i := posisi; i < *jumlah-1; i++ {
		daftar[i] = daftar[i+1]
	}

	daftar[*jumlah-1] = Presensi{}
	*jumlah--

	fmt.Println("Data presensi berhasil dihapus.")
}

func cariPresensi(
	daftar [1000]Presensi,
	jumlah int,
	nim string,
	mataKuliah string,
) int {

	for i := 0; i < jumlah; i++ {
		if daftar[i].Nim == nim &&
			daftar[i].MataKuliah == mataKuliah {

			return i
		}
	}

	return -1
}

func bacaStatus() string {
	for {
		status := bacaTeks(
			"Status (hadir/izin/sakit/alpa): ",
		)

		if status == "hadir" ||
			status == "izin" ||
			status == "sakit" ||
			status == "alpa" {

			return status
		}

		fmt.Println("Status tidak valid.")
	}
}

func menuCariMahasiswa(
	daftar *[100]Mahasiswa,
	jumlah int,
) {
	kembali := false

	for !kembali {
		fmt.Println()
		fmt.Println("=== Cari Mahasiswa ===")
		fmt.Println("1. Cari NIM")
		fmt.Println("2. Urutkan Berdasarkan Nama")
		fmt.Println("0. Kembali")

		pilihan := bacaAngka("Pilih menu: ")

		switch pilihan {
		case 1:
			cariMahasiswaMenu(*daftar, jumlah)

		case 2:
			menuUrutNama(daftar, jumlah)

		case 0:
			kembali = true

		default:
			fmt.Println("Menu tidak tersedia.")
		}
	}
}

func cariMahasiswaMenu(
	daftar [100]Mahasiswa,
	jumlah int,
) {
	nim := bacaTeks("Masukkan NIM: ")

	posisi := cariMahasiswaBinary(
		daftar,
		jumlah,
		nim,
	)

	if posisi == -1 {
		fmt.Println("Mahasiswa tidak ditemukan.")
	} else {
		fmt.Println()
		fmt.Printf("%-14s %-25s %-14s\n",
			"NIM",
			"Nama",
			"Kelas")

		fmt.Printf("%-14s %-25s %-14s\n",
			daftar[posisi].Nim,
			daftar[posisi].Nama,
			daftar[posisi].Kelas)
	}
}

func menuUrutNama(
	daftar *[100]Mahasiswa,
	jumlah int,
) {
	fmt.Println()
	fmt.Println("1. Ascending")
	fmt.Println("2. Descending")

	pilihan := bacaAngka("Pilih urutan: ")

	switch pilihan {
	case 1:
		urutNamaAscending(daftar, jumlah)
		tampilkanMahasiswa(*daftar, jumlah)

	case 2:
		urutNamaDescending(daftar, jumlah)
		tampilkanMahasiswa(*daftar, jumlah)

	default:
		fmt.Println("Pilihan tidak tersedia.")
	}
}

func urutNamaAscending(
	daftar *[100]Mahasiswa,
	jumlah int,
) {
	for i := 1; i < jumlah; i++ {
		dataDipilih := daftar[i]
		j := i - 1

		for j >= 0 &&
			daftar[j].Nama > dataDipilih.Nama {

			daftar[j+1] = daftar[j]
			j--
		}

		daftar[j+1] = dataDipilih
	}
}

func urutNamaDescending(
	daftar *[100]Mahasiswa,
	jumlah int,
) {
	for i := 1; i < jumlah; i++ {
		dataDipilih := daftar[i]
		j := i - 1

		for j >= 0 &&
			daftar[j].Nama < dataDipilih.Nama {

			daftar[j+1] = daftar[j]
			j--
		}

		daftar[j+1] = dataDipilih
	}
}

func menuStatistik(
	daftarMahasiswa [100]Mahasiswa,
	jumlahMahasiswa int,
	daftarPresensi [1000]Presensi,
	jumlahPresensi int,
) {
	kembali := false

	for !kembali {
		fmt.Println()
		fmt.Println("=== Statistik ===")
		fmt.Println("1. Persentase Kehadiran Per Kelas")
		fmt.Println("2. Mahasiswa Dengan Alpa Terbanyak")
		fmt.Println("0. Kembali")

		pilihan := bacaAngka("Pilih menu: ")

		switch pilihan {
		case 1:
			statistikKehadiranKelas(
				daftarMahasiswa,
				jumlahMahasiswa,
				daftarPresensi,
				jumlahPresensi,
			)

		case 2:
			alpaTerbanyak(
				daftarMahasiswa,
				jumlahMahasiswa,
				daftarPresensi,
				jumlahPresensi,
			)

		case 0:
			kembali = true
		}
	}
}

func statistikKehadiranKelas(
	daftarMahasiswa [100]Mahasiswa,
	jumlahMahasiswa int,
	daftarPresensi [1000]Presensi,
	jumlahPresensi int,
) {
	kelas := bacaTeks("Masukkan kelas: ")

	totalPresensi := 0
	totalHadir := 0

	for i := 0; i < jumlahMahasiswa; i++ {

		if daftarMahasiswa[i].Kelas == kelas {

			for j := 0; j < jumlahPresensi; j++ {

				if daftarPresensi[j].Nim ==
					daftarMahasiswa[i].Nim {

					totalPresensi++

					if daftarPresensi[j].Status == "hadir" {
						totalHadir++
					}
				}
			}
		}
	}

	if totalPresensi == 0 {
		fmt.Println("Belum ada data presensi.")
		return
	}

	persentase :=
		float64(totalHadir) /
			float64(totalPresensi) * 100

	fmt.Printf(
		"Persentase Kehadiran %.2f%%\n",
		persentase,
	)
}

func alpaTerbanyak(
	daftarMahasiswa [100]Mahasiswa,
	jumlahMahasiswa int,
	daftarPresensi [1000]Presensi,
	jumlahPresensi int,
) {
	fmt.Println()
	fmt.Println("Mahasiswa Dengan Alpa Terbanyak")
	fmt.Println()

	for i := 0; i < jumlahMahasiswa; i++ {

		jumlahAlpa := 0

		for j := 0; j < jumlahPresensi; j++ {

			if daftarPresensi[j].Nim ==
				daftarMahasiswa[i].Nim &&
				daftarPresensi[j].Status == "alpa" {

				jumlahAlpa++
			}
		}

		fmt.Printf(
			"%-14s %-20s Alpa: %d\n",
			daftarMahasiswa[i].Nim,
			daftarMahasiswa[i].Nama,
			jumlahAlpa,
		)
	}
}