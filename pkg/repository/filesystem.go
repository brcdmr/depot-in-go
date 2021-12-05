package repository

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type FileSystem struct {
	Path string
	Name string
	//lock sync.Mutex
}

func (fs *FileSystem) IsFileExist() bool {
	if fs.Path == "" {
		return false
	}

	_, err := os.Stat(fs.Path)
	return !os.IsNotExist(err)
}

func (fs *FileSystem) WriteFile(data map[string]string) {
	dataToJson, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("File data marshal error: %s %v", data, err)
	}
	err = os.WriteFile(fs.Path, dataToJson, 0777)
	if err != nil {
		log.Fatalf("File system writing error: %s %v", dataToJson, err)
	}
}

func (fs *FileSystem) ReadFile() map[string]string {
	fileDataBytes, err := ioutil.ReadFile(fs.Path)

	if err != nil {
		log.Fatalf("File reading error: %s %v", fs.Path, err)

	}

	return fs.convertFileData(fileDataBytes)
}

func (fs *FileSystem) convertFileData(fileData []byte) map[string]string {
	var convertedData map[string]string
	err := json.Unmarshal(fileData, &convertedData)
	if err != nil {
		log.Fatalf("File data convert error: %s %v", fs.Path, err)
	}
	return convertedData
}

/*
func (fs *FileSystem) ReadFile() map[string]string {
	osfile, err := os.OpenFile(fs.Name, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("File opening error: %s %v", fs.Name, err)

	}

	fileDataBytes, err := ioutil.ReadAll(osfile)
	if err != nil {
		log.Fatalf("File reading error: %s %v", fs.Name, err)

	}

	return fs.convertFileData(fileDataBytes)
}
*/
