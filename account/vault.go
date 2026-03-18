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

type VaultWithDb struct {
	Vault
	db files.JsonDb
}

func NewVault(db files.JsonDb) *VaultWithDb {
	file, err := db.Read()
	if err != nil {
		return &VaultWithDb{
			Vault: Vault{
				Accounts:  []*Account{},
				UpdatedAt: time.Now(),
			},
			db : db,
			
		}
	}
	var vault VaultWithDb
	err = json.Unmarshal(file, &vault)
	vault.db = db
	if err != nil {
		fmt.Println("не удалось преобразовать файл")
		return &VaultWithDb{
			Vault: Vault{
				Accounts:  []*Account{},
				UpdatedAt: time.Now(),
			},
			db : db,
			
		}
	}
	return &vault
}

func (vault *VaultWithDb) CreateVault(acc *Account) {
	vault.Vault.Accounts = append(vault.Vault.Accounts, acc)
	vault.Vault.UpdatedAt = time.Now()
	file, err := vault.ToVault()
	if err != nil {
		fmt.Println("не удалость преобразовать в байты")
		return
	}
	vault.db.Write(file)
}

func (vault *VaultWithDb) ToVault() ([]byte, error) {
	file, err := json.Marshal(*vault)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (vault *VaultWithDb) FindVault(url string) []Account {
	var arr []Account
	for _, value := range vault.Accounts {
		if strings.Contains(value.Url, url) {
			arr = append(arr, *value)
		}
	}
	return arr
}

func (vault *VaultWithDb) DeleteVault(url string) {
	NewVault := make([]*Account, 0, len(vault.Accounts))
	for _, value := range vault.Accounts {
		if value.Url != url {
			NewVault = append(NewVault, value)

		}
	}
	vault.Accounts = NewVault
	vault.UpdatedAt = time.Now()
	file, _ := vault.ToVault()
	vault.db.Write(file)
}
