package main

import (
	"fmt"
	"strings"

	"a21hc3NpZ25tZW50/helper"
)

var Students = []string{
	"A1234_Aditira_TI",
	"B2131_Dito_TK",
	"A3455_Afis_MI",
}

var StudentStudyPrograms = map[string]string{
	"TI": "Teknik Informatika",
	"TK": "Teknik Komputer",
	"SI": "Sistem Informasi",
	"MI": "Manajemen Informasi",
}

type studentModifier func(string, *string)

func Login(id string, name string) string {
	if id == "" || name == "" {
		return "ID or Name is undefined!"
	}

	if len(id) != 5 {
		return "ID must be 5 characters long!"
	}

	for _, mhs := range Students[:] {
		mhs = strings.TrimSpace(mhs)
		arr := strings.Split(mhs, "_")
		idMhs := arr[0]
		nameMhs := arr[1]

		if id == idMhs && name == nameMhs {
			return fmt.Sprintf("Login berhasil: %s", name)
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

	for _, mhs := range Students[:] {
		arr := strings.Split(mhs, "_")
		idMhs := arr[0]
		if id == idMhs {
			return "Registrasi gagal: id sudah digunakan"
		}

	}

	Students = append(Students, fmt.Sprintf("%s_%s_%s", id, name, major))
	return fmt.Sprintf("Registrasi berhasil: %s (%s)", name, major)
}

func GetStudyProgram(code string) string {
	if code == "" {

		return "Kode program studi tidak ditemukan"
	} else {

		for key, value := range StudentStudyPrograms {
			if code == key {
				return value
			}
		}

	}

	return "Kode program studi tidak ditemukan"
}

func ModifyStudent(programStudi, nama string, fn studentModifier) string {
	for i, mhs := range Students[:] {
		arr := strings.Split(mhs, "_")
		nameMhs := arr[1]

		if nameMhs == nama {

			fn(programStudi, &Students[i]) //memanggil studentModifier yang menerima string dan *string
			Students[i] = fmt.Sprintf("%s_%s_%s", arr[0], arr[1], programStudi)
			return "Program studi mahasiswa berhasil diubah."
		}
	}
	return "Mahasiswa tidak ditemukan."
}

func UpdateStudyProgram(programStudi string, students *string) {
	*students = programStudi
}

func main() {
	fmt.Println("Selamat datang di Student Portal!")

	for {
		helper.ClearScreen()
		for i, student := range Students {
			fmt.Println(i+1, student)
		}

		fmt.Println("\nPilih menu:")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("3. Get Program Study")
		fmt.Println("4. Change student study program")
		fmt.Println("5. Keluar")

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
			helper.ClearScreen()
			var nama, programStudi string
			fmt.Print("Masukkan nama mahasiswa: ")
			fmt.Scanln(&nama)
			fmt.Print("Masukkan program studi baru: ")
			fmt.Scanln(&programStudi)

			fmt.Println(ModifyStudent(programStudi, nama, UpdateStudyProgram))
			helper.Delay(5)
		case "5":
			fmt.Println("Terima kasih telah menggunakan Student Portal.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
