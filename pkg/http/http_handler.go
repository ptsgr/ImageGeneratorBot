package http

import (
	"log"
	"net/http"

	"github.com/ptsgr/ImageGeneratorBot/pkg/image_creator"
)

func InitHandlers() http.Handler {
	ServeMux := http.NewServeMux()
	ServeMux.HandleFunc("/", imageHandler)

	return ServeMux
}

func imageHandler(w http.ResponseWriter, r *http.Request) {
	img := new(image_creator.Image)
	if err := img.CreateImage(w); err != nil {
		log.Fatalf("Error imageGenerator running http server: %s", err.Error())
	}
}
