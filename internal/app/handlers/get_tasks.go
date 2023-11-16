package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/GbSouza15/apiToDoGo/internal/app/models"
	"github.com/GbSouza15/apiToDoGo/internal/app/response"
	"github.com/GbSouza15/apiToDoGo/internal/authenticator"
)

func (h handler) GetTasksForUserHandler(w http.ResponseWriter, r *http.Request) {
	userId := authenticator.UserIDFromContext(r.Context())

	res, err := h.DB.Query("SELECT id, title, description FROM tdlist.tasks WHERE user_id = $1", userId)
	if err != nil {
		response.SendResponse(500, []byte("Error fetching tasks"), w)
		return
	}

	tasks := []models.Task{}

	for res.Next() {
		var task models.Task
		if err := res.Scan(&task.ID, &task.Title, &task.Description); err != nil {
			response.SendResponse(500, []byte("Error in articles"), w)
			fmt.Println("Error scanning articles: ", err.Error())
			return
		}
		tasks = append(tasks, task)
	}

	responseJSON, err := json.Marshal(tasks)
	if err != nil {
		response.SendResponse(500, []byte("Error converting to JSON"), w)
		return
	}

	response.SendResponse(200, responseJSON, w)
}
