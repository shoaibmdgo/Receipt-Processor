package utils

import (
    "math"
    "regexp"
    "strconv"
    "strings"
    "time"
    "github.com/google/uuid"
    "receipt-processor/models"
)

// GenerateID generates a unique ID for a receipt
func GenerateID() string {
    return uuid.New().String()
}

// CalculatePoints calculates points based on rules
func CalculatePoints(receipt models.Receipt) int {
    points := 0

    // Rule 1: One point for every alphanumeric character in the retailer name
    alphanumeric := regexp.MustCompile(`[a-zA-Z0-9]`)
    points += len(alphanumeric.FindAllString(receipt.Retailer, -1))

    // Rule 2: 50 points if the total is a round dollar amount with no cents
    total, _ := strconv.ParseFloat(receipt.Total, 64)
    if total == float64(int(total)) {
        points += 50
    }

    // Rule 3: 25 points if the total is a multiple of 0.25
    if math.Mod(total, 0.25) == 0 {
        points += 25
    }

    // Rule 4: 5 points for every two items
    points += (len(receipt.Items) / 2) * 5

    // Rule 5: Check item description length multiple of 3
    for _, item := range receipt.Items {
        descLen := len(strings.TrimSpace(item.ShortDescription))
        if descLen%3 == 0 {
            itemPrice, _ := strconv.ParseFloat(item.Price, 64)
            points += int(math.Ceil(itemPrice * 0.2))
        }
    }

    // Rule 6: 6 points if the day in the purchase date is odd
    date, _ := time.Parse("2006-01-02", receipt.PurchaseDate)
    if date.Day()%2 != 0 {
        points += 6
    }

    // Rule 7: 10 points if the time of purchase is between 2:00pm and 4:00pm
    purchaseTime, _ := time.Parse("15:04", receipt.PurchaseTime)
    if purchaseTime.Hour() == 14 || (purchaseTime.Hour() == 15 && purchaseTime.Minute() < 60) {
        points += 10
    }

    return points
}