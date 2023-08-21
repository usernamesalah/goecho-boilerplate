package dto

import (
	"github.com/dgrijalva/jwt-go"
)

type ErrorData struct {
	Code    int    `json:"code"`
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type Response struct {
	Success bool        `json:"success"`
	Error   *ErrorData  `json:"errors"`
	Data    interface{} `json:"data"`
	Paging  *PageData   `json:"paging"`
}

type PageData struct {
	HasNext     bool  `json:"has_next"`
	CurrentPage int   `json:"current_page"`
	TotalPages  int   `json:"total_pages"`
	TotalData   int64 `json:"total_data"`
	Limit       int   `json:"limit"`
}

type ListQueryParams struct {
	Limit  int
	Offset int
}

type JWTClaims struct {
	UID  string `json:"uid"`
	Name string `json:"name"`
	jwt.StandardClaims
}
