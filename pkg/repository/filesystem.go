package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
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

	_, err := os.Stat(fs.Path + fs.Name)
	return !os.IsNotExist(err)
}

func (fs *FileSystem) WriteFile(data map[string]string, fileName string) string {
	dataToJson, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("File data marshal error: %s %v", data, err)
	}
	writeErr := ioutil.WriteFile(fs.Path+fileName, dataToJson, 0777)

	if writeErr != nil {
		log.Fatalf("File system writing error: %s %v", dataToJson, err)
	}

	return fmt.Sprintf("File was saved: %s!", fs.Path)

}

func (fs *FileSystem) ReadFile(fileName string) map[string]string {
	fileDataBytes, err := ioutil.ReadFile(fs.Path + fileName)

	if err != nil {
		log.Fatalf("File reading error: %s %v", fs.Path, err)

	}

	return fs.convertFileData(fileDataBytes)
}

func (fs *FileSystem) RemoveFile(fileName string) error {
	err := os.Remove(fs.Path + fileName)

	if err != nil {
		return err
		//log.Fatalf("File remove error: %s %v", fs.Path, err)
	}
	return nil
}

func (fs *FileSystem) convertFileData(fileData []byte) map[string]string {
	var convertedData map[string]string
	err := json.Unmarshal(fileData, &convertedData)
	if err != nil {
		log.Fatalf("File data convert error: %s %v", fs.Path, err)
	}
	return convertedData
}

func (fs *FileSystem) SearchSavedFileName() string {
	var files []string

	fileNameRegEx, _ := regexp.Compile("^.*-data.json$")
	fileNames, _ := ioutil.ReadDir(fs.Path)

	for _, fl := range fileNames {

		if fileNameRegEx.MatchString(fl.Name()) {
			files = append(files, fl.Name())

		}
	}

	if len(files) == 1 {
		return files[0]
	}
	return ""

}
