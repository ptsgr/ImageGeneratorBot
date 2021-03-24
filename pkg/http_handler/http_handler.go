package http_handler

import (
	"fmt"
	"log"
	"net/http"
)

func Run() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":9000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from ImageGeneratorServer!")
}
