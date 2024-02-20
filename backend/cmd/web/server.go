package web

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/guilhermefill/FinTrak/pkg/utils"
	"github.com/rs/cors"
)

var (
	port   = utils.GetEnvVars("PORT")
	origin = utils.GetEnvVars("ORIGIN")
	debug  = utils.GetEnvVars("DEBUG")
)

func Server() *http.Server {
	r := mux.NewRouter()
	RoutesRegistrar(r)
	d, err := strconv.ParseBool(debug)
	if err != nil {
		log.Fatalf("Error parsing DEBUG environment variable: %v", err)
	}

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{origin},
		AllowedMethods:   []string{"GET", "HEAD", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization, Content-Type, Accept"},
		AllowCredentials: true,
		MaxAge:           86400,
		Debug:            d,
	})

	handler := c.Handler(r)
	srv := &http.Server{
		Addr:              ":" + port,
		Handler:           handler,
		ReadHeaderTimeout: 20 * time.Second,
		ReadTimeout:       1 * time.Minute,
		WriteTimeout:      2 * time.Minute,
	}

	log.Println("Server started on port: ", srv.Addr)
	return srv
}
