package api

import (
	"net/http"

	"go-final-project/pkg/db"
)

// deleteTaskHandler обрабатывает DELETE-запросы для удаления задачи
func deleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	// Получение ID задачи из параметра
	id := r.FormValue("id")
	if id == "" {
		writeJSON(w, map[string]string{"error": "Не указан идентификатор"})
		return
	}

	// Удаление задачи из базы данных
	err := db.DeleteTask(id)
	if err != nil {
		writeJSON(w, map[string]string{"error": "Задача не найдена"})
		return
	}

	// Возвращение пустого JSON
	writeJSON(w, map[string]interface{}{})
}
