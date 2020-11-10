package handlers

import (
	"log"
	"net/http"
	"fmt"
	"io/ioutil"
)


//create a hello Struct
type Hello struct{
	l *log.Logger
}

//function to create a new Hello
func NewHello(l *log.Logger) *Hello{
	return &Hello{l}
}

//serve mux is gonna call ServeHTTP function by default
func (h *Hello) ServeHTTP(rw http.ResponseWriter,r *http.Request){
	h.l.Println("Hello World")
	d,err := ioutil.ReadAll(r.Body)
		if err!=nil{
			http.Error(rw,"OOps ",http.StatusBadRequest)
			return
		}	
		fmt.Fprintf(rw,"hello %s",d)



}