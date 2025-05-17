package v1

import (
	"net/http"
	"os"
)

type HealthEndpoint struct {
	Ip       string `json:"ip"`
	Message  string `json:"message"`
	Hostname string `json:"hostname"`
}

// HealthHandler helps to track and check all the health params of the service
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	hostname, err := os.Hostname()
	if err != nil {
		jsonResponse(http.StatusInternalServerError, w, ErrorResponse{Error: err.Error(), ErrorMessage: "service is not healthy"})
		return
	}

	jsonResponse(http.StatusOK, w, HealthEndpoint{
		Ip:       r.Host,
		Message:  "service is healthy",
		Hostname: hostname,
	})
}
