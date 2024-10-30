package main

import (
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/gorilla/handlers"  // Import the CORS handlers library
    receiptHandlers "receipt-processor/handlers" // Alias to avoid naming conflict
)

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/receipts/process", receiptHandlers.ProcessReceipt).Methods("POST")
    r.HandleFunc("/receipts/{id}/points", receiptHandlers.GetPoints).Methods("GET")

    // Enable CORS for all origins
    corsHandler := handlers.CORS(
        handlers.AllowedOrigins([]string{"*"}),           // Allows all origins; specify if needed
        handlers.AllowedMethods([]string{"GET", "POST"}), // Allows only GET and POST methods
        handlers.AllowedHeaders([]string{"Content-Type"}), // Allows Content-Type header
    )

    log.Println("Listening on :8080")
    if err := http.ListenAndServe(":8080", corsHandler(r)); err != nil {
        log.Fatal(err)
    }
}
