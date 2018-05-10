package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/crdueck/interface-demo/service"
	"github.com/crdueck/interface-demo/storage"
	"github.com/julienschmidt/httprouter"
)

func parseInt64Param(ps httprouter.Params, name string) (int64, error) {
	return strconv.ParseInt(ps.ByName(name), 10, 64)
}

type Controller struct {
	service service.KeyValueService
}

func newController(svc service.KeyValueService) *Controller {
	return &Controller{
		service: svc,
	}
}

func (c *Controller) GetString(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	key, err := parseInt64Param(ps, "key")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	val, err := c.service.GetString(key)
	if err != nil {
		if err == storage.ErrNotFound {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	type response struct {
		Value string `json:"value"`
	}

	json.NewEncoder(w).Encode(response{val})
}

func (c *Controller) PutString(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	key, err := parseInt64Param(ps, "key")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	val := ps.ByName("value")
	if val == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	c.service.PutString(key, val)
}
