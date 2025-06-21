package models

type Event struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
	Date     string `json:"date"`
}
