package model

import "time"

type Workout struct {
	ID           int       `json:"id"`
	User         *User     `json:"user"`
	ActivityDate time.Time `json:"activityDate"`
	Activity     *Activity `json:"activity"`
	Comment      *string   `json:"comment"`
	UserID       int
	ActivityID   int
	Duration     *int       `json:"duration"`
	Calories     *float64   `json:"calories"`
	Intensity    *float64   `json:"intensity"`
	Distance     *float64   `json:"distance"`
	Attribute    *int       `json:"attribute" gorm:"column:attribute_id"`
	StartTime    *time.Time `json:"startTime"`
	EndTime      *time.Time `json:"endTime"`
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
	DeletedAt    *time.Time
}

func (Workout) TableName() string {
	return "user_activity"
}
