package main

import (
	"fmt"
	"strings"
)

var id, nama, jurusan string

var Students string = "A1234_Aditira_TI, B2131_Dito_TK, A3455_Afis_MI, WE123_Yiyi_JK"
var StudentStudyPrograms string = "TI_Teknik Informatika, TK_Teknik Komputer, SI_Sistem Informasi, MI_Manajemen Informasi"

func main() {
	fmt.Print("Enter your ID name major:")
	fmt.Scanf("%s %s %s", &id, &nama, &jurusan)

	if len(id) < 5 {
		fmt.Println("ID must be 5 characters long!")
		if id == "" || nama == "" {
			fmt.Println("ID or Name is undefined!")
			return
		}
		return
	}

	// Memisahkan elemen menjadi ID, nama, dan jurusan
	parts := strings.Split(Students, ",")
	for _, mhs := range parts {
		arr := strings.Fields(mhs)          // Memisahkan elemen menjadi ID, nama, dan jurusan
		idMhs := strings.Split(arr[0], "_") // Memisahkan ID dari nama dan jurusan
		if id == idMhs[0] {
			fmt.Println("ID is already taken!")
			return
		}
	}

	// Menambahkan data baru ke dalam string Students
	Students += fmt.Sprintf(", %s_%s_%s", id, nama, jurusan)

	fmt.Println("Register success")
	fmt.Println(Students)
}
