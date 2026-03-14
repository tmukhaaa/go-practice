package account

import (
	"demo/password/files"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Vault struct {
	Accounts  []*Account `json:"Accounts"`
	UpdatedAt time.Time  `json:"UpdatedAt"`
}

func NewVault() *Vault {
	db := files.NewJsonDb("pass.json")
	file, err := db.Read()
	if err != nil {
		return &Vault{
			Accounts:  []*Account{},
			UpdatedAt: time.Now(),
		}
	}
	var vault Vault
	err = json.Unmarshal(file, &vault)
	if err != nil {
		fmt.Println("не удалось преобразовать файл")
		return nil
	}
	return &vault
}

func (vault *Vault) CreateVault(acc *Account) {
	vault.Accounts = append(vault.Accounts, acc)
	vault.UpdatedAt = time.Now()
	file, err := vault.ToVault()
	if err != nil {
		fmt.Println("не удалость преобразовать в байты")
		return
	}
	db := files.NewJsonDb("pass.json")
	db.Write(file)
}

func (vault *Vault) ToVault() ([]byte, error) {
	file, err := json.Marshal(vault)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (vault *Vault) FindVault(url string) []Account {
	var arr []Account
	for _, value := range vault.Accounts {
		if strings.Contains(value.Url, url) {
			arr = append(arr, *value)
		}
	}
	return arr
}

func (vault *Vault) DeleteVault(url string) {
	NewVault := make([]*Account, 0, len(vault.Accounts))
	for _, value := range vault.Accounts {
		if value.Url != url {
			NewVault = append(NewVault, value)

		}
	}
	vault.Accounts = NewVault
	vault.UpdatedAt = time.Now()
	file, _ := vault.ToVault()
	db := files.NewJsonDb("pass.json")
	db.Write(file)
}
