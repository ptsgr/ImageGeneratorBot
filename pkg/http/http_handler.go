package http

import (
	"log"
	"net/http"
	"strconv"

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
	img := new(image_creator.Image)
	img.Properties.InitImageProperties()

	buffer, err := img.CreateImage()
	if err != nil {
		log.Fatalf("Error generate image: %s", err.Error())
	}
	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))

	if _, err := w.Write(buffer.Bytes()); err != nil {
		log.Fatalf("Error return image from byte buffer: %s", err.Error())
	}

}
