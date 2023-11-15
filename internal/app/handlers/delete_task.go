package handlers

import (
	"net/http"

	"github.com/GbSouza15/apiToDoGo/internal/app/response"
	"github.com/gorilla/mux"
)

func (h handler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskId := vars["taskId"]

	_, err := h.DB.Exec("DELETE FROM tdlist.tasks WHERE id = $1", taskId)
	if err != nil {
		response.SendResponse(500, []byte("Erro ao deletar tarefa"), w)
		return
	}

	response.SendResponse(200, []byte("Tarefa deletada com sucesso"), w)
}
