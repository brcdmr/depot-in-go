# A RESTful API with Go --> depot-in-go

## Download & Run
```bash
# Download this project
go get github.com/brcdmr/depot-in-go
```

```bash
# Build and Run
cd depot-in-go
go build
./depot-int-go

You can call with default port, or you can change it.

# API Endpoint : http://localhost:8888

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