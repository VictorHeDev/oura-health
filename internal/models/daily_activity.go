package models

type DailyActivity struct {
    Summary DailyActivitySummary `json:"summary"`
    Data    []DailyActivityData  `json:"data"`
}

type DailyActivitySummary struct {
    Steps         int `json:"steps"`
    ActiveTime    int `json:"active_time"`
    RestingTime   int `json:"resting_time"`
    TotalCalories int `json:"total_calories"`
}

type DailyActivityData struct {
    Date          string `json:"date"`
    Steps         int    `json:"steps"`
    ActiveTime    int    `json:"active_time"`
    RestingTime   int    `json:"resting_time"`
    TotalCalories int    `json:"total_calories"`
}