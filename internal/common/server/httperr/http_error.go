package httperr

import (
	"github.com/google/jsonapi"
	"github.com/our-edu/new-go-commin/internal/common/errors"
	"github.com/our-edu/new-go-commin/internal/common/logs"
	"net/http"
	"strconv"
)

func InternalError(slug string, err error, w http.ResponseWriter, r *http.Request) {
	httpRespondWithError(err, slug, w, r, "Internal server error", http.StatusInternalServerError)
}

func Unauthorised(slug string, err error, w http.ResponseWriter, r *http.Request) {
	httpRespondWithError(err, slug, w, r, "Unauthorised", http.StatusUnauthorized)
}

func BadRequest(slug string, err error, w http.ResponseWriter, r *http.Request) {
	httpRespondWithError(err, slug, w, r, "Bad request", http.StatusBadRequest)
}

func RespondWithSlugError(err error, w http.ResponseWriter, r *http.Request) {
	slugError, ok := err.(errors.SlugError)
	if !ok {
		InternalError("internal-server-error", err, w, r)
		return
	}
	switch slugError.ErrorType() {
	case errors.ErrorTypeAuthorization:
		Unauthorised(slugError.Slug(), slugError, w, r)
	case errors.ErrorTypeIncorrectInput:
		BadRequest(slugError.Slug(), slugError, w, r)
	default:
		InternalError(slugError.Slug(), slugError, w, r)
	}
}

func httpRespondWithError(err error, slug string, w http.ResponseWriter, r *http.Request, logMSg string, status int) {
	logs.GetLogEntry(r).WithError(err).WithField("error-slug", slug).Warn(logMSg)
	//resp := ErrorResponse{slug, status}
	w.Header().Set("Content-Type", "application/vnd.api+json")
	w.WriteHeader(status)

	jsonapi.MarshalErrors(w, []*jsonapi.ErrorObject{{
		Title:  logMSg,
		Detail: err.Error(),
		Status: strconv.Itoa(status),
	}})
	//if err := render.Render(w, r, resp); err != nil {
	//	panic(err)
	//}
	return
}

type ErrorResponse struct {
	Slug       string `json:"slug"`
	httpStatus int
}

func (e ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(e.httpStatus)
	return nil
}
