package server

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

// DefaultLimit elements in page in paged requests
const DefaultLimit = 20

// MaxLimit max elements in page in paged requests
const MaxLimit = 200

type PageParams struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

func WriteJSON(w http.ResponseWriter, model interface{}) {
	WriteJSONStatus(w, model, http.StatusOK)
}

func WriteJSONStatus(w http.ResponseWriter, model interface{}, httpStatus int) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(httpStatus)

	if err := json.NewEncoder(w).Encode(model); err != nil {
		panic(err)
	}
}

func parsePageParams(w http.ResponseWriter, r *http.Request, params *PageParams) {
	parseParams(w, r, params)

	if params.Limit > MaxLimit {
		params.Limit = MaxLimit
	}

	if params.Limit == 0 {
		params.Limit = DefaultLimit
	}
}

func parseParams(w http.ResponseWriter, r *http.Request, params interface{}) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if len(body) == 0 {
		return
	}

	if err := json.Unmarshal(body, &params); err != nil {
		panic(err)
	}
}
