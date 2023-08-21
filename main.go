package main

import (
	"goecho-boilerplate/library/config"

	echojwt "github.com/labstack/echo-jwt/v4"
)

func main() {

	echojwt.WithConfig(echojwt.Config{
		SigningMethod: "RS256",
		SigningKey:    config.Get().JWTRS256PubKey,
	})
}
