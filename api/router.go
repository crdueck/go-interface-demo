package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type KeyValueService interface {
	GetString(int64) (string, error)
	PutString(int64, string) error
}

func New(svc KeyValueService) http.Handler {
	c := newController(svc)
	r := httprouter.New()

	r.GET("/:key", c.GetString)
	r.PUT("/:key/:value", c.PutString)

	return r
}
