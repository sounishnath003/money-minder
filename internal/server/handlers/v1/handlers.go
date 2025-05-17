package v1

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/sounishnath003/money-minder/internal/models"
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

// GetTransactionsHandler handler helps to get the total balance available for this month
func GetTransactionsHandler(w http.ResponseWriter, r *http.Request) {
	fromDate := r.URL.Query().Get("from")
	toDate := r.URL.Query().Get("to")

	fromDateParsed, toDateParsed, err := CheckDateRange(fromDate, toDate)
	if err != nil {
		jsonResponse(http.StatusBadRequest, w, ErrorResponse{Error: err.Error(), ErrorMessage: "date range error"})
		return
	}

	log.Printf("Fetching transactions from %s to %s", fromDateParsed.Format("2006-01-02"), toDateParsed.Format("2006-01-02"))

	jsonResponse(http.StatusOK, w, []models.Transaction{
		{ID: 1, Name: "Salary", CatagoryID: "cat-001", UserID: 1, Timestamp: time.Now()},
	})
}
