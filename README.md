# A RESTful API with Go --> depot-in-go

A RESTful API project developed for

## Project Guide 
```bash
# Download this project
go get github.com/brcdmr/depot-in-go
```

```bash
# Build and Run
cd depot-in-go
make
// after running  -make- command, building will be started along with below build command 
// go build -ldflags "-X main.version=v1.0.0-7-g5222465-dirty" -o depot ./cmd/depot

./depot
// after running executable, listening port and version number can be seen

You can call with default port, or you can change it.

# Default API Endpoint : http://localhost:8888
```

## Project Structure
```
├── build
│   └── DockerFile
├── cmd
│   └── depot          
│       └── main.go   
├── pkg
│   ├── api          
│   │   ├── api.go   
│   │   ├── api_test.go  
│   │   ├── logging.go  
│   │   ├── response.go  
│   │   ├── api.out
│   └── repository          
│       ├── filesystem.go   
│       ├── filesystem_test.go  
│       ├── store.go  
│       ├── store_test.go  
│       └── repository.out
├── tmp
│   └── test.json
└── go.mod
└── Makefile
```


## API 
#### /getvalue/:key
* `GET` : Get a value from in memoryStore

#### /setvalue
* `POST` : Send key / value pair in request body to store data in memoryStore

#### /flush
* `GET` : Flush store data
