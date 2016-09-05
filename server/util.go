package server

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-errors/errors"
	"strconv"
)

// DefaultLimit elements in page in paged requests
const DefaultLimit = 20

// MaxLimit max elements in page in paged requests
const MaxLimit = 200

type PageParams struct {
	Offset int `json:"offset" form:"offset"`
	Limit  int `json:"limit" form:"limit"`
}

func (pp *PageParams) fix() {
	if pp.Limit > MaxLimit {
		pp.Limit = MaxLimit
	}

	if pp.Limit == 0 {
		pp.Limit = DefaultLimit
	}
}

type PageResponse struct {
	TotalCount int `json:"totalCount"`
	PageSize   int `json:"pageSize"`
	Count      int `json:"count"`

	Content interface{} `json:"content"`
}

func WriteJSONStatus(c *gin.Context, model interface{}, httpStatus int) {
	c.JSON(httpStatus, model)
}

func parseParamsGet(c *gin.Context, params interface{}) {
	if err := c.BindWith(params, binding.Form); err != nil {
		panic(err)
	}
}

// Atoi64 parses int64 from string
func Atoi64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

// GetPathInt64Param returns path param with key
func GetPathInt64Param(c *gin.Context, key string) int64 {
	value, exists := c.Params.Get(key)
	if !exists {
		panic(errors.New("path param not exists"))
	}

	intVal, err := Atoi64(value)
	if err != nil {
		panic(err)
	}
	return intVal
}
