package main

import (
	"fmt"
	"os"
	"strconv"
)

// Struct untuk menyimpan data teman
type Friend struct {
	ID        int
	Nama      string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

// Slice untuk menyimpan data teman-teman
var friends = []Friend{
	{1, "Rifqy", "Alamat1", "Pekerjaan1", "Alasan1"},
	{2, "Andi", "Alamat2", "Pekerjaan2", "Alasan2"},
	{3, "Budi", "Alamat3", "Pekerjaan3", "Alasan3"},
	{4, "Fitri", "Alamat4", "Pekerjaan4", "Alasan4"},
	{5, "Ryan", "Alamat5", "Pekerjaan5", "Alasan5"},
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Tolong masukan nama atau nomor absen")
		fmt.Println("Contoh: 'go run main.go Rifqy' atau 'go run main.go'")
		return
	}

	searchKey := os.Args[1]

	foundFriends := findFriends(searchKey)

	if len(foundFriends) == 0 {
		fmt.Println("Tidak ada teman yang ditemukan.")
	} else {
		fmt.Println("Data teman yang ditemukan:")
		for _, friend := range foundFriends {
			printFriendDetails(friend)
		}
	}
}

// Fungsi untuk mencari teman berdasarkan nama atau ID
func findFriends(searchKey string) []Friend {
	var foundFriends []Friend
	for _, friend := range friends {
		if strconv.Itoa(friend.ID) == searchKey || friend.Nama == searchKey {
			foundFriends = append(foundFriends, friend)
		}
	}
	return foundFriends
}

// Fungsi untuk mencetak detail teman
func printFriendDetails(friend Friend) {
	fmt.Printf("ID: %d\n", friend.ID)
	fmt.Printf("Nama: %s\n", friend.Nama)
	fmt.Printf("Alamat: %s\n", friend.Alamat)
	fmt.Printf("Pekerjaan: %s\n", friend.Pekerjaan)
	fmt.Printf("Alasan memilih kelas Golang: %s\n", friend.Alasan)
	fmt.Println("----------")
}
