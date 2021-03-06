package main

import (
	handlers "NicJackson/handlers"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	//create a log
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	//assigning a hello handler
	hh := handlers.NewHello(l)

	//assigning a goodbye handler
	gg := handlers.NewGoodbye(l)

	//assigning a products handler
	ph := handlers.NewProducts(l)
	//create a servemux
	sm := http.NewServeMux()

	sm.Handle("/", ph)
	sm.Handle("/helloWorld", hh)
	sm.Handle("/goodbye", gg)

	//creating a new server with properties
	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	//start server
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	//create a channel of signal
	sigchan := make(chan os.Signal)
	signal.Notify(sigchan, os.Interrupt)
	signal.Notify(sigchan, os.Kill)

	//assigning the channel to sig
	sig := <-sigchan
	l.Println("received instructions , Terminated Gracefully ", sig)

	//creating a context because the shutdown method accepts context as input
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)

	//shutting down
	s.Shutdown(tc)

}
