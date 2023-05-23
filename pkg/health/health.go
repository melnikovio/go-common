package health

import (
	"fmt"
	"net/http"

	"github.com/melnikovio/go-common/pkg/common"

	"github.com/gorilla/mux"
)


type Handler struct {

}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi")
}

func GetStatusHandler() (handler http.Handler) {
	handler = &Handler{}

	return handler
}

func AddStatusHandler(router *mux.Router) {
	router.Methods(common.Get).
		Path(fmt.Sprintf(common.APIContextPattern, common.ContextPath, common.APIHealthStatusUrl)).
		Name(common.APIHealthStatusUrl).
		Handler(GetStatusHandler())
}