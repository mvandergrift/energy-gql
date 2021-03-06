package model

import "time"

type FoodEaten struct {
	ID        int      `json:"id"`
	Name      string   `json:"name"`
	Size      *float64 `json:"size"`
	Calories  *int     `json:"calories"`
	Food      *Food    `json:"food"`
	FoodID    int      `gorm:"primary_key"`
	MealID    int      `gorm:"primary_key"`
	UnitID    int
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}

func (FoodEaten) TableName() string {
	return "meal_food"
}
