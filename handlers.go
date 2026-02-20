package main

import (
	"net/http"
	"strconv"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title: "Go + HTMX",
	}
	err := templates.ExecuteTemplate(w, "index", data)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}

func ShowAddProductForm(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "add-product", nil)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	price, _ := strconv.ParseFloat(r.FormValue("price"), 64)
	stock, _ := strconv.Atoi(r.FormValue("stock"))

	product := Products{
		Name:  r.FormValue("name"),
		Price: price,
		Stock: stock,
	}

	db.Create(&product)

	w.Header().Set("HX-Trigger", "productAdded")

	w.Write([]byte(`
    <div class="alert alert-success py-2 shadow-sm border-none">
        <svg xmlns="http://www.w3.org/2000/svg" class="stroke-current shrink-0 h-6 w-6" fill="none" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" /></svg>
        <span class="text-sm font-medium">Saved to Database!</span>
    </div>
`))
}

func GetProductHandler(w http.ResponseWriter, r *http.Request) {
	var products []Products

	db.Find(&products)

	err := templates.ExecuteTemplate(w, "products-lists", products)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}

func DeleteProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	result := db.Delete(&Products{}, id)

	if result.Error != nil {
		http.Error(w, "Could not delete item", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
