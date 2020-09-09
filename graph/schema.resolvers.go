package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"
	"time"

	"github.com/mvandergrift/energy-gql/graph/generated"
	"github.com/mvandergrift/energy-gql/graph/model"
)

func (r *mutationResolver) DeleteFood(ctx context.Context, id int) (*model.Food, error) {
	var food model.Food
	r.DB.Where("id = ?", id).First(&food)
	err := r.DB.Delete(&food).Error
	return &food, err
}

func (r *mutationResolver) AddFood(ctx context.Context, food model.NewFood) (*model.Food, error) {
	newFood := model.Food{Name: food.Name, Calories: food.Calories, ImgURL: food.FoodImg, UnitID: *food.UnitID}
	result := r.DB.Omit("ID").Create(&newFood)
	log.Println(&newFood)
	r.DB.Preload("Unit").Preload("Unit.UnitType").First(&newFood, result.Value)

	return &newFood, result.Error
}

func (r *queryResolver) AllMeals(ctx context.Context, userID *int) ([]*model.Meal, error) {
	var meals []*model.Meal
	tx := r.DB.Order("meal_date, meal_type_id")

	if userID != nil {
		tx.Where("user_id = ?", userID)
	}

	err := tx.
		Preload("MealType").
		Preload("FoodEaten").
		Preload("FoodEaten.Food").
		Preload("FoodEaten.Food.Unit").
		Preload("FoodEaten.Food.Unit.UnitType").
		Find(&meals).Error

	return meals, err
}

func (r *queryResolver) MealsForDay(ctx context.Context, userID int, date time.Time) ([]*model.Meal, error) {
	var meals []*model.Meal
	err := r.DB.
		Preload("MealType").
		Preload("FoodEaten").
		Preload("FoodEaten.Food").
		Preload("FoodEaten.Food.Unit").
		Preload("FoodEaten.Food.Unit.UnitType").
		Order("meal_date, meal_type_id").
		Where("user_id = ? and meal_date = ? ", userID, date).
		Find(&meals).Error

	return meals, err
}

func (r *queryResolver) AllFoods(ctx context.Context, userID *int) ([]*model.Food, error) {
	var foods []*model.Food
	err := r.DB.Preload("Unit").Preload("Unit.UnitType").Order("unit_type_id, name").Find(&foods).Error
	return foods, err
}

func (r *queryResolver) AllUnits(ctx context.Context) ([]*model.Unit, error) {
	var units []*model.Unit
	err := r.DB.Preload("UnitType").Find(&units).Error
	return units, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
