package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"

	"github.com/bobTheBuilder7/bunrouter"
	"github.com/bobTheBuilder7/bunrouter/extra/reqlog"
)

func main() {
	router := bunrouter.New(
		bunrouter.Use(reqlog.NewMiddleware()),
	)

	group := router.Use(errorHandler)
	group.GET("/", indexHandler)

	log.Println("listening on http://localhost:9999")
	log.Println(http.ListenAndServe(":9999", router))
}

func indexHandler(w http.ResponseWriter, req bunrouter.Request) error {
	rnd := rand.Float64()
	if rnd < 0.3 {
		return io.EOF
	}
	if rnd < 0.6 {
		return sql.ErrNoRows
	}
	return fmt.Errorf("hello world")
}

func errorHandler(next bunrouter.HandlerFunc) bunrouter.HandlerFunc {
	return func(w http.ResponseWriter, req bunrouter.Request) error {
		err := next(w, req)

		switch err := err.(type) {
		case nil:
			// no error
		case HTTPError:
			w.WriteHeader(err.statusCode)
			_ = bunrouter.JSON(w, err)
		default:
			httpErr := NewHTTPError(err)
			w.WriteHeader(httpErr.statusCode)
			_ = bunrouter.JSON(w, httpErr)
		}

		return err
	}
}

type HTTPError struct {
	statusCode int

	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e HTTPError) Error() string {
	return e.Message
}

func NewHTTPError(err error) HTTPError {
	switch err {
	case io.EOF:
		return HTTPError{
			statusCode: http.StatusBadRequest,

			Code:    "eof",
			Message: "EOF reading HTTP request body",
		}
	case sql.ErrNoRows:
		return HTTPError{
			statusCode: http.StatusNotFound,

			Code:    "not_found",
			Message: "Page Not Found",
		}
	}

	return HTTPError{
		statusCode: http.StatusInternalServerError,

		Code:    "internal",
		Message: "Internal server error",
	}
}
