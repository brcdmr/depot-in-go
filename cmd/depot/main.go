package main

import (
	"depot/pkg/api"
	"depot/pkg/repository"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

var version string

type App struct {
	Repo    *repository.InMemoryStore
	Server  api.ApiServer
	FileSys repository.FileSystem
	Mux     *http.ServeMux
}

func main() {
	a := App{}

	var Port string = os.Getenv("PORT")
	var LogToFile string = os.Getenv("WRITELOGFILE")
	var Interval time.Duration = 30

	if Port == "" {
		Port = "8888"
	}

	if LogToFile == "YES" {
		logFileName := "application.log"
		logFile, _ := os.OpenFile(logFileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
		defer logFile.Close()
		log.SetOutput(logFile)
	}

	a.initialize()

	a.routes()
	go a.startFileScheduler(Interval)
	a.run(Port)
	select {}

}

func (a *App) initialize() {
	dir, err := filepath.Abs("./../../")
	if err != nil {
		log.Fatal(err)
	}

	a.FileSys.Name = "test.json"
	a.FileSys.Path = dir + "/tmp/"

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
	// a.Mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "index.html")
	// })
	a.Mux.HandleFunc("/", a.Server.Home().ServeHTTP)
	a.Mux.HandleFunc("/getvalue", a.Server.GetValue().ServeHTTP)
	a.Mux.HandleFunc("/setvalue", a.Server.AddValue().ServeHTTP)
	a.Mux.HandleFunc("/flush", a.Server.Flush().ServeHTTP)
}

func (a *App) run(port string) {
	log.Printf("version %s listening on port %s", version, port)
	err := http.ListenAndServe(":"+port, api.LoggingMiddleware(a.Mux))

	if err != nil {
		log.Fatal(err)
	}
}

func (a *App) startFileScheduler(duration time.Duration) {

	for range time.Tick(duration * time.Second) { // to do: change to min
		// timeStamp := time.Now().Unix()

		// a.FileSys.Name = fmt.Sprintf("%d-data.json", timeStamp)
		log.Printf("%s - %s minutes have passed and called write file function!!", a.FileSys.Name, duration.String())
		a.FileSys.WriteFile(a.Repo.GetAllStoreData())
	}
	// time.AfterFunc(duration, func() {
	// 	a.FileSys.WriteFile(a.Repo.GetAllStoreData())
	// })
}
