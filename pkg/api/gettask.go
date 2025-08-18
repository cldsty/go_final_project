package api

import (
	"net/http"

	"go-final-project/pkg/db"
)

// getTaskHandler обрабатывает GET-запросы для получения задачи по ID
func getTaskHandler(w http.ResponseWriter, r *http.Request) {
	// Получение ID задачи из параметра
	id := r.FormValue("id")
	if id == "" {
		writeJSON(w, map[string]string{"error": "Не указан идентификатор"})
		return
	}

	// Получение задачи из базы данных
	task, err := db.GetTask(id)
	if err != nil {
		writeJSON(w, map[string]string{"error": "Задача не найдена"})
		return
	}

	// Возвращение задачи в формате JSON
	writeJSON(w, task)
}
