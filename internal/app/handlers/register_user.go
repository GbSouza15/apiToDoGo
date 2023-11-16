package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/GbSouza15/apiToDoGo/internal/app/models"
	"github.com/GbSouza15/apiToDoGo/internal/app/response"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (h handler) RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		response.SendResponse(500, []byte("Error reading the request body"), w)
		fmt.Println(err.Error())
		return
	}
	defer r.Body.Close()

	var newUser models.User
	userId := uuid.NewString()

	if err := json.Unmarshal(body, &newUser); err != nil {
		response.SendResponse(500, []byte("Error decoding JSON"), w)
		fmt.Println(err.Error())
		return
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), 14)
	if err != nil {
		fmt.Println("Error in Hash.")
	}

	_, err = h.DB.Exec("INSERT INTO tdlist.users (id, name, email, password) VALUES ($1, $2, $3, $4)", userId, newUser.Name, newUser.Email, bytes)
	if err != nil {
		response.SendResponse(500, []byte("Error registering user."), w)
		fmt.Println(err.Error())
		return
	}

	response.SendResponse(200, []byte("User registered successfully."), w)
}
