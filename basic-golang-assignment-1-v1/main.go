package main

import (
	"fmt"
	"strings"

	"a21hc3NpZ25tZW50/helper"
)

var Students string = "A1234_Aditira_TI, B2131_Dito_TK, A3455_Afis_MI"
var StudentStudyPrograms string = "TI_Teknik Informatika, TK_Teknik Komputer, SI_Sistem Informasi, MI_Manajemen Informasi"

func Login(id string, name string) string {
	var major string
	if id == "" || name == "" {
		return "ID or Name is undefined!"
	}

	if len(id) != 5 {
		return "ID must be 5 characters long!"
	}

	parts := strings.Split(Students, ",")
	for _, mhs := range parts {
		mhs = strings.TrimSpace(mhs)
		arr := strings.Split(mhs, "_")
		idMhs := arr[0]
		nameMhs := arr[1]
		majorMhs := arr[2]

		if id == idMhs && name == nameMhs {
			major = majorMhs
			return fmt.Sprintf("Login berhasil: %s (%s)", name, major)
		}
	}
	return "Login gagal: data mahasiswa tidak ditemukan"
}

func Register(id string, name string, major string) string {
	if id == "" || name == "" || major == "" {
		return "ID, Name or Major is undefined!"
	}

	if len(id) != 5 {
		return "ID must be 5 characters long!"
	}

	parts := strings.Split(Students, ",")
	for _, mhs := range parts {
		arr := strings.Split(mhs, "_")
		idMhs := arr[0]
		if id == idMhs {
			return "Registrasi gagal: id sudah digunakan"
		}
	}

	Students += fmt.Sprintf(",%s_%s_%s", id, name, major)
	return fmt.Sprintf("Registrasi berhasil: %s (%s)", name, major)
}

func GetStudyProgram(code string) string {
	if code == "" {

		return "Code is undefined!"
	} else {
		parts := strings.Split(StudentStudyPrograms, ",")

		for _, prody := range parts {
			arr := strings.Fields(prody)
			code_prody := strings.Split(arr[0], "_")
			if code == code_prody[0] {
				switch code {
				case "TI":
					return "Teknik Informatika"
				case "TK":
					return "Teknik Komputer"
				case "SI":
					return "Sistem Informasi"
				case "MI":
					return "Manajemen Informasi"
				}

			}
		}
	}

	return "Code is undefined!"
}

func main() {
	fmt.Println("Selamat datang di Student Portal!")

	for {
		helper.ClearScreen()
		fmt.Println("Students: ", Students)
		fmt.Println("Student Study Programs: ", StudentStudyPrograms)

		fmt.Println("\nPilih menu:")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("3. Get Program Study")
		fmt.Println("4. Keluar")

		var pilihan string
		fmt.Print("Masukkan pilihan Anda: ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case "1":
			helper.ClearScreen()
			var id, name string
			fmt.Print("Masukkan id: ")
			fmt.Scanln(&id)
			fmt.Print("Masukkan name: ")
			fmt.Scanln(&name)

			fmt.Println(Login(id, name))

			helper.Delay(5)
		case "2":
			helper.ClearScreen()
			var id, name, jurusan string
			fmt.Print("Masukkan id: ")
			fmt.Scanln(&id)
			fmt.Print("Masukkan name: ")
			fmt.Scanln(&name)
			fmt.Print("Masukkan jurusan: ")
			fmt.Scanln(&jurusan)
			fmt.Println(Register(id, name, jurusan))

			helper.Delay(5)
		case "3":
			helper.ClearScreen()
			var kode string
			fmt.Print("Masukkan kode: ")
			fmt.Scanln(&kode)
			fmt.Println(GetStudyProgram(kode))
			helper.Delay(5)
		case "4":
			fmt.Println("Terima kasih telah menggunakan Student Portal.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
