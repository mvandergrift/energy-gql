package model

import "time"

type NewMeal struct {
	ID         *int      `json:"id"`
	MealDate   time.Time `json:"mealDate"`
	MealTypeID int       `json:"mealTypeId"`
	UserID     int       `json:"userId"`
}

func (NewMeal) TableName() string {
	return "meal"
}
