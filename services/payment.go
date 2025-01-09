package services

import (
    "fmt"
    "time"
)

func ProcessPayment(amount float64, done chan<- bool) {
    go func() {
        defer close(done)
        // Simulate payment processing time
        time.Sleep(3 * time.Second)

        // Process payment logic (e.g., interacting with Stripe)
        // For simplicity, we just print a message here
        fmt.Printf("Processing payment of $%.2f\n", amount)

        // Notify that payment is done
        done <- true
    }()
}
