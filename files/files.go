package files

import (
	"demo/password/output"
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
		output.ErrorOutput("Ошибка в чтении файла!")
		return nil, err
	}
	return file, nil

}

func (db *JsonDb) Write(content []byte){
	file, err := os.Create(db.filename)
	if err != nil {
		output.ErrorOutput("Ошибка в создании файла!")
		return
	}
	defer file.Close()
	_, err = file.Write(content)
	if err != nil {
		file.Close()
		output.ErrorOutput("Не удалось записать данные!")
		return 
	}
	fmt.Println("Запись прошла успешно!")
	
}