package model

import "time"

type Food struct {
	ID        int      `json:"id"`
	Name      string   `json:"name"`
	Calories  int      `json:"calories"`
	Fat       *float64 `json:"fat"`
	SatFat    *float64 `json:"satFat"`
	Carbs     *float64 `json:"carbs"`
	Fiber     *float64 `json:"fiber"`
	Sugar     *float64 `json:"sugar"`
	Sodium    *float64 `json:"sodium"`
	Protein   *float64 `json:"protein"`
	ImgURL    *string  `json:"imgUrl"`
	Unit      *Unit    `json:"unit"`
	UnitID    int
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}
