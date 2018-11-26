package main

import (
	"log"
	"net/http"

	"github.com/sachinnay/Gophercises/assignment18/transform/handler"
)

var listenAndServeFunc = http.ListenAndServe

//Start point of application
func main() {
	log.Fatal(listenAndServeFunc(":3003", handler.GetMux()))
}
