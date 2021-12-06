package main

import (
	"depot/pkg/api"
	"depot/pkg/repository"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

//"github.com/brcdmr/depot-in-go/pkg/api"
//"github.com/brcdmr/depot-in-go/pkg/repository"

var version string

//http://localhost:8888/getvalue?key=burcu

type App struct {
	Port    string
	Repo    *repository.InMemoryStore
	Server  api.ApiServer
	FileSys repository.FileSystem
}

func main() {

	a := App{}

	a.Port = os.Getenv("PORT")

	if a.Port == "" {
		a.Port = "8888"
	}

	a.initialize()

	a.routes()
	go a.startFileScheduler(10)
	a.run()
	log.Printf("version %s listening on port %s", version, a.Port)
	select {}

}

func (a *App) initialize() {
	// testpath, err := filepath.Abs(filepath.Dir("test.json"))

	// fmt.Print(testpath)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	a.FileSys.Name = "test.json"
	a.FileSys.Path = "/Users/burcudemirel/Desktop/Go/depot-in-go/tmp/test.json"

	isExist := a.FileSys.IsFileExist()

	if isExist {
		storeData := a.FileSys.ReadFile()
		a.Repo = repository.NewInMemoryStore(storeData)
	} else {
		a.Repo = repository.NewInMemoryStore(make(map[string]string))
	}

	a.Server = api.NewApiServer(a.Repo)

}

func (a *App) routes() {
	http.HandleFunc("/getvalue", a.Server.GetValue().ServeHTTP)
	http.HandleFunc("/setvalue", a.Server.AddValue().ServeHTTP)
	http.HandleFunc("/flush", a.Server.Flush().ServeHTTP)
}

func (a *App) run() {

	err := http.ListenAndServe(":"+a.Port, nil)

	if err != nil {
		log.Fatal(err)
	}

}

func (a *App) startFileScheduler(duration time.Duration) {

	for range time.Tick(duration * time.Second) {
		str := "Hi! " + duration.String() + " seconds have passed"
		echo(str)
		a.FileSys.WriteFile(a.Repo.GetAllStoreData())
	}
	// time.AfterFunc(duration, func() {
	// 	a.FileSys.WriteFile(a.Repo.GetAllStoreData())
	// })
}

func echo(s string) {
	fmt.Println(s)
}
