package main

import (
    "fmt"
    "time"
)

func main() {
    // Get current date
    currentDate := time.Now()

    // Add one year and two days
    oneYearTwoDaysLater := currentDate.AddDate(1, 0, 2)

    // Format the date in the desired format
    formattedDate := oneYearTwoDaysLater.Format("2006-01-02")

    // Print the result
    fmt.Println("One year and two days later:", formattedDate)
}
