package model

import (
	"time"
)

type Meal struct {
	ID         int          `json:"id"`
	MealDate   time.Time    `json:"mealDate"`
	MealTypeID int          `json:"mealTypeId"`
	MealType   *MealType    `json:"mealType"`
	User       *User        `json:"user"`
	FoodEaten  []*FoodEaten `json:"foodEaten"`
}
