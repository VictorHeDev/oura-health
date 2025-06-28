package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/VictorHeDev/oura-health/internal/models"
)

type OuraClient struct {
	Token string
}

type DailyActivityResponse struct {
	Data []models.DailyActivity `json:"data"`
}

func NewOuraClient(ouraToken string) *OuraClient {
	return &OuraClient{
		Token: ouraToken,
	}
}

func (c *OuraClient) GetDailyActivity(startDate, endDate string) (*DailyActivityResponse, error) {
	url := fmt.Sprintf("https://api.ouraring.com/v2/usercollection/daily_activity?start_date=%s&end_date=%s", startDate, endDate)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+c.Token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch daily activity: %s", resp.Status)
	}

	var dailyActivityResponse DailyActivityResponse
	if err := json.NewDecoder(resp.Body).Decode(&dailyActivityResponse); err != nil {
		return nil, err
	}

	return &dailyActivityResponse, nil
}
