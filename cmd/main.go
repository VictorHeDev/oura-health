package main

import (
	"fmt"
	"log"
	"strings"

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
	for _, activity := range dailyActivity.Data {
		fmt.Printf("📅 Date: %s\n", activity.Day)
		fmt.Printf("🚶 Steps: %d\n", activity.Steps)
		fmt.Printf("🔥 Active Calories: %d\n", activity.ActiveCalories)
		fmt.Printf("🛌 Resting Time (min): %d\n", activity.RestingTime/60)
		fmt.Printf("⚡ Score: %d\n", activity.Score)
		fmt.Println(strings.Repeat("-", 30))
	}

}
