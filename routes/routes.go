package routes

import (
	"net/http"
	"ticketing-system/controllers"
	"ticketing-system/middleware"

	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("🎟️ Ticketing Backend Running"))
	}).Methods("GET")
	r.Handle("/dashboard", middleware.JWTAuth(http.HandlerFunc(controllers.Dashboard))).Methods("GET")

	// Auth routes
	r.HandleFunc("/signup", controllers.Signup).Methods("POST")
	r.HandleFunc("/login", controllers.Login).Methods("POST")
	r.Handle("/book-ticket", middleware.JWTAuth(http.HandlerFunc(controllers.BookTicket))).Methods("POST")
	r.Handle("/my-tickets", middleware.JWTAuth(http.HandlerFunc(controllers.MyTickets))).Methods("GET")
	r.Handle("/cancel-ticket/{id}", middleware.JWTAuth(http.HandlerFunc(controllers.CancelTicket))).Methods("DELETE")
	r.HandleFunc("/events", controllers.CreateEvent).Methods("POST")
	r.HandleFunc("/events", controllers.GetEvents).Methods("GET")



	return r
}
