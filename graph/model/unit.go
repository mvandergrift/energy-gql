package model

type Unit struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	UnitTypeId int
	UnitType   *UnitType `json:"unitType"`
}
