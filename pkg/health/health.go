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
	fmt.Fprint(w, common.GetReponse())
}

func GetStatusHandler() (handler http.Handler) {
	handler = &Handler{}

	return handler
}

func AddStatusHandler(router *mux.Router, contextPath string) {
	if contextPath == "" {
		contextPath = common.ContextPath
	}

	router.Methods(common.Get).
		Path(fmt.Sprintf(common.APIContextPattern, contextPath, common.GetAPIHealthStatusUrl())).
		Name(common.APIHealthStatusUrl).
		Handler(GetStatusHandler())
}
