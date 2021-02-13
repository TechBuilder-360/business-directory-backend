package App

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func Start()  {
	router := mux.NewRouter()

	serverAddr := "localhost"
	serverPort := "8100"

	log.Printf("Server starting on %s:%s", serverAddr, serverPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", serverAddr, serverPort), router))
}