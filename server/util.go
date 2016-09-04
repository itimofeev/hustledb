package server

import (
	"encoding/json"
	"github.com/gorilla/schema"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

// DefaultLimit elements in page in paged requests
const DefaultLimit = 20

// MaxLimit max elements in page in paged requests
const MaxLimit = 200

type PageParams struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type PageResponse struct {
	TotalCount int `json:"totalCount"`
	PageSize   int `json:"pageSize"`
	Count      int `json:"count"`

	Content interface{} `json:"content"`
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
	parseParamsJSONBody(w, r, params)

	if params.Limit > MaxLimit {
		params.Limit = MaxLimit
	}

	if params.Limit == 0 {
		params.Limit = DefaultLimit
	}
}

func parseParamsJSONBody(w http.ResponseWriter, r *http.Request, params interface{}) {
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

func parseParamsGet(w http.ResponseWriter, r *http.Request, params interface{}) {
	decoder := schema.NewDecoder()
	r.ParseForm()

	if err := decoder.Decode(params, r.Form); err != nil {
		panic(err)
	}

	if pp, ok := params.(*PageParams); ok {
		if pp.Limit > MaxLimit {
			pp.Limit = MaxLimit
		}

		if pp.Limit == 0 {
			pp.Limit = DefaultLimit
		}
	}
}

// Atoi64 parses int64 from string
func Atoi64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}
