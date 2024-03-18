package main

import (
	"fmt"
)

var menu, ID, name, major string

func main() {
	fmt.Println("---- Menu ----")
	fmt.Println("1. Login")
	fmt.Println("2. Register")
	fmt.Println("3. Get Study Program")
	fmt.Println("Pilih menu: ")
	fmt.Scan(&menu)

	switch menu {
	case "1":
		Login()
	case "2":
		Register()

	case "3":
		GetStudyProgram()
	default:
		fmt.Println("Pilih menu yang benar!")
	}

}

func Register() {
	// fmt.Print("Enter your ID name major:")
	// fmt.Scanln()
	// fmt.Scanf("%s %s %s", &ID, &name, &major)

	// if len(ID) < 5 {
	// 	fmt.Println("ID must be 5 characters long!")
	// 	if ID == "" || name == "" {
	// 		fmt.Println("ID or Name is undefined!")
	// 		return
	// 	}
	// 	return
	// }

	// account := []string{ID, name, major}
	// //fmt.Println(account)

	// newAccount := append(account, account...)
	// fmt.Println(newAccount)
	// fmt.Println(account)

	// fmt.Println("Register success: ", name, "(", major, ")")

}

func Login()           {}
func GetStudyProgram() {}
