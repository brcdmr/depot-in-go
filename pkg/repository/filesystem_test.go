package repository_test

import (
	"depot/pkg/repository"
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

type StubFileSystem struct {
	Path string
	Name string
	File *os.File
}

func createTempFile(t testing.TB, initialData []byte) (StubFileSystem, func()) {
	dir, _ := filepath.Abs("./")

	tmp := StubFileSystem{}

	tmp.Path = dir

	var err error
	tmp.File, err = ioutil.TempFile(tmp.Path, "*.json")
	if err != nil {
		t.Fatalf("Create temp file err %v", err)
	}
	tmp.Path = tmp.File.Name()
	tmp.File.Write([]byte(initialData))
	t.Helper()
	RemoveFile := func() {
		tmp.File.Close()
		os.Remove(tmp.File.Name())
	}

	return tmp, RemoveFile
	//ioutil.WriteFile(tmp.Path, json.Marshal(initData), 0777)

}

func Test_ReadFile(t *testing.T) {

	t.Run(t.Name(), func(t *testing.T) {

		want := map[string]string{
			"key1": "value1",
			"key2": "value2",
		}

		testFile, cleanTestFile := createTempFile(t, convertToJson(want))
		defer cleanTestFile()

		fileSys := &repository.FileSystem{Path: testFile.Path, Name: testFile.Name}

		got := fileSys.ReadFile()

		equal := reflect.DeepEqual(want, got)
		if !equal {
			t.Fatalf("ReadFile error: %s", convertToJson(got))
		}
		//testFile.File.RemoveFile()
	})
}

func convertToJson(data map[string]string) []byte {

	converted, _ := json.Marshal(data)

	return converted
}

func Test_IsFileExist_True(t *testing.T) {

	t.Run(t.Name(), func(t *testing.T) {
		fileSys := &repository.FileSystem{Path: "/Go/depot-in-go/tmp/test.json", Name: "test.json"}

		got := fileSys.IsFileExist()
		want := false

		if got != want {
			t.Fatalf("fileExist should return %t, but got value %t", want, got)
		}

	})
}
func Test_IsFileExist_False(t *testing.T) {

	t.Run(t.Name(), func(t *testing.T) {
		fileSys := &repository.FileSystem{Path: "", Name: ""}

		got := fileSys.IsFileExist()
		want := false

		if got != want {
			t.Fatalf("fileExist should return %t, but got value %t", want, got)
		}
	})
}

// func Test_WriteFile(t *testing.T) {
// 	t.Run(t.Name(), func(t *testing.T) {
// 		fileSys := &repository.FileSystem{Path: "/testfile.json", Name: ""}
// 		fileData := map[string]string{
// 			"hello": "file",
// 		}
// 		fileSys.WriteFile(fileData)
// 		got := fileSys.IsFileExist()

// 		if !got {
// 			t.Fatalf("Write file error")
// 		}

// 	})
// }
