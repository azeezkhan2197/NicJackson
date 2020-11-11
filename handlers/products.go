package handlers

import (
	data "NicJackson/data"
	"context"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Printf("Handle Get Products")
	lp := data.GetProducts()
	err := lp.ToJson(rw)
	if err != nil {
		http.Error(rw, "unable to convert", http.StatusBadRequest)
	}
}

//adding product
func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {

	p.l.Println("Handle Post Product")

	//taking value using middleware function
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	p.l.Printf("the value is %#v", prod)
	//adding product to the list
	data.AddProduct(&prod)
}

//updating product
func (p Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT Product")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		p.l.Println("error is returned while converting from string to int")
	}

	//taking value from middle ware
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	err = data.UpdateProduct(id, &prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}

type KeyProduct struct{}

//middleWare
func (p Products) MiddleWareProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := data.Product{}

		err := prod.FromJson(r.Body)
		if err != nil {
			http.Error(rw, "Unable to unmarshal", http.StatusBadRequest)
			return
		}
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(rw, r)
	})

}
