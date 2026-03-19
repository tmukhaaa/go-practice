package cloud

type CloudDb struct {
	cloudDb string
}

func NewCloudDb(url string) *CloudDb {
	return &CloudDb{
		cloudDb: url,
	}
}

func (db *CloudDb) Read() ([]byte, error) {
	return nil, nil

}

func (db *CloudDb) Write(content []byte){
	
}