package main

import (
	"fmt"
	"html/template"
	"net/http"

	"gorm.io/gorm"
)

type PageData struct {
	Title   string
	Message string
}

var db *gorm.DB
var templates *template.Template

func init() {
	templates = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	db = InitDB()
	fs := http.FileServer(http.Dir("./static/"))

	http.Handle("/static/", http.StripPrefix("/static", fs))

	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/add-product-form", ShowAddProductForm)
	http.HandleFunc("GET /get-products", GetProductHandler)

	http.HandleFunc("POST /add-product", CreateProduct)
	http.HandleFunc("DELETE /delete-product/{id}", DeleteProductHandler)

	fmt.Println("Server running at localhost:8080")
	http.ListenAndServe(":8080", nil)
}
