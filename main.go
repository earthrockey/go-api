package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Book struct {
	gorm.Model
	Firstname string
	Code      int
}

func getBookAll(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	db := dbConn()
	var books []Book
	db.Find(&books)
	json.NewEncoder(w).Encode(books)
}

func newBook(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	db := dbConn()
	decoder := json.NewDecoder(r.Body)
	var b Book
	err := decoder.Decode(&b)
	if err != nil {
		panic(err)
	}
	db.Create(&b)
	json.NewEncoder(w).Encode(b)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	fmt.Fprint(w, "Welcome to the HomePage!")
}

func dbConn() (db *gorm.DB) {
	db, err := gorm.Open("mysql", "root:e575g73wk@tcp(172.17.0.2:3306)/go_api?charset=utf8&parseTime=True&loc=Local")
	// db, err := gorm.Open("mysql", "root:e575g73wk@/go_api?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	return db
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/getBook", getBookAll)
	http.HandleFunc("/newBook", newBook)
	http.ListenAndServe(":8080", nil)
}

func main() {
	db := dbConn()
	db.AutoMigrate(&Book{})
	defer db.Close()
	handleRequest()
}
