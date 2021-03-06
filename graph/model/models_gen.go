// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type Activity struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	ImgURL *string `json:"imgURL"`
}

type MealType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type NewActivity struct {
	ID     *int    `json:"id"`
	Name   string  `json:"name"`
	UserID int     `json:"userId"`
	ImgURL *string `json:"imgURL"`
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

type NewWorkout struct {
	ID           *int       `json:"id"`
	UserID       int        `json:"userId"`
	ActivityDate time.Time  `json:"activityDate"`
	ActivityID   int        `json:"activityId"`
	Duration     *int       `json:"duration"`
	Calories     *float64   `json:"calories"`
	Intensity    *float64   `json:"intensity"`
	Distance     *float64   `json:"distance"`
	StartTime    *time.Time `json:"startTime"`
	EndTime      *time.Time `json:"endTime"`
	Comment      *string    `json:"comment"`
}

type Predicate struct {
	Name           *string         `json:"name"`
	Values         []string        `json:"values"`
	Operator       *Operator       `json:"operator"`
	InnerPredicate *PredicateGroup `json:"innerPredicate"`
}

type PredicateGroup struct {
	Predicates []*Predicate `json:"predicates"`
	Logic      Logic        `json:"logic"`
}

type PredicateResult struct {
	Query *string `json:"query"`
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

type Logic string

const (
	LogicAnd Logic = "AND"
	LogicOr  Logic = "OR"
)

var AllLogic = []Logic{
	LogicAnd,
	LogicOr,
}

func (e Logic) IsValid() bool {
	switch e {
	case LogicAnd, LogicOr:
		return true
	}
	return false
}

func (e Logic) String() string {
	return string(e)
}

func (e *Logic) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Logic(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Logic", str)
	}
	return nil
}

func (e Logic) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Operator string

const (
	OperatorEqual        Operator = "EQUAL"
	OperatorNotEqual     Operator = "NOT_EQUAL"
	OperatorGreator      Operator = "GREATOR"
	OperatorLess         Operator = "LESS"
	OperatorGreatorEqual Operator = "GREATOR_EQUAL"
	OperatorLessEqual    Operator = "LESS_EQUAL"
	OperatorContains     Operator = "CONTAINS"
	OperatorStartsWith   Operator = "STARTS_WITH"
	OperatorIsNull       Operator = "IS_NULL"
	OperatorIsNotNull    Operator = "IS_NOT_NULL"
)

var AllOperator = []Operator{
	OperatorEqual,
	OperatorNotEqual,
	OperatorGreator,
	OperatorLess,
	OperatorGreatorEqual,
	OperatorLessEqual,
	OperatorContains,
	OperatorStartsWith,
	OperatorIsNull,
	OperatorIsNotNull,
}

func (e Operator) IsValid() bool {
	switch e {
	case OperatorEqual, OperatorNotEqual, OperatorGreator, OperatorLess, OperatorGreatorEqual, OperatorLessEqual, OperatorContains, OperatorStartsWith, OperatorIsNull, OperatorIsNotNull:
		return true
	}
	return false
}

func (e Operator) String() string {
	return string(e)
}

func (e *Operator) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Operator(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Operator", str)
	}
	return nil
}

func (e Operator) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
