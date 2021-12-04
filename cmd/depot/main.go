package main

import (
	"depot/pkg/api"
	"depot/pkg/repository"
	"log"
	"net/http"
	"os"
)

//"github.com/brcdmr/depot-in-go/pkg/api"
//"github.com/brcdmr/depot-in-go/pkg/repository"

var version string

//http://localhost:8888/getvalue?key=burcu

type App struct {
	Port   string
	Repo   *repository.InMemoryStore
	Server api.ApiServer
}

func main() {

	a := App{}

	a.Port = os.Getenv("PORT")

	if a.Port == "" {
		a.Port = "8888"
	}

	a.initialize()
	//  myNewStore := repository.NewInMemoryStore()
	//  myApiServer := api.NewApiServer(myNewStore)
	a.routes()
	a.run()

	log.Printf("version %s listening on port %s", version, a.Port)
}

func (a *App) initialize() {
	a.Repo = repository.NewInMemoryStore()
	a.Server = api.NewApiServer(a.Repo)

}

func (a *App) routes() {
	http.HandleFunc("/getvalue", a.Server.GetItem().ServeHTTP)
	http.HandleFunc("/setvalue", a.Server.AddItem().ServeHTTP)
}

func (a *App) run() {

	err := http.ListenAndServe(":"+a.Port, nil)

	if err != nil {
		log.Fatal(err)
	}
}
