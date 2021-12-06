package main

import (
	"depot/pkg/api"
	"depot/pkg/repository"
	"log"
	"net/http"
	"os"
	"time"
)

//"github.com/brcdmr/depot-in-go/pkg/api"
//"github.com/brcdmr/depot-in-go/pkg/repository"

var version string

type App struct {
	Port      string
	LogToFile string
	Repo      *repository.InMemoryStore
	Server    api.ApiServer
	FileSys   repository.FileSystem
	Mux       *http.ServeMux
}

func main() {

	a := App{}

	a.Port = os.Getenv("PORT")
	a.LogToFile = os.Getenv("WRITELOGFILE")

	if a.Port == "" {
		a.Port = "8888"
	}

	if a.LogToFile == "YES" {

		logFileName := "application.log"
		logFile, _ := os.OpenFile(logFileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
		defer logFile.Close()

		log.SetOutput(logFile)

	}

	a.initialize()

	a.routes()
	go a.startFileScheduler(10)
	a.run()

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
	a.Mux = http.NewServeMux()
	//Mux := http.NewServeMux()
	a.Mux.HandleFunc("/getvalue", a.Server.GetValue().ServeHTTP)
	a.Mux.HandleFunc("/setvalue", a.Server.AddValue().ServeHTTP)
	a.Mux.HandleFunc("/flush", a.Server.Flush().ServeHTTP)
}

func (a *App) run() {

	log.Printf("version %s listening on port %s", version, a.Port)
	err := http.ListenAndServe(":"+a.Port, api.LoggingMiddleware(a.Mux))

	if err != nil {
		log.Fatal(err)
	}

}

func (a *App) startFileScheduler(duration time.Duration) {

	for range time.Tick(duration * time.Second) {
		// str := "Hi! " + duration.String() + " seconds have passed and file saved!!"
		// echo(str)
		log.Printf("Hi! %s seconds have passed and called write file function!!", duration.String())
		a.FileSys.WriteFile(a.Repo.GetAllStoreData())
	}
	// time.AfterFunc(duration, func() {
	// 	a.FileSys.WriteFile(a.Repo.GetAllStoreData())
	// })
}

// func echo(s string) {
// 	fmt.Println(s)
// }
