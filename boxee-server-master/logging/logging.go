package logging

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

// PrintHTTP outputs the HTTP request-response to the specified logger
// and uses StringPrinter to make the string representation.
func PrintHTTP(logger *log.Logger, req *http.Request, resp *http.Response) {
	logger.Println(StringPrinter(req, resp))
}

// StringPrinter creates a string representation of a request-response
// done using HTTP, it mainly used for debugging.
func StringPrinter(req *http.Request, resp *http.Response) string {
	var builder strings.Builder
	builder.WriteString("==============================\n")
	builder.WriteString(fmt.Sprintf("Remote Address: %s\n", req.RemoteAddr))
	builder.WriteString(fmt.Sprintf("%v %v\n", req.Method, req.RequestURI))
	for k, v := range req.Header {
		builder.WriteString(fmt.Sprintln(k, ":", v))
	}
	builder.WriteString("------------------------------\n")
	if resp != nil {
		builder.WriteString(fmt.Sprintf("HTTP/1.1 %v\n", resp.Status))
		for k, v := range resp.Header {
			builder.WriteString(fmt.Sprintln(k, ":", v))
		}
		fmt.Println(resp.Body)
	} else {
		builder.WriteString("No response...\n")
	}
	builder.WriteString("==============================\n")
	return builder.String()
}
