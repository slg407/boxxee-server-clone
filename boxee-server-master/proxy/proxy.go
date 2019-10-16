package proxy

import (
	"io"
	"log"
	"net/http"
)

// HandleRequest receives the incoming requests and performs
// it on behalf of the caller without any modifications
func HandleRequest(client *http.Client, logger *log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Printf("Proxying: %s %s%s", r.Method, r.RequestURI, r.URL.Path)
		req, err := http.NewRequest(r.Method, r.RequestURI, r.Body)
		for name, value := range r.Header {
			req.Header.Set(name, value[0])
		}
		resp, err := client.Do(req)
		r.Body.Close()

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			logger.Printf("Error on proxy request: %s", err)
			return
		}

		for k, v := range resp.Header {
			w.Header().Set(k, v[0])
		}
		w.WriteHeader(resp.StatusCode)
		io.Copy(w, resp.Body)
		resp.Body.Close()
	}
}
