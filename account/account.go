package account

import (

	

	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"

	"github.com/fatih/color"
)

type Account struct {
	Login       string    `json:"login"`
	Password    string    `json:"password"`
	Url         string    `json:"url"`
	CreatedTime time.Time `json:"createdTime"`
	UpdatedTime time.Time `json:"updatedTime"`
}


func (acc Account) PrintStruct() {
	color.Red(`--------------------
your login is: %s
your password is: %s
your url is: %s
created time: %v
updated time: %v
--------------------`,
		acc.Login, acc.Password, acc.Url,acc.CreatedTime, acc.UpdatedTime,)
}

func NewAccount(login, password, urlString string) (*Account, error) {
	_, err := url.ParseRequestURI(urlString)
	if login == "" {
		return nil, errors.New("empty login")
	}

	if err != nil {
		return nil, errors.New("VALIDATION_ERROR")
	}
	newAcc := &Account{
		CreatedTime: time.Now(),
		UpdatedTime: time.Now(),
		Password:    password,
		Login:       login,
		Url:         urlString,
	}
	if password == "" {
		newAcc.getPassword(10)
	}
	return newAcc, nil
}

func GetInfo (msg string) string {
	var info string
	fmt.Print(msg)
	fmt.Scanln(&info)
	return info
}

func (acc *Account) getPassword(n int) {
	symbArray := make([]rune, 0, n)
	for range n {
		symbArray = append(symbArray, rune(rand.IntN(57)+65))
	}
	acc.Password = string(symbArray)
}

func FindAccount(vault VaultWithDb) {
	var url string
	fmt.Print("Введите URL-адрес по которому хотите совершить поиск: ")
	fmt.Scan(&url)
	acc := vault.FindVault(url)
	for _, acc := range acc {
		acc.PrintStruct()
	}
}

func DeleteAccount(vault VaultWithDb) {
	url := GetInfo("Введите URL адрес:")
	vault.DeleteVault(url)
}
