package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/mvandergrift/energy-gql/graph/generated"
	"github.com/mvandergrift/energy-gql/graph/model"
	"github.com/mvandergrift/energy-gql/internal/generate"
)

func (r *mutationResolver) AddFood(ctx context.Context, food model.NewFood) (*model.Food, error) {
	newFood := model.Food{Name: food.Name, Calories: food.Calories, ImgURL: food.FoodImg, UnitID: *food.UnitID}

	var result *gorm.DB

	if food.ID == nil {
		result = r.DB.Omit("ID").Create(&newFood)
	} else {
		newFood.ID = *food.ID
		result = r.DB.Save(newFood)
	}

	//log.Println(&newFood)
	r.DB.Preload("Unit").Preload("Unit.UnitType").First(&newFood, result.Value)
	return &newFood, result.Error
}

func (r *mutationResolver) DeleteFood(ctx context.Context, id int) (*model.Food, error) {
	var food model.Food
	r.DB.Where("id = ?", id).First(&food)
	err := r.DB.Delete(&food).Error
	return &food, err
}

func (r *mutationResolver) AddFoodEaten(ctx context.Context, foodEaten model.NewFoodEaten) (*model.FoodEaten, error) {
	newFoodEaten := model.FoodEaten{
		MealID: foodEaten.MealID,
		FoodID: foodEaten.FoodID,
		Size:   &foodEaten.Size,
		UnitID: foodEaten.UnitID,
	}

	var (
		matches int64
		err     error
	)
	r.DB.Model(&newFoodEaten).Where("food_id = ? and meal_id = ?", foodEaten.FoodID, foodEaten.MealID).Count(&matches)

	if matches == 0 {
		err = r.DB.Omit("ID", "Name", "Calories", "Food").Create(&newFoodEaten).Error
	} else {
		err = r.DB.Model(&newFoodEaten).Omit("ID", "Name", "Calories", "Food").Where("food_id = ? and meal_id = ?", foodEaten.FoodID, foodEaten.MealID).Update(&newFoodEaten).Error
	}

	return &newFoodEaten, err
}

func (r *mutationResolver) DeleteFoodEaten(ctx context.Context, id int) (*model.FoodEaten, error) {
	var foodEaten model.FoodEaten
	err := r.DB.Delete(&foodEaten, "id = ?", id).Error
	return &foodEaten, err
}

func (r *mutationResolver) AddMealForDay(ctx context.Context, meal model.NewMeal) (*model.Meal, error) {
	var d model.Meal

	result := r.DB.Preload("MealType").First(&d, "meal_date = ? and meal_type_id = ? and user_id = ?", meal.MealDate, meal.MealTypeID, meal.UserID)

	if result.RowsAffected == 0 {
		result = r.DB.Create(&meal)
		r.DB.Preload("MealType").First(&d, "meal_date = ? and meal_type_id = ? and user_id = ?", meal.MealDate, meal.MealTypeID, meal.UserID)
	}

	return &d, result.Error
}

func (r *mutationResolver) AddNote(ctx context.Context, note model.NewNote) (*model.Note, error) {
	newNote := model.Note{Content: note.Content, Subject: note.Subject, NoteDate: note.NoteDate, UserID: note.UserID}
	var result *gorm.DB

	if note.ID == nil {
		result = r.DB.Omit("ID").Create(&newNote)
	} else {
		newNote.ID = *note.ID
		result = r.DB.Save(newNote)
	}

	return &newNote, result.Error
}

func (r *mutationResolver) DeleteNote(ctx context.Context, id int) (*model.Note, error) {
	var note model.Note
	err := r.DB.Delete(&note, "id = ?", id).Error
	return &note, err
}

func (r *mutationResolver) DeleteWorkout(ctx context.Context, id int) (*model.Workout, error) {
	var workout model.Workout
	err := r.DB.Delete(&workout, "id = ?", id).Error
	return &workout, err
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
	err := r.DB.Preload("Unit").Preload("Unit.UnitType").Order("name").Find(&foods).Error
	return foods, err
}

func (r *queryResolver) AllUnits(ctx context.Context) ([]*model.Unit, error) {
	var units []*model.Unit
	err := r.DB.Preload("UnitType").Order("unit_type_id, name").Find(&units).Error
	return units, err
}

func (r *queryResolver) UnitsForFood(ctx context.Context, foodID *int) ([]*model.Unit, error) {
	var (
		food  model.Food
		units []*model.Unit
	)

	if err := r.DB.Preload("Unit").Preload("Unit.UnitType").First(&food, "id = ?", foodID).Error; err != nil {
		return units, err
	}

	log.Printf("food %+v", food)
	units = append(units, food.Unit)
	log.Printf("units %+v", units)
	return units, nil
}

func (r *queryResolver) Notes(ctx context.Context, userID *int, date *time.Time) ([]*model.Note, error) {
	var notes []*model.Note
	tx := r.DB.Order("note_date desc")

	if userID != nil {
		tx = tx.Where("user_id = ?", userID)
	}

	if date != nil {
		tx = tx.Where("note_date = ?", date)
	}

	err := tx.Find(&notes).Error
	return notes, err
}

func (r *queryResolver) WorkoutsForDay(ctx context.Context, userID int, date *time.Time, attributes []int) ([]*model.Workout, error) {
	var workouts []*model.Workout
	tx := r.DB.Order("start_time")

	if date != nil {
		log.Println("Date filter", date)
		tx = tx.Where("(start_time > ? OR end_time > ?) AND (start_time < ? OR end_time < ?)", date, date, date.AddDate(0, 0, 1), date.AddDate(0, 0, 1))
	}

	if len(attributes) > 0 {
		tx = tx.Where("attribute_id IN (?)", attributes)
	}

	err := tx.Preload("Activity").Where("user_id = ?", userID).Find(&workouts).Error
	return workouts, err
}

func (r *queryResolver) WorkoutQuery(ctx context.Context, filter *model.PredicateGroup) ([]*model.Workout, error) {
	var workouts []*model.Workout
	tx := r.DB.Order("activity_date desc")
	c := generate.BuildPredicateGroup(filter, generate.Collection{})
	err := tx.Preload("Activity").Where(c.Query, c.Values...).Find(&workouts).Error
	return workouts, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
