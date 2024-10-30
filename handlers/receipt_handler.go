package handlers

import (
    "encoding/json"
    "net/http"
    "receipt-processor/models"
    "receipt-processor/utils"
    "sync"
)

var (
    receiptStore = make(map[string]models.Receipt) // In-memory store
    mu           sync.Mutex                        // Mutex for concurrent access
)

// ProcessReceipt handles the /receipts/process endpoint
func ProcessReceipt(w http.ResponseWriter, r *http.Request) {
    var receipt models.Receipt
    err := json.NewDecoder(r.Body).Decode(&receipt)
    if err != nil {
        http.Error(w, "Invalid JSON format", http.StatusBadRequest)
        return
    }

    // Calculate points and generate ID
    points := utils.CalculatePoints(receipt)
    receiptID := utils.GenerateID()
    receipt.Points = points

    // Store the receipt in memory
    mu.Lock()
    receiptStore[receiptID] = receipt
    mu.Unlock()

    // Respond with the generated ID
    w.WriteHeader(http.StatusOK) // Set response status code to 200
    json.NewEncoder(w).Encode(map[string]string{"id": receiptID})
}

// GetPoints handles the /receipts/{id}/points endpoint
func GetPoints(w http.ResponseWriter, r *http.Request) {
    // Extracting the ID from the URL
    id := r.URL.Path[len("/receipts/") : len(r.URL.Path)-len("/points")]

    mu.Lock()
    receipt, exists := receiptStore[id]
    mu.Unlock()

    if !exists {
        http.Error(w, "Receipt ID not found", http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusOK) // Set response status code to 200
    json.NewEncoder(w).Encode(map[string]int{"points": receipt.Points})
}
