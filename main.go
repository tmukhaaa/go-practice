package main

import (
	"demo/password/account"
	"demo/password/files"
	"fmt"
)



func main(){
	menu()
	
}

func menu() {
	var input int
itteration:
	for {
		fmt.Println(`1. Создать аккаунт
2. Найти аккаунт
3. Удалить аккаунт
4. Выход`)
		fmt.Print("Введите пункт меню(1-4): ")

		fmt.Scan(&input)
		switch input {
		case 1:
			createAccount()
		case 2:
			account.FindAccount()
		case 3:
			account.DeleteAccount()
		case 4:
			break itteration
		default:
			fmt.Println("Неверный ввод! Введите пункт меню(1-4):")
	}
	}
}

func createAccount(){
	login := account.GetInfo("write your login: ")
	password:= account.GetInfo("write your password: ")
	url := account.GetInfo("write your url: ")
	account1, err := account.NewAccount(login,password,url)
	if err != nil {
		fmt.Println(err)
		return
	}
	vault := account.NewVault(*files.NewJsonDb("pass.json"))
	vault.CreateVault(account1)
}



