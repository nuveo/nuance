package main

import (
	"fmt"
	"net/http"

	"github.com/auth0/go-jwt-middleware"
	"github.com/caarlos0/env"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"

	"github.com/dgrijalva/jwt-go"
	"github.com/nuveo/nuance/config"
	"github.com/nuveo/nuance/controllers"
	"github.com/nuveo/nuance/omnipage"
)

func main() {
	cfg := config.Nuance{}
	env.Parse(&cfg)

	n := negroni.Classic()
	n.Use(negroni.HandlerFunc(handlerSet))
	if cfg.JWTKey != "" {
		n.Use(jwtMiddleware(cfg.JWTKey))
	}

	op := omnipage.New()

	err := op.SetLicense(cfg.OemLicenseFile, cfg.OemCode)
	if err != nil {
		fmt.Println("SetLicense failed:", err)
		return
	}

	err = op.Init(cfg.CompanyName, cfg.ProductName)
	if err != nil {
		fmt.Println("Init failed:", err)
		return
	}

	controllers.SetOmnipage(&op)
	controllers.SetConfig(&cfg)

	r := mux.NewRouter()
	r.HandleFunc("/omnipage/totext", controllers.ImgToText).Methods("POST")
	r.HandleFunc("/omnipage/ocrwithtemplate", controllers.ImgWithTemplate).Methods("POST")

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
