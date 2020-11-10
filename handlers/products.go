package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"
	"usercreated/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	//for get method
	if r.Method == http.MethodGet {
		p.getProduct(rw, r)
		return
	}

	//for post methods
	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
		return
	}

	if r.Method == http.MethodPut {
		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(r.URL.Path, -1)

		p.l.Println(" g is ", g)
		if len(g) != 1 {
			http.Error(rw, "bad URL", http.StatusBadRequest)
			return
		}

		if len(g[0]) != 2 {
			p.l.Println("Invalid URI more than one capture group")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		idString := g[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			p.l.Println("Invalid URI unable to convert to numer", idString)
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		p.updateProducts(id, rw, r)
		return

	}

	rw.WriteHeader(http.StatusMethodNotAllowed)

}

func (p *Products) getProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Printf("Handle Get Products")
	lp := data.GetProducts()
	err := lp.ToJson(rw)
	if err != nil {
		http.Error(rw, "unable to convert", http.StatusBadRequest)
	}
}

//adding product
func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {

	p.l.Println("Handle Post Product")
	product := &data.Product{}
	err := product.FromJson(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal", http.StatusBadRequest)
		return
	}
	p.l.Printf("the value is %#v", product)
	//adding product to the list
	data.AddProduct(product)
}

func (p Products) updateProducts(id int, rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT Product")

	prod := &data.Product{}

	err := prod.FromJson(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	err = data.UpdateProduct(id, prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}
}
