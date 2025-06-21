package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"ticketing-system/config"
	"ticketing-system/models"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	log.Println("🧠 Inside Signup Handler")
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	// Hash password
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	user.Password = string(hashedPassword)

	config.DB.Create(&user)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User created"})
}

func Login(w http.ResponseWriter, r *http.Request) {
	log.Println("🔐 Inside Login Handler")

	var creds models.User
	json.NewDecoder(r.Body).Decode(&creds)

	var user models.User
	result := config.DB.Where("email = ?", creds.Email).First(&user)
	if result.Error != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password))
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// ✅ Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.ID,
		"exp":    time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		http.Error(w, "Could not generate token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}
func Dashboard(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userId").(uint)

	json.NewEncoder(w).Encode(map[string]any{
		"message": "Welcome to your dashboard!",
		"userId":  userID,
	})
}
