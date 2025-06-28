package models

type DailyActivityContributors struct {
	MeetDailyTargets  int `json:"meet_daily_targets"`
	MoveEveryHour     int `json:"move_every_hour"`
	RecoveryTime      int `json:"recovery_time"`
	StayActive        int `json:"stay_active"`
	TrainingFrequency int `json:"training_frequency"`
	TrainingVolume    int `json:"training_volume"`
}

type DailyActivityMet struct {
	Interval  float64   `json:"interval"`
	Items     []float64 `json:"items"`
	Timestamp string    `json:"timestamp"`
}

type DailyActivity struct {
	Id                string  `json:"id"`
	Class5Min         string  `json:"class_5_min"`
	Score             int     `json:"score"`
	ActiveCalories    int     `json:"active_calories"`
	AverageMetMinutes float64 `json:"average_met_minutes"`
	// Contributors              DailyActivityContributors `json:"contributors"`
	EquivalentWalkingDistance int     `json:"equivalent_walking_distance"`
	HighActivityMetMinutes    float64 `json:"high_activity_met_minutes"`
	HighActivityTime          float64 `json:"high_activity_time"`
	InactivityAlerts          int     `json:"inactivity_alerts"`
	LowActivityMetMinutes     int     `json:"low_activity_met_minutes"`
	LowActivityTime           int     `json:"low_activity_time"`
	MediumActivityMetMinutes  int     `json:"medium_activity_met_minutes"`
	MediumActivityTime        int     `json:"medium_activity_time"`
	// Met                       DailyActivityMet          `json:"met"`
	MetersToTarget      int    `json:"meters_to_target"`
	NonWearTime         int    `json:"non_wear_time"`
	RestingTime         int    `json:"resting_time"`
	SedentaryMetMinutes int    `json:"sedentary_met_minutes"`
	SedentaryTime       int    `json:"sedentary_time"`
	Steps               int    `json:"steps"`
	TargetCalories      int    `json:"target_calories"`
	TargetMeters        int    `json:"target_meters"`
	TotalCalories       int    `json:"total_calories"`
	Day                 string `json:"day"`
	Timestamp           string `json:"timestamp"`
}
