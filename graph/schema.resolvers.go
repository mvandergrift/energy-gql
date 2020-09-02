package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	"github.com/mvandergrift/energy-gql/graph/generated"
	"github.com/mvandergrift/energy-gql/graph/model"
)

func (r *mealResolver) Food(ctx context.Context, obj *model.Meal) ([]*model.MealFood, error) {
	var food []*model.MealFood

	if err := r.DB.Table("meal_food").Select("meal_food.id, meal_food.size, food.calories, food.name, food.img_url as food_img").Joins("join meal ON meal.id = meal_food.meal_id").Joins("join food on food.id = meal_food.food_id").Where("meal.id = ?", obj.ID).Scan(&food).Error; err != nil {
		return food, err
	}

	return food, nil
}

func (r *queryResolver) AllMeals(ctx context.Context, userID *int) ([]*model.Meal, error) {
	var meals []*model.Meal
	var tx = r.DB.Table("meal").Select("meal.id, meal.meal_date, meal_type.name as meal_type").Joins("join meal_type ON meal_type.id = meal.meal_type_id").Order("meal_date, meal_type_id")

	if userID != nil {
		tx.Where("user_id = ?", userID)
	}

	if err := tx.Scan(&meals).Error; err != nil {
		return meals, err
	}

	return meals, nil
}

func (r *queryResolver) MealsForDay(ctx context.Context, userID int, date time.Time) ([]*model.Meal, error) {
	var meals []*model.Meal
	var tx = r.DB.Table("meal").Select("meal.id, meal.meal_date, meal_type.name as meal_type").Joins("join meal_type ON meal_type.id = meal.meal_type_id").Where("user_id = ? and meal.meal_date = ?", userID, date).Order("meal_date, meal_type_id")

	if err := tx.Scan(&meals).Error; err != nil {
		return nil, err
	}

	return meals, nil
}

func (r *queryResolver) AllFoods(ctx context.Context, userID *int) ([]*model.Food, error) {
	var foods []*model.Food
	var err = r.DB.Find(&foods).Error
	return foods, err
}

// Meal returns generated.MealResolver implementation.
func (r *Resolver) Meal() generated.MealResolver { return &mealResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mealResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
