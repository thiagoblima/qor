package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/qor/qor/admin"
	"github.com/qor/qor/resource"

	"net/http"
)

type User struct {
	Id   int64
	Name string
	Role string
}

var db gorm.DB

func init() {
	db, _ = gorm.Open("sqlite3", "/tmp/qor.db")
	db.AutoMigrate(&User{})

	db.FirstOrCreate(&User{}, User{Name: "jinzhu", Role: "admin"})
	db.FirstOrCreate(&User{}, User{Name: "juice", Role: "dev"})
}

func main() {
	mux := http.NewServeMux()

	user := resource.New(&User{})
	admin := admin.New(&db)
	admin.AddResource(user)
	admin.AddToMux("/admin", mux)

	fmt.Println("listening on :8080")
	http.ListenAndServe(":8080", mux)
}
