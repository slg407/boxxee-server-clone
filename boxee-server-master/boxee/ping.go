package boxee

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

const pingResponseTemplate = `<?xml version='1.0' encoding='ISO-8859-1' ?>
<ping><cmds ping_version='9'></cmds><timestamp utc='%d'/></ping>`

// Ping verifies if network connection is still working.
// The response from the request is a timestamp.
func Ping(logger *log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		value := strings.TrimSpace(vars["pingid"])
		if strings.Compare(value, "") != 0 {
			intValue, err := strconv.Atoi(value)
			if err == nil {
				value = fmt.Sprintf("(%d) ", intValue)
			}
		}

		logger.Printf("Ping %sfrom %s", value, r.RemoteAddr)
		timestamp := time.Now().UnixNano() / int64(time.Millisecond)
		_, err := w.Write([]byte(fmt.Sprintf(pingResponseTemplate, timestamp)))
		if err != nil {
			logger.Printf("Error while responding to ping: %s", err)
		}
	}
}
