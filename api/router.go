package api

import (
	"net/http"

	"github.com/crdueck/interface-demo/service"
	"github.com/julienschmidt/httprouter"
)

func New(svc service.KeyValueService) http.Handler {
	c := newController(svc)
	r := httprouter.New()

	r.GET("/:key", c.GetString)
	r.PUT("/:key/:value", c.PutString)

	return r
}
