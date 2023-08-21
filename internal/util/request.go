package util

import (
	"encoding/json"
	"goecho-boilerplate/library/derrors"
	"net/http"
	"strconv"
)

// Decode reads the body of an HTTP request looking for a JSON document. The
// body is decoded into the provided value.
func Decode(r *http.Request, val interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(val); err != nil {
		return ErrBadRequest(err, "Bad Request: "+err.Error())
	}
	return nil
}

func ParsePagination(r *http.Request) (uint64, uint64, error) {
	page := r.URL.Query().Get("page")
	if page == "" {
		page = "1"
	}
	limit := r.URL.Query().Get("limit")
	if limit == "" {
		limit = "10"
	}
	pageInt, err := strconv.ParseUint(page, 10, 64)
	if err != nil {
		return 0, 0, derrors.WrapStack(err, derrors.InvalidArgument, "limit not valid")
	}
	limitInt, err := strconv.ParseUint(limit, 10, 64)
	if err != nil {
		return 0, 0, derrors.WrapStack(err, derrors.InvalidArgument, "limit not valid")
	}
	if limitInt > 25 {
		limitInt = 25
	}
	return pageInt, limitInt, nil
}
