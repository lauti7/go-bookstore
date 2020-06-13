# Go-Bookstore API 
## Bookstore API built with Go Lang, Gin-Gonic Framework And Gorm ORM

### To Run API
* `git clone https://github.com/lauti7/go-bookstore.git` in your $GOPATH.
* Download API Dependencies:
  * `go get github.com/gin-gonic/gin`
  * `go get github.com/jinzhu/gorm`
* Turn on your MySql Server
* In "go-bookstore/models/setup.go" remove my Mysql Localhost Server URL and type your URL
* Crate Database in Mysql.
* Inside project folder, run: `go run main.go`
* Open `127.0.0.1:8080/api/books`
