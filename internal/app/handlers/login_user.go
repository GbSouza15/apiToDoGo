package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/GbSouza15/apiToDoGo/internal/app/models"
	"github.com/GbSouza15/apiToDoGo/internal/app/response"
	"golang.org/x/crypto/bcrypt"
)

func (h handler) LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	var userLogin models.UserLogin
	var user models.User

	body, err := io.ReadAll(r.Body)
	if err != nil {
		response.SendResponse(500, []byte("Error logging in"), w)
	}
	defer r.Body.Close()

	if err := json.Unmarshal(body, &userLogin); err != nil {
		response.SendResponse(500, []byte("JSON decoding error"), w)
		return
	}

	err = h.DB.QueryRow("SELECT * FROM tdlist.users WHERE email = $1", userLogin.Email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			response.SendResponse(404, []byte("No record of this user"), w)
			return
		}
		response.SendResponse(401, []byte("Server error"), w)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userLogin.Password)); err != nil {
		response.SendResponse(401, []byte("Incorrect password"), w)
		fmt.Println(err)
		return
	}

	claims := &models.Claims{UserId: user.ID.String(), RegisteredClaims: jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
	}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := os.Getenv("SECRET")
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		response.SendResponse(500, []byte("Error generating the token"), w)
		return
	}

	tokenResponse := map[string]string{"token": tokenString}
	responseJSON, err := json.Marshal(tokenResponse)
	if err != nil {
		response.SendResponse(500, []byte("Error encoding the JSON"), w)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: time.Now().Add(time.Hour * 24),
	})
	response.SendResponse(200, responseJSON, w)
}
