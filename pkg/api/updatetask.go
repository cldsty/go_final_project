package api

import (
	"encoding/json"
	"log"
	"net/http"

	"go-final-project/pkg/db"
)

// updateTaskHandler обрабатывает PUT-запросы для обновления задачи
func updateTaskHandler(w http.ResponseWriter, r *http.Request) {
	// Десериализация JSON в структуру Task
	var task db.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		log.Printf("Ошибка десериализации JSON: %v", err)
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Ошибка десериализации JSON"})
		return
	}

	// Проверка указания ID задачи
	if task.ID == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Не указан идентификатор задачи"})
		return
	}

	// Проверка наличия заголовка задачи
	if task.Title == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Не указан заголовок задачи"})
		return
	}

	// Проверка и корректировка даты
	if err := checkDate(&task); err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	// Обновление задачи в базе данных
	err := db.UpdateTask(&task)
	if err != nil {
		log.Printf("Ошибка обновления задачи с ID %s: %v", task.ID, err)
		writeJSON(w, http.StatusNotFound, map[string]string{"error": "Задача не найдена"})
		return
	}

	writeJSON(w, http.StatusOK, map[string]interface{}{})
}
