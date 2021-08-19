package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"productAPI/data"
	middleware "productAPI/midddleware"
	"strconv"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l: l}
}

func (p *Products) GetProducts(w http.ResponseWriter, r *http.Request){
	p.l.Println("Handling GET Request to List Products")
	d := data.GetProducts()
	err := d.ToJSON(w)
	if err != nil {
		http.Error(w, "error marshalling data", http.StatusBadRequest)
		return
	}
}

func (p *Products) AddProduct(w http.ResponseWriter, r *http.Request){
	p.l.Println("Handling Post Request to ADD a Product")
	key := middleware.ProductKey
	product := r.Context().Value(key).(data.Product)
	data.AddProduct(&product)
	fmt.Println(data.ProductList)
}

func (p *Products) UpdateProduct(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w,"Error in retrieving the ID from URL", http.StatusBadRequest)
		return
	}
	p.l.Println("Handling PUT Request")
	product := r.Context().Value(middleware.ProductKey).(data.Product)

	err = data.UpdateProduct(id, &product)
	if err != nil {
		http.Error(w,"Product not found", http.StatusNotFound)
		return
	}


}

func (p *Products) DeleteProduct(w http.ResponseWriter, r *http.Request) {

}