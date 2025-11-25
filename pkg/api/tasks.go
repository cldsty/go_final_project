package api

import (
	"log"
	"net/http"

	"go-final-project/pkg/db"
)

// TasksResp структура для ответа со списком задач
type TasksResp struct {
	Tasks []*db.Task `json:"tasks"`
}

// tasksHandler обрабатывает GET-запросы для получения списка задач
func tasksHandler(w http.ResponseWriter, r *http.Request) {

	tasks, err := db.Tasks(20)
	if err != nil {
		log.Printf("Ошибка получения задач из базы данных: %v", err)
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "Ошибка получения задач из базы данных"})
		return
	}

	writeJSON(w, http.StatusOK, TasksResp{
		Tasks: tasks,
	})
}
