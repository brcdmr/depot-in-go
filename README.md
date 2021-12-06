# A RESTful API with Go --> depot-in-go

## Project Guide 
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
