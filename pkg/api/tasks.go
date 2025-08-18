package api

import (
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
		writeJSON(w, map[string]string{"error": "Ошибка получения задач из базы данных"})
		return
	}

	writeJSON(w, TasksResp{
		Tasks: tasks,
	})
}
