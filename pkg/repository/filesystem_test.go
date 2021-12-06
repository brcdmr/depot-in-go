package repository_test

import (
	"depot/pkg/repository"
	"testing"
)

func Test_IsFileExist(t *testing.T) {

	t.Run(t.Name(), func(t *testing.T) {
		fileSys := &repository.FileSystem{Path: "/Go/depot-in-go/tmp/test.json", Name: "test.json"}

		got := fileSys.IsFileExist()
		want := false

		if got != want {
			t.Fatalf("fileExist should return %t, but got value %t", want, got)
		}

	})
}

// func Test_WriteFile_MarshalError(t *testing.T) {

// 	t.Run(t.Name(), func(t *testing.T) {
// 		fileSys := &repository.FileSystem{Path: "/Go/depot-in-go/tmp/test.json", Name: "test.json"}
// 		want := map[string]string{
// 			"test1":"testvak1",
// 		}
// 		fileSys.WriteFile(textData)
// 		got := fileSys.ReadFile(fileSys.Path)

// 		if got != want {
// 			t.Fatalf("fileExist should return %t, but got value %t", want, got)
// 		}

// 	})
// }
