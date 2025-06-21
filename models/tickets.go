package models

import "gorm.io/gorm"

type Ticket struct {
	gorm.Model
	EventName string `json:"event_name"`
	BookedBy  uint   `json:"booked_by"` 
	IsBooked  bool   `json:"is_booked"`
}
