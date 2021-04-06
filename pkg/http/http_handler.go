package http

import (
	"html"
	"log"
	"net/http"

	"github.com/ptsgr/ImageGeneratorBot/pkg/image_creator"
)

func InitHandlers() http.Handler {
	ServeMux := http.NewServeMux()
	ServeMux.HandleFunc("/", imageHandler)
	ServeMux.HandleFunc("/favicon.ico", faviconHandler)
	return ServeMux
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
}

func imageHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Hello, %q\n", html.EscapeString(r.URL.Path))
	img := new(image_creator.Image)
	if err := img.CreateImage(w); err != nil {
		log.Fatalf("Error imageGenerator running http server: %s", err.Error())
	}
}
