package main

import (
	"boxee-server/cmd"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

const defaultPort = 7777

var usageText = fmt.Sprintf(`
Usage:	boxee-server [OPTIONS] COMMAND

A replacement Boxee server

Options:
  -p, --port               Port used by Boxee server (default %d)

Commands:
  dns         Start server in DNS mode
  proxy       Start server in Proxy mode`, defaultPort)

func printUsage(errorText string) {
	fmt.Printf("%s\n%s", errorText, usageText)
	os.Exit(2)
}
func main() {

	logger := log.New(os.Stdout, "[BOXEE SERVER] ", log.Ldate|log.Ltime)
	portFlag := flag.Int("port", defaultPort, "Port used by Boxee server")
	flag.Parse()

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	router := mux.NewRouter()
	host := "0.0.0.0"
	port := defaultPort
	portValue, ok := os.LookupEnv("BOXEE_SERVER_PORT")
	if ok {
		portNumber, err := strconv.Atoi(portValue)
		if err == nil {
			logger.Println("Setting custom port from environment")
			port = portNumber
		}
	} else {
		if *portFlag != defaultPort {
			port = *portFlag
		}
	}

	args := flag.Args()
	if len(args) > 1 {
		printUsage("Unexpected number of arguments passed")
	}

	if len(args) == 0 {
		cmd.ProxyMode(router, client, logger)
	} else {
		command := args[0]
		switch command {
		case "dns":
			cmd.DNSMode(router, client, logger)
		case "proxy":
			cmd.ProxyMode(router, client, logger)
		default:
			printUsage("Unknown command passed")
		}
	}

	addr := fmt.Sprintf("%s:%d", host, port)
	srv := &http.Server{
		Addr:         addr,
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		ErrorLog:     logger,
	}

	logger.Printf("Boxee Proxy Server listening on %s", addr)
	if err := srv.ListenAndServe(); err != nil {
		logger.Fatalln(err)
	}
}
