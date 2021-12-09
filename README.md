# A RESTful API with Go --> depot-in-go

A RESTful API project developed with Go which has below properties / functions:
--> https://depotingo.herokuapp.com/

    * deployed on Heroku and linked to GitHub (main) for automatic deployment
    * can be accessed via https://depotingo.herokuapp.com/
    * three different endpoints available (getvalue / setvalue / flush)
    * thread safe InMemoryStorage
    * InMemoryStorage data is saved at given interval
    * httpServer with LoggingMiddleware, enables exporting .log file
    * has DockerFile 
    * has MakeFile with Golang Lint 
    * code coverage documentations are generated 
        - apiPackage Test Coverage, for html view run -> go tool cover -html=apitestcoverage.out
        - repositoryPackage Test Coverage, for html view run -> go tool cover -html=repository.out



## Project Guide 
```bash
# Download this project
go get github.com/brcdmr/depot-in-go
```

```bash
# Build and Run
cd depot-in-go
make
--> after running  -make- command, building will be started along with below build command 
--> go build -ldflags "-X main.version=v1.0.0-7-g5222465-dirty" -o depot ./cmd/depot

./depot
--> after running executable, listening port and version number can be seen

make docker
--> deploy to docker

You can call with default port, or you can change it.

# Default API Endpoint : http://localhost:8888
```

## Project Layout
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
│   │   ├── apiDoc.out  
│   │   └── apitestcoverage.out 
│   └── repository          
│       ├── filesystem.go   
│       ├── filesystem_test.go  
│       ├── store.go  
│       ├── store_test.go  
│       ├── repositoryDoc.out  
│       └── repositorytestcoverage.out 
├── tmp
│   └── samples.json
└── go.mod
└── go.sum
└── Makefile
```


## API 
#### /getvalue/:key
* `GET` : Get a value from in memoryStore

#### /setvalue
* `POST` : Send key / value pair in request body to store data in memoryStore

```
# Sample Post Body

    {
        "Key":"Paris",
        "Value":"France"
    }

```

#### /flush
* `GET` : Flush store data



