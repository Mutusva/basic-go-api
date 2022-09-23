package server

import (
	"context"
	"fmt"
	"go-api/api/v1/routes"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func RunServer(port string) {

	r := routes.InitialiseRoutes()

	s := &http.Server{
		Addr:         ":" + port,
		Handler:      r,
		ReadTimeout:  5 * time.Minute,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// graceful shut down - allows cleaning up of resources.
	go func() {
		fmt.Printf("Listening on port %s\n", port)
		err := s.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	log.Println("Received terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
