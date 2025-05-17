package v1

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

// ErrorResponse sends the response whenever the error occured in the service
type ErrorResponse struct {
	Error        string `json:"error"`
	ErrorMessage string `json:"errorMessage"`
}

// SuccessResponse central success response while
// sending a SUCESS JSON response back to client
type SuccessResponse struct {
	Data interface{} `json:"data"`
}

// FaliureResponse central failure response while
// sending a FAILURE JSON response back to client
type FaliureResponse struct {
	Error interface{} `json:"error"`
}

// jsonResponse sends the response as response to the clients
// use the method to efficiently send any success or failure responses to the clients
func jsonResponse(statusCode int, w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Request-ID", uuid.New().String())
	w.WriteHeader(statusCode)

	// All sucess statuc codes
	if statusCode >= 200 && statusCode < 400 {
		succ := SuccessResponse{Data: data}
		if err := json.NewEncoder(w).Encode(succ); err != nil {
			http.Error(w, "Error encoding response", http.StatusInternalServerError)
			return
		}

	} else {
		fail := FaliureResponse{Error: data}
		if err := json.NewEncoder(w).Encode(fail); err != nil {
			http.Error(w, "Error encoding response", http.StatusInternalServerError)
			return
		}
	}
}
