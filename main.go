package main

import (
	"fmt"
	"net/http"

	"github.com/auth0/go-jwt-middleware"
	"github.com/caarlos0/env"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"

	"github.com/dgrijalva/jwt-go"
	"github.com/nuveo/nuance/config"
	"github.com/nuveo/nuance/controllers"
)

func main() {
	cfg := config.Nuance{}
	env.Parse(&cfg)

	n := negroni.Classic()
	n.Use(negroni.HandlerFunc(handlerSet))
	if cfg.JWTKey != "" {
		n.Use(jwtMiddleware(cfg.JWTKey))
	}

	fmt.Println(cfg.OemLicenseFile, cfg.OemCode )

	r := mux.NewRouter()
	r.HandleFunc("/omnipage", controllers.PostImg).Methods("POST")

	n.UseHandler(r)
	n.Run(fmt.Sprintf(":%v", cfg.HTTPPort))
}

func handlerSet(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	w.Header().Set("Content-Type", "application/json")
	next(w, r)
}

func jwtMiddleware(key string) negroni.Handler {
	jwtm := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(key), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})
	return negroni.HandlerFunc(jwtm.HandlerWithNext)
}
