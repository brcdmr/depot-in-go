package repository_test

import (
	"depot/pkg/repository"
	"testing"
)

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
