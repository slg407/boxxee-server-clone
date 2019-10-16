package boxee

import (
	"boxee-server/logging"
	"fmt"
	"log"
	"net/http"
)

// HandleBoxeeTraffic is the default way of handling traffic going to
// *.boxee.tv URLs. This simply dumps the callers requests into the response
// making sure that the request receives a 200 OK.
//
// More specific handlers are create to return the appropriate data.
// Note: This is mainly used for debugging.
func HandleBoxeeTraffic(logger *log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		loggedContent := logging.StringPrinter(r, nil)
		logger.Println(loggedContent)
		_, err := fmt.Fprintf(w, "Handling boxee.tv request:\n%s", loggedContent)
		if err != nil {
			logger.Println(err)
		}
	}
}
