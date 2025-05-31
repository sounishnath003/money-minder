package v1

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/sounishnath003/money-minder/internal/core"
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

	co := r.Context().Value("co").(*core.Core)
	co.Logger.Printf("Health check requested from %s", r.Host)

	jsonResponse(http.StatusOK, w, HealthEndpoint{
		Ip:       r.Host,
		Message:  "service is healthy",
		Hostname: hostname,
	})
}

func GetAllCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	co := r.Context().Value("co").(*core.Core)
	co.Logger.Printf("Fetching all active categories")

	categories, err := co.BQClient.GetActiveCategories()
	if err != nil {
		co.Logger.Printf("Error fetching categories: %v", err)
		jsonResponse(http.StatusInternalServerError, w, ErrorResponse{Error: err.Error(), ErrorMessage: "unable to fetch categories"})
		return
	}

	co.Logger.Printf("Successfully fetched %d categories", len(categories))
	jsonResponse(http.StatusOK, w, categories)
}

func GetAllPaymentMethodOptionsHandler(w http.ResponseWriter, r *http.Request) {
	co := r.Context().Value("co").(*core.Core)
	co.Logger.Printf("Fetching all active categories")

	paymentMethods := []models.PaymentMethod{
		models.UPI,
		models.Cash,
		models.BankDeposit,
		models.CreditCard,
	}

	co.Logger.Printf("sending all available payment methods: %d", len(paymentMethods))
	jsonResponse(http.StatusOK, w, paymentMethods)
}

// GetTransactionsHandler handler helps to get the total balance available for this month
func GetTransactionsHandler(w http.ResponseWriter, r *http.Request) {
	fromDate := r.URL.Query().Get("from")
	toDate := r.URL.Query().Get("to")
	userID := 1

	fromDateParsed, toDateParsed, err := CheckDateRange(fromDate, toDate)
	if err != nil {
		co := r.Context().Value("co").(*core.Core)
		co.Logger.Printf("Invalid date range: %v", err)
		jsonResponse(http.StatusBadRequest, w, ErrorResponse{Error: err.Error(), ErrorMessage: "date range error"})
		return
	}

	co := r.Context().Value("co").(*core.Core)
	co.Logger.Printf("Fetching transactions from %s to %s", fromDateParsed.Format("2006-01-02 "), toDateParsed.Format("2006-01-02"))

	out, err := co.BQClient.GetTransactionsByUserId(userID, fromDateParsed, toDateParsed)
	if err != nil {
		co.Logger.Printf("Error fetching transactions: %v", err)
		jsonResponse(http.StatusInternalServerError, w, ErrorResponse{Error: err.Error(), ErrorMessage: "unable to fetch transactions"})
		return
	}

	co.Logger.Printf("Successfully fetched %d transactions for userID: %d", len(out), userID)
	jsonResponse(http.StatusOK, w, out)
}

// AddTransactionHandler helps to add new transaction with specified DTO
func AddTransactionHandler(w http.ResponseWriter, r *http.Request) {
	var transaction models.Transaction
	err := json.NewDecoder(r.Body).Decode(&transaction)

	if err != nil {
		co := r.Context().Value("co").(*core.Core)
		co.Logger.Printf("Error parsing transaction request: %v", err)
		jsonResponse(http.StatusBadRequest, w, ErrorResponse{Error: err.Error(), ErrorMessage: "unable to parse request body"})
		return
	}

	co := r.Context().Value("co").(*core.Core)
	co.Logger.Printf("Adding new transaction: %+v", transaction)

	success, err := co.BQClient.AddTransaction(transaction)
	if err != nil {
		co.Logger.Printf("Error creating transaction: %v", err)
		jsonResponse(http.StatusBadRequest, w, ErrorResponse{Error: err.Error(), ErrorMessage: "unable to create transaction"})
		return
	}

	co.Logger.Printf("Successfully added transaction with ID: %d", transaction.ID)
	jsonResponse(http.StatusOK, w, success)
}

// SpendByCategoryHandler helps to get the spend by category of a User.
// Group of aggregation query will be performed
func GetTotalSpendByCategoryHandler(w http.ResponseWriter, r *http.Request) {
	fromDate := r.URL.Query().Get("from")
	toDate := r.URL.Query().Get("to")
	userID := 1

	fromDateParsed, toDateParsed, err := CheckDateRange(fromDate, toDate)
	if err != nil {
		co := r.Context().Value("co").(*core.Core)
		co.Logger.Printf("Invalid date range: %v", err)
		jsonResponse(http.StatusBadRequest, w, ErrorResponse{Error: err.Error(), ErrorMessage: "date range error"})
		return
	}

	co := r.Context().Value("co").(*core.Core)
	co.Logger.Printf("Fetching transactions from %s to %s", fromDateParsed.Format("2006-01-02"), toDateParsed.Format("2006-01-02"))

	out, err := co.BQClient.GetTotalSpendsByCategories(userID, fromDateParsed, toDateParsed)
	if err != nil {
		co.Logger.Printf("Error fetching transactions: %v", err)
		jsonResponse(http.StatusInternalServerError, w, ErrorResponse{Error: err.Error(), ErrorMessage: "unable to fetch transactions"})
		return
	}

	co.Logger.Printf("Successfully fetched %d transactions for userID: %d", len(out), userID)
	jsonResponse(http.StatusOK, w, out)
}

// GetDailyTotalSpendByTimeframeHandler helps to find the analytical handler which
// helps to get the daily total spend (incl. any category)
func GetDailyTotalSpendByTimeframeHandler(w http.ResponseWriter, r *http.Request) {
	fromDate := r.URL.Query().Get("from")
	toDate := r.URL.Query().Get("to")
	userID := 1

	fromDateParsed, toDateParsed, err := CheckDateRange(fromDate, toDate)
	if err != nil {
		co := r.Context().Value("co").(*core.Core)
		co.Logger.Printf("Invalid date range: %v", err)
		jsonResponse(http.StatusBadRequest, w, ErrorResponse{Error: err.Error(), ErrorMessage: "date range error"})
		return
	}

	co := r.Context().Value("co").(*core.Core)
	co.Logger.Printf("Fetching transactions from %s to %s", fromDateParsed.Format("2006-01-02"), toDateParsed.Format("2006-01-02"))

	out, err := co.BQClient.GetDailyTotalSpendByTimeframe(userID, fromDateParsed, toDateParsed)
	if err != nil {
		co.Logger.Printf("Error fetching transactions: %v", err)
		jsonResponse(http.StatusInternalServerError, w, ErrorResponse{Error: err.Error(), ErrorMessage: "unable to fetch transactions"})
		return
	}

	co.Logger.Printf("Successfully fetched %d daily total spend analytical for userID: %d", len(out), userID)
	jsonResponse(http.StatusOK, w, out)
}

func GetSpendOnCategoriesMonthOnMonthHandler(w http.ResponseWriter, r *http.Request) {
	fromDate := r.URL.Query().Get("from")
	toDate := r.URL.Query().Get("to")
	userID := 1

	fromDateParsed, toDateParsed, err := CheckDateRange(fromDate, toDate)
	if err != nil {
		co := r.Context().Value("co").(*core.Core)
		co.Logger.Printf("Invalid date range: %v", err)
		jsonResponse(http.StatusBadRequest, w, ErrorResponse{Error: err.Error(), ErrorMessage: "date range error"})
		return
	}

	co := r.Context().Value("co").(*core.Core)
	co.Logger.Printf("Fetching transactions from %s to %s", fromDateParsed.Format("2006-01-02"), toDateParsed.Format("2006-01-02"))

	out, err := co.BQClient.GetSpendOnCategoriesMonthOnMonth(userID, fromDateParsed, toDateParsed)
	if err != nil {
		co.Logger.Printf("Error fetching transactions: %v", err)
		jsonResponse(http.StatusInternalServerError, w, ErrorResponse{Error: err.Error(), ErrorMessage: "unable to fetch transactions"})
		return
	}

	co.Logger.Printf("Successfully fetched %d spend month on month analytical for userID: %d", len(out), userID)
	jsonResponse(http.StatusOK, w, out)

}
