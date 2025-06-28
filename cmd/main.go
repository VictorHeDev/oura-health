package main

import (
	"fmt"
	"log"

	"github.com/VictorHeDev/oura-health/internal/client"
	"github.com/VictorHeDev/oura-health/internal/config"
)

func main() {
	// Load configuration
	config.LoadEnv()
	ouraToken := config.GetOuraToken()
	if ouraToken == "" {
		log.Fatal("Oura API token is not set")
	}

	// Create Oura client
	ouraClient := client.NewOuraClient(ouraToken)

	// Fetch daily activity data
	dailyActivity, err := ouraClient.GetDailyActivity("2025-06-01", "2025-06-28")
	if err != nil {
		log.Fatalf("Error fetching daily activity: %v", err)
	}

	// Print daily activity data
	fmt.Printf("Daily Activity Data: %+v\n", dailyActivity)
}
