package main

import (
	"context"
	"crm/middleware"
	"encoding/json"
	"fmt"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	mux := http.NewServeMux()


	mux.Handle("/api/messages", middleware.VerifyJwtTokenMiddleware(http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

		for k, v := range request.Header {
			fmt.Printf("k : %s, v : %+v \n", k, v)
		}

		var res = "{\"de\":\"ddd\"}"
		resbytes, _ := json.Marshal(res)
		_, err := writer.Write(resbytes)
		if err != nil {
			log.Println(err)
		}
	})))
	// verify jwt token





	srv := &http.Server{
		Addr:    "localhost:9001",
		Handler: cors.New(cors.Options{
			AllowedOrigins: []string{"https://localhost:*"},
			AllowedHeaders: []string{"Authorization"},
			AllowCredentials: true,

		}).Handler(mux),
	}

	log.Printf("starting server on %s \n", srv.Addr)

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()


	var quit = make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, os.Kill)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
	log.Println("server exit")

}
