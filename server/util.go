package server

import (
	"encoding/json"
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
