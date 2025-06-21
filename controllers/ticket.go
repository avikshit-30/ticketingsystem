package controllers

import (
	"encoding/json"
	"net/http"
	"ticketing-system/config"
	"ticketing-system/models"

	"github.com/gorilla/mux"
)

func BookTicket(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userId").(uint)

	var req struct {
		EventName string `json:"event_name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	ticket := models.Ticket{
		EventName: req.EventName,
		BookedBy:  userID,
		IsBooked:  true,
	}

	if err := config.DB.Create(&ticket).Error; err != nil {
		http.Error(w, "Failed to book ticket", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]any{
		"message": "Ticket booked successfully",
		"ticket":  ticket,
	})
}
func MyTickets(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userId").(uint)

	var tickets []models.Ticket
	if err := config.DB.Where("booked_by = ?", userID).Find(&tickets).Error; err != nil {
		http.Error(w, "Failed to fetch tickets", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(tickets)
}
func CancelTicket(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userId").(uint)
	ticketID := mux.Vars(r)["id"]

	var ticket models.Ticket
	if err := config.DB.First(&ticket, ticketID).Error; err != nil {
		http.Error(w, "Ticket not found", http.StatusNotFound)
		return
	}

	if ticket.BookedBy != userID {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	if err := config.DB.Delete(&ticket).Error; err != nil {
		http.Error(w, "Failed to delete ticket", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Ticket canceled"})
}
func CreateEvent(w http.ResponseWriter, r *http.Request) {
	var event models.Event
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := config.DB.Create(&event).Error; err != nil {
		http.Error(w, "Failed to create event", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]any{
		"message": "Event created",
		"event":   event,
	})
}
func GetEvents(w http.ResponseWriter, r *http.Request) {
	var events []models.Event
	if err := config.DB.Find(&events).Error; err != nil {
		http.Error(w, "Failed to fetch events", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(events)
}
