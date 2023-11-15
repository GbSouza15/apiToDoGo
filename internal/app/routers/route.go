package routers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/GbSouza15/apiToDoGo/internal/app/handlers"
	"github.com/GbSouza15/apiToDoGo/internal/authenticator"
	"github.com/gorilla/mux"
)

func RoutesApi(db *sql.DB) error {

	r := mux.NewRouter()
	h := handlers.New(db)
	// GET
	r.HandleFunc("/tasks", authenticator.CheckTokenIsValid(h.GetTasksForUserHandler)).Methods(http.MethodGet)
	// POST
	r.HandleFunc("/register", h.RegisterUserHandler).Methods(http.MethodPost)
	r.HandleFunc("/login", h.LoginUserHandler).Methods(http.MethodPost)
	r.HandleFunc("/tasks/new", authenticator.CheckTokenIsValid(h.CreateTasks)).Methods(http.MethodPost)
	// DELETE
	r.HandleFunc("/{taskId}/task/delete", authenticator.CheckTokenIsValid(h.DeleteTask)).Methods(http.MethodDelete)

	http.Handle("/", r)
	fmt.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		return err
	}

	return nil
}
