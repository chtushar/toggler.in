package handlers

import (
	"net/http"
	"regexp"
	"runtime/debug"

	"go.uber.org/zap"
)

type recoveryHandler struct {
	handler    http.Handler
	log        *zap.Logger
	// jsonWriter *response.JSONWriter
	apiRegex   *regexp.Regexp
}

// RecoveryHandler is HTTP middleware that recovers from a panic,
// logs the panic, writes http.StatusInternalServerError, and
// continues to the next handler.
func RecoveryHandler(log *zap.Logger) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		r := &recoveryHandler{
			handler:  h,
			log:      log,
			apiRegex: regexp.MustCompile(`^/api/`),
		}

		return r
	}
}


func (h *recoveryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer h.recoverPanic(w, r)
	h.handler.ServeHTTP(w, r)
}

func (h *recoveryHandler) recoverPanic(w http.ResponseWriter, r *http.Request) {
	err := recover()
	if err == nil {
		return
	}

	h.logPanic(r, err)

	if h.apiRegex.MatchString(r.URL.Path) {
		// h.jsonResponse(w, r)
		return
	}

	// h.htmlResponse(w, r)
}

// func (h *recoveryHandler) htmlResponse(w http.ResponseWriter, _ *http.Request) {
// 	w.WriteHeader(http.StatusInternalServerError)
// }

// func (h *recoveryHandler) jsonResponse(w http.ResponseWriter, r *http.Request) {
// 	h.jsonWriter.DefaultError(w, r)
// }

func (h *recoveryHandler) logPanic(_ *http.Request, v interface{}) {
	h.log.Error(
		"server panic recovered",
		zap.Stack("stack"),
		zap.Any("error", v),
	)

	debug.PrintStack()
}

