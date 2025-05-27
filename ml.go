package main

import (
	"fmt"
	"strings"
)

const maxHero = 100

type Role string

const (
	Tank     Role = "Tank"
	Fighter  Role = "Fighter"
	Assassin Role = "Assassin"
	Mage     Role = "Mage"
	Marksman Role = "Marksman"
	Support  Role = "Support"
)

type Hero struct {
	Nama             string
	Role             Role
	TingkatKesulitan int
	HP               int
	Damage           int
	WinRate          float64
}

var dataHero [maxHero]Hero
var jumlahHero int = 0

func validasiRole(input string) bool {
	switch Role(strings.Title(input)) {
	case Tank, Fighter, Assassin, Mage, Marksman, Support:
		return true
	}
	return false
}

func isiDataAwal() {
	dataHero[0] = Hero{"Layla", Marksman, 2, 2500, 150, 51.3}
	dataHero[1] = Hero{"Tigreal", Tank, 4, 4000, 100, 48.7}
	dataHero[2] = Hero{"Fanny", Assassin, 9, 2200, 210, 45.1}
	dataHero[3] = Hero{"Nana", Support, 3, 2700, 120, 55.4}
	dataHero[4] = Hero{"Gusion", Assassin, 8, 2300, 200, 49.9}
	dataHero[5] = Hero{"Lancelot", Assassin, 7, 2400, 195, 50.2}
	dataHero[6] = Hero{"Alice", Mage, 6, 2600, 180, 53.0}
	dataHero[7] = Hero{"Miya", Marksman, 2, 2450, 140, 52.5}
	dataHero[8] = Hero{"Franco", Tank, 5, 4100, 110, 46.7}
	dataHero[9] = Hero{"Eudora", Mage, 3, 2300, 175, 54.8}
	dataHero[10] = Hero{"Aldous", Fighter, 6, 3000, 190, 49.2}
	dataHero[11] = Hero{"Khufra", Tank, 7, 4200, 105, 47.5}
	dataHero[12] = Hero{"Beatrix", Marksman, 8, 2500, 160, 50.9}
	dataHero[13] = Hero{"Angela", Support, 4, 2600, 110, 56.2}
	dataHero[14] = Hero{"Lylia", Mage, 7, 2200, 200, 53.9}
	dataHero[15] = Hero{"X.Borg", Fighter, 5, 3100, 185, 48.4}
	dataHero[16] = Hero{"Hayabusa", Assassin, 7, 2350, 205, 50.0}
	dataHero[17] = Hero{"Cyclops", Mage, 3, 2400, 170, 54.1}
	dataHero[18] = Hero{"Granger", Marksman, 6, 2600, 165, 51.7}
	dataHero[19] = Hero{"Akai", Tank, 4, 4050, 95, 49.3}
	jumlahHero = 20
}

func tambahHero() {
	if jumlahHero >= maxHero {
		fmt.Println("Data hero penuh.")
		return
	}
	var nama, roleInput string
	var role Role
	var kesulitan, hp, dmg int
	var winRate float64

	fmt.Print("Nama Hero: ")
	fmt.Scan(&nama)

	for {
		fmt.Print("Role (Tank/Fighter/Assassin/Mage/Marksman/Support): ")
		fmt.Scan(&roleInput)
		if validasiRole(roleInput) {
			role = Role(strings.Title(roleInput))
			break
		}
		fmt.Println("Role tidak valid.")
	}

	fmt.Print("Tingkat Kesulitan (1-10): ")
	fmt.Scan(&kesulitan)
	fmt.Print("HP: ")
	fmt.Scan(&hp)
	fmt.Print("Damage: ")
	fmt.Scan(&dmg)
	fmt.Print("Win Rate (contoh: 59.5): ")
	fmt.Scan(&winRate)

	dataHero[jumlahHero] = Hero{nama, role, kesulitan, hp, dmg, winRate}
	jumlahHero++
	fmt.Println("Hero berhasil ditambahkan.")
}

func tampilkanData() {
	if jumlahHero == 0 {
		fmt.Println("Belum ada data.")
		return
	}
	fmt.Println("=== Daftar Hero ===")
	for i := 0; i < jumlahHero; i++ {
		h := dataHero[i]
		fmt.Printf("%d. %s | Role: %s | Kesulitan: %d | HP: %d | DMG: %d | Win Rate: %.1f%%\n",
			i+1, h.Nama, h.Role, h.TingkatKesulitan, h.HP, h.Damage, h.WinRate)
	}
}

func sequentialSearch(nama string) int {
	for i := 0; i < jumlahHero; i++ {
		if strings.EqualFold(dataHero[i].Nama, nama) {
			return i
		}
	}
	return -1
}

func binarySearch(nama string) int {
	left, right := 0, jumlahHero-1
	for left <= right {
		mid := (left + right) / 2
		if strings.EqualFold(dataHero[mid].Nama, nama) {
			return mid
		} else if strings.ToLower(nama) < strings.ToLower(dataHero[mid].Nama) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func selectionSortByNama(asc bool) {
	for i := 0; i < jumlahHero-1; i++ {
		selected := i
		for j := i + 1; j < jumlahHero; j++ {
			if (asc && strings.ToLower(dataHero[j].Nama) < strings.ToLower(dataHero[selected].Nama)) ||
				(!asc && strings.ToLower(dataHero[j].Nama) > strings.ToLower(dataHero[selected].Nama)) {
				selected = j
			}
		}
		dataHero[i], dataHero[selected] = dataHero[selected], dataHero[i]
	}
}

func insertionSortByWinRate(asc bool) {
	for i := 1; i < jumlahHero; i++ {
		key := dataHero[i]
		j := i - 1
		for j >= 0 && ((asc && dataHero[j].WinRate > key.WinRate) || (!asc && dataHero[j].WinRate < key.WinRate)) {
			dataHero[j+1] = dataHero[j]
			j--
		}
		dataHero[j+1] = key
	}
}

func hapusHero() {
	var nama string
	fmt.Print("Masukkan nama hero yang ingin dihapus: ")
	fmt.Scan(&nama)
	idx := sequentialSearch(nama)
	if idx == -1 {
		fmt.Println("Hero tidak ditemukan.")
		return
	}
	for i := idx; i < jumlahHero-1; i++ {
		dataHero[i] = dataHero[i+1]
	}
	jumlahHero--
	fmt.Println("Hero berhasil dihapus.")
}

func editHero() {
	var nama string
	fmt.Print("Masukkan nama hero yang ingin diedit: ")
	fmt.Scan(&nama)
	idx := sequentialSearch(nama)
	if idx == -1 {
		fmt.Println("Hero tidak ditemukan.")
		return
	}
	var roleInput string
	var kesulitan, hp, dmg int
	var winRate float64

	fmt.Printf("Data lama: %+v\n", dataHero[idx])
	fmt.Print("Role (Tank/Fighter/Assassin/Mage/Marksman/Support): ")
	fmt.Scan(&roleInput)
	if validasiRole(roleInput) {
		dataHero[idx].Role = Role(strings.Title(roleInput))
	}
	fmt.Print("Tingkat Kesulitan: ")
	fmt.Scan(&kesulitan)
	fmt.Print("HP: ")
	fmt.Scan(&hp)
	fmt.Print("Damage: ")
	fmt.Scan(&dmg)
	fmt.Print("Win Rate: ")
	fmt.Scan(&winRate)

	dataHero[idx].TingkatKesulitan = kesulitan
	dataHero[idx].HP = hp
	dataHero[idx].Damage = dmg
	dataHero[idx].WinRate = winRate

	fmt.Println("Data hero berhasil diperbarui.")
}

func menuUrut() {
	var opsi int
	fmt.Println("1. Urutkan berdasarkan Nama (Ascending)")
	fmt.Println("2. Urutkan berdasarkan Nama (Descending)")
	fmt.Println("3. Urutkan berdasarkan Win Rate (Ascending)")
	fmt.Println("4. Urutkan berdasarkan Win Rate (Descending)")
	fmt.Print("Pilih: ")
	fmt.Scan(&opsi)
	if opsi == 1 {
		selectionSortByNama(true)
	} else if opsi == 2 {
		selectionSortByNama(false)
	} else if opsi == 3 {
		insertionSortByWinRate(true)
	} else if opsi == 4 {
		insertionSortByWinRate(false)
	} else {
		fmt.Println("Pilihan tidak valid.")
	}
	fmt.Println("Data telah diurutkan.")
}

func menu() {
	var pilihan int
	for pilihan != 8 {
		fmt.Println("\n=== MENU MOBILE LEGENDS HERO MANAGER ===")
		fmt.Println("1. Tambah Hero")
		fmt.Println("2. Tampilkan Semua Hero")
		fmt.Println("3. Cari Hero (Sequential Search)")
		fmt.Println("4. Cari Hero (Binary Search)")
		fmt.Println("5. Edit Data Hero")
		fmt.Println("6. Hapus Hero")
		fmt.Println("7. Urutkan Hero")
		fmt.Println("8. Keluar")
		fmt.Print("Pilihan: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tambahHero()
		case 2:
			tampilkanData()
		case 3:
			var nama string
			fmt.Print("Masukkan nama hero: ")
			fmt.Scan(&nama)
			idx := sequentialSearch(nama)
			if idx != -1 {
				h := dataHero[idx]
				fmt.Printf("Ditemukan: %s - %s - Kesulitan: %d - HP: %d - DMG: %d - Win Rate: %.1f%%\n", h.Nama, h.Role, h.TingkatKesulitan, h.HP, h.Damage, h.WinRate)
			} else {
				fmt.Println("Hero tidak ditemukan.")
			}
		case 4:
			selectionSortByNama(true)
			var nama string
			fmt.Print("Masukkan nama hero: ")
			fmt.Scan(&nama)
			idx := binarySearch(nama)
			if idx != -1 {
				h := dataHero[idx]
				fmt.Printf("Ditemukan (Binary): %s - %s - Kesulitan: %d - HP: %d - DMG: %d - Win Rate: %.1f%%\n", h.Nama, h.Role, h.TingkatKesulitan, h.HP, h.Damage, h.WinRate)
			} else {
				fmt.Println("Hero tidak ditemukan.")
			}
		case 5:
			editHero()
		case 6:
			hapusHero()
		case 7:
			menuUrut()
		case 8:
			fmt.Println("Terima kasih telah menggunakan program.")
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}

func main() {
	isiDataAwal()
	menu()
}
