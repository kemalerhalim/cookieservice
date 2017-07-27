# cookieservice.go

Simple cookie service made for golang.

# Installation
### First, get the package
```
go get github.com/kemalerhalim/cookieservice
```

### Then import it
```
import(
    "github.com/kemalerhalim/cookieservice"
)
```
# Usage

  - Getting value of a cookie:
```go
        func HomePage(w http.ResponseWriter,req *http.Request){
            var value=cookieservice.ReadCookie(req,"cookieName")
            fmt.Println(value)
            ...
        }
```
  - Creating a cookie (Which will expire after 48 hours):
```go
        func HomePage(w http.ResponseWriter,req *http.Request){
            cookieservice.NewCookie(w,"cookieName","cookieValue",48)
            ...
        }
```
  - Deleting a cookie:
```go
        func HomePage(w http.ResponseWriter,r *http.Request){
            cookieservice.DeleteCookie(w,"cookieName")
            ...
        }
```


