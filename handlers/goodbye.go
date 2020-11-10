package handlers

import(
	"log"
	"net/http"
	"io/ioutil"
	"fmt"

)

//create a goodbye struct
type Goodbye struct{
	l *log.Logger
}


//function to create a new good bye
func NewGoodbye(l *log.Logger) *Goodbye{
	return &Goodbye{l}
}


//serve mux defaultly call the ServeHTTP function
func (g *Goodbye) ServeHTTP(rw http.ResponseWriter,r *http.Request){
	g.l.Println("Bye World")
	d,error := ioutil.ReadAll(r.Body)
	if error!=nil {
		g.l.Println("error exist")
		fmt.Fprintf(rw,"OOPS ",http.StatusBadRequest)
		return
	}
	fmt.Fprintf(rw,"Bye %s",d)
	

}