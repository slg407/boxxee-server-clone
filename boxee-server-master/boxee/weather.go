package boxee

import (
	"boxee-server/common"
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

const weatherProviderURL = "https://wttr.in/?0&T&Q&A"
const importantLines = 2
const valueStartIndex = 16

type weatherResponse struct {
	Temperature string `json:"temperature"`
	RealFeel    string `json:"realfeel"`
	WeatherIcon string `json:"weathericon"`
}

// FetchWeatherData get the current weather from wttr.in.
// This provider takes care of getting the weather based
// on the caller's location from the IP so there's no need
// to send the location
func FetchWeatherData(client *http.Client, logger *log.Logger) http.HandlerFunc {
	var response weatherResponse

	resp, err := client.Get(weatherProviderURL)
	if err != nil {
		text := fmt.Sprintf("Error while fetching weather data: %s", err)
		return common.ReturnError(logger, text, http.StatusInternalServerError)
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	i := 0
	for scanner.Scan() && i < importantLines {
		line := scanner.Text()
		if i == importantLines-1 {
			weatherTemp := strings.Split(strings.TrimSpace(line[valueStartIndex:]), " ")
			if len(weatherTemp) == 2 {
				value := strings.TrimSpace(weatherTemp[0])
				response.Temperature = value
				response.RealFeel = value
			} else {
				text := fmt.Sprintf("Unexpected weather response format: %s", line)
				return common.ReturnError(logger, text, http.StatusInternalServerError)
			}
		}
		i++
	}
	if err != nil {
		text := fmt.Sprintf("Error while reading weather response data: %s", err)
		return common.ReturnError(logger, text, http.StatusInternalServerError)
	}

	response.WeatherIcon = "1"
	data, err := json.Marshal(&response)
	if err != nil {
		text := fmt.Sprintf("Error while marshalling weather response data: %s", err)
		return common.ReturnError(logger, text, http.StatusInternalServerError)
	}
	return func(w http.ResponseWriter, r *http.Request) {
		index := strings.Index(r.RemoteAddr, ":")
		logger.Printf("Sending weather data to %s ...", r.RemoteAddr[:index])
		w.Write(data)
	}
}
