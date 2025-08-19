package api

import (
	"log"
	"net/http"

	"go-final-project/pkg/db"
)

// getTaskHandler обрабатывает GET-запросы для получения задачи по ID
func getTaskHandler(w http.ResponseWriter, r *http.Request) {
	// Получение ID задачи из параметра
	id := r.FormValue("id")
	if id == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Не указан идентификатор"})
		return
	}

	// Получение задачи из базы данных
	task, err := db.GetTask(id)
	if err != nil {
		log.Printf("Ошибка получения задачи с ID %s: %v", id, err)
		writeJSON(w, http.StatusNotFound, map[string]string{"error": "Задача не найдена"})
		return
	}

	// Возвращение задачи в формате JSON
	writeJSON(w, http.StatusOK, task)
}
