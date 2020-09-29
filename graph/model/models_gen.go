// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"
)

type MealType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type NewFood struct {
	Name     string   `json:"name"`
	Calories int      `json:"calories"`
	FoodImg  *string  `json:"foodImg"`
	Fat      *float64 `json:"fat"`
	SatFat   *float64 `json:"satFat"`
	Carbs    *float64 `json:"carbs"`
	Fiber    *float64 `json:"fiber"`
	Sugar    *float64 `json:"sugar"`
	Sodium   *float64 `json:"sodium"`
	Protein  *float64 `json:"protein"`
	ImgURL   *string  `json:"imgUrl"`
	UnitID   *int     `json:"unitId"`
	ID       *int     `json:"id"`
}

type NewFoodEaten struct {
	MealID int     `json:"mealId"`
	FoodID int     `json:"foodId"`
	Size   float64 `json:"size"`
	UnitID int     `json:"unitId"`
}

type NewNote struct {
	ID       *int      `json:"id"`
	NoteDate time.Time `json:"noteDate"`
	Subject  string    `json:"subject"`
	Content  string    `json:"content"`
	UserID   int       `json:"userId"`
}

type UnitType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
