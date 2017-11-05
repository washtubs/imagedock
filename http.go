package imagedock

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type HttpHandler struct {
	mux *mux.Router
}

func (h *HttpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	c := context.WithValue(context.Background(), contextKeys.requestId, nextRequestId())

	r = r.WithContext(c)

	log.Debugf("Request ID: %v", r.Context().Value(contextKeys.requestId))
	log.Debugf("Request method: %v", r.Method)
	log.Debugf("Request path: %v", r.URL.EscapedPath())

	h.mux.ServeHTTP(w, r)
}

func handleImageGET(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["imageId"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Expected id. Got "+idStr, http.StatusBadRequest)
		log.Errorf("Error parsing %v: %v", idStr, err)
		return
	}

	log.Debugf("Getting image %v", id)
}

func handleImagePOST(w http.ResponseWriter, r *http.Request) {
	p.provideTagModel(r.Context())
}

func NewHttpHandler() http.Handler {

	rtr := mux.NewRouter()

	// Images
	rtr.HandleFunc("/image/{imageId:[0-9]+}", handleImageGET).Methods("GET")
	rtr.HandleFunc("/image}", handleImagePOST).Methods("POST")

	return &HttpHandler{rtr}
}
