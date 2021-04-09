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
	httpKeys := r.URL.Query()
	log.Println(httpKeys)
	img := new(image_creator.Image)
	img.Properties.InitImageProperties(httpKeys)

	buffer, err := img.CreateImage()
	if err != nil {
		log.Printf("Error generate image: %s", err.Error())
	}
	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))

	if _, err := w.Write(buffer.Bytes()); err != nil {
		log.Printf("Error return image from byte buffer: %s", err.Error())
	}

}
