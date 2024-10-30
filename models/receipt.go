package models

// Item represents an item in the receipt
type Item struct {
    ShortDescription string `json:"shortDescription"`
    Price            string `json:"price"`
}

// Receipt represents the receipt data
type Receipt struct {
    Retailer     string `json:"retailer"`
    PurchaseDate string `json:"purchaseDate"`
    PurchaseTime string `json:"purchaseTime"`
    Items        []Item `json:"items"`
    Total        string `json:"total"`
    Points       int    `json:"-"`
}