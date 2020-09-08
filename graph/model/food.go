package model

import "time"

type Food struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Calories  int     `json:"calories"`
	ImgURL    *string `json:"imgUrl"`
	Unit      *Unit   `json:"unit"`
	UnitID    int
	CreatedAt *time.Time
	UpdatedAt *time.Time
	DeletedAt *time.Time
}
