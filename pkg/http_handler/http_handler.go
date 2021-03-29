package http_handler

import (
	"log"
	"net/http"

	"github.com/ptsgr/ImageGeneratorBot/pkg/image_creator"
)

func Run() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":9000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	image_creator.CreateImage(w)
}
