package cmd

import (
	"boxee-server/boxee"
	"boxee-server/proxy"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func ProxyMode(router *mux.Router, client *http.Client, logger *log.Logger) {
	logger.Println("Starting server in Proxy mode")

	router.Host("ping.boxee.tv").HandlerFunc(boxee.Ping(logger))
	router.Host("{pingid:[0-9]+}.ping.boxee.tv").HandlerFunc(boxee.Ping(logger))
	boxeeRouter := router.Host("{subdomain:[a-zA-Z]+}.boxee.tv").Subrouter()
	boxeeRouter.HandleFunc("/ping/dlink", boxee.Ping(logger))
	boxeeRouter.HandleFunc("/api/login", boxee.Login(logger))
	boxeeRouter.HandleFunc("/location/weather", boxee.FetchWeatherData(client, logger))
	boxeeRouter.PathPrefix("/").HandlerFunc(boxee.HandleBoxeeTraffic(logger))
	router.PathPrefix("/").Handler(proxy.HandleRequest(client, logger))
}

func DNSMode(router *mux.Router, client *http.Client, logger *log.Logger) {
	logger.Println("Starting server in DNS mode")

	// Handle ping.boxee.tv
	router.HandleFunc("/ping", boxee.Ping(logger))
	router.PathPrefix("/ping").HandlerFunc(boxee.Ping(logger))

	// Handle app.boxee.tv
	router.HandleFunc("/app/ping/dlink", boxee.Ping(logger))
	router.HandleFunc("/app/api/login", boxee.Login(logger))
	router.HandleFunc("/app/location/weather", boxee.FetchWeatherData(client, logger))
	router.PathPrefix("/app").HandlerFunc(boxee.HandleBoxeeTraffic(logger))

	// Handle res.boxee.tv
	router.PathPrefix("/res").HandlerFunc(boxee.HandleBoxeeTraffic(logger))

	// Handle s3.boxee.tv
	router.PathPrefix("/s3").HandlerFunc(boxee.HandleBoxeeTraffic(logger))

	// Fallback
	router.PathPrefix("/").HandlerFunc(boxee.HandleBoxeeTraffic(logger))
}
