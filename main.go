package main

import "fmt"

type account struct {
	login    string
	password string
	url      string
}

func main() {
	login := promptData("Введите логин")
	password := promptData("Введите пароль")
	url := promptData("Введите URL")

	// account1 := account{
	// 	password,
	// 	"",
	// 	url,
	// }
	account1 := account{
		password: password,
		url:      url,
		login:    login,
	}

	outputPassword(login, password, url)
}

func promptData(prompt string) string {
	fmt.Print(prompt)
	var res string
	fmt.Scan(&res)
	return res
}

func outputPassword(login, password, url string) {
	fmt.Println(login, password, url)
}
