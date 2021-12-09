package main

import (
	"depot/pkg/api"
	"depot/pkg/repository"
	"fmt"
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
	var Interval time.Duration = 1

	if Port == "" {
		Port = "8888"
	}

	// this flag is used to export .log file
	if LogToFile == "YES" {
		logFileName := "application.log"
		logFile, _ := os.OpenFile(logFileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
		defer logFile.Close()
		log.SetOutput(logFile)
	}

	a.initialize(Interval)
	a.routes()
	go a.startFileScheduler(Interval)
	a.run(Port)
	select {}

}

// Initialize the application
func (a *App) initialize(interval time.Duration) {

	dir, err := filepath.Abs("./../../")
	if err != nil {
		log.Fatal(err)
	}

	a.FileSys.Path = dir + "/tmp/"

	initialStoreData := a.getInitStoreData()
	a.Repo = repository.NewInMemoryStore(initialStoreData)
	a.Server = api.NewApiServer(a.Repo)

}

// Configures the routes definitions
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

// Runs the application
func (a *App) run(port string) {
	log.Printf("version %s listening on port %s", version, port)
	err := http.ListenAndServe(":"+port, api.LoggingMiddleware(a.Mux))

	if err != nil {
		log.Fatal(err)
	}
}

// Gets InitialStoreData,
// If there is a saved file within given format, returns fileData
func (a *App) getInitStoreData() map[string]string {

	found := a.FileSys.SearchSavedFileName()
	if found != "" {
		a.FileSys.Name = found
		return a.FileSys.ReadFile(a.FileSys.Name)
	}

	return make(map[string]string)
}

// Starts a scheduler for saving data to file in given interval (Minute)
func (a *App) startFileScheduler(duration time.Duration) {

	for range time.Tick(duration * time.Minute) {
		timeStamp := time.Now().Unix()

		newFileName := fmt.Sprintf("%d-data.json", timeStamp)

		log.Printf("%s - %s minutes have passed and called write file function!!", newFileName, duration.String())
		a.FileSys.WriteFile(a.Repo.GetAllStoreData(), newFileName)

		if len(a.FileSys.Name) > 0 {
			a.FileSys.RemoveFile(a.FileSys.Name)
		}

		a.FileSys.Name = newFileName

	}

}
