package model

type Unit struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	UnitTypeID int
	UnitType   *UnitType `json:"unitType"`
}
