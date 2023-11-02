# Go-Swagger-UI
Swagger UI tool for Go-Micro and more.

## How to use
At first, you need to install this lib:

```shell
go get -u github.com/devexps/go-swagger-ui
```

### Directly use
```go
import (
    "net/http"

    swaggerUI "github.com/devexps/go-swagger-ui"
)

func main() {
    swaggerHandler := swaggerUI.New(
        "Petstore",
        "https://petstore3.swagger.io/api/v3/openapi.json",
        "/docs/",
    )

    http.Handle("/docs/", swaggerHandler)
    http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
        _, _ = writer.Write([]byte("Hello World!"))
    })

    println("docs at http://localhost:8080/docs/")
    _ = http.ListenAndServe("localhost:8080", http.DefaultServeMux)
}
```