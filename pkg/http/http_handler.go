package http

import (
	"net/http"

	"github.com/ptsgr/ImageGeneratorBot/pkg/image_creator"
)

func InitHandlers() http.Handler {
	ServeMux := http.NewServeMux()
	ServeMux.HandleFunc("/", imageHandler)

	return ServeMux
}

func imageHandler(w http.ResponseWriter, r *http.Request) {
	image_creator.CreateImage(w)
}
