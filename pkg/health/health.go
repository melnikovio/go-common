package health

import (
	"fmt"
	"net/http"
)


type Handler struct {

}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi")
}

func StatusHandler() (handler http.Handler) {
	handler = &Handler{}

	return handler
}

