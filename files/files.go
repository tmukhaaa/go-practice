package files

import (
	"fmt"
	"os"
)

type JsonDb struct {
	filename string
}

func NewJsonDb(name string) *JsonDb {
	return &JsonDb{
		filename : name,
	}
}

func (db *JsonDb) Read() ([]byte, error) {
	file, err := os.ReadFile(db.filename)
	if err != nil {
		fmt.Println("Ошибка в чтении файла!")
		return nil, err
	}
	return file, nil

}

func (db *JsonDb) Write(content []byte){
	file, err := os.Create(db.filename)
	if err != nil {
		fmt.Println("Ошибка в создании файла!", err)
		return
	}
	defer file.Close()
	_, err = file.Write(content)
	if err != nil {
		file.Close()
		fmt.Println("Не удалось записать данные!")
		return 
	}
	fmt.Println("Запись прошла успешно!")
	
}