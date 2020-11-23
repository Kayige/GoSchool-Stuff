package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"restapi/routes"

	"github.com/joho/godotenv"

	"time"
)

func Start() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("API starting....")

	port := os.Getenv("API_PORT")

	r := routes.New()

	srv := &http.Server{
		Handler:      r.Router(),
		Addr:         "127.0.0.1:" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("API server listening on :%v", port)
	log.Fatal(srv.ListenAndServe())
}
