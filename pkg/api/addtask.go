package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"go-final-project/pkg/db"
)

// addTaskHandler обрабатывает POST-запросы для добавления задач
func addTaskHandler(w http.ResponseWriter, r *http.Request) {
	// Десериализация JSON в структуру Task
	var task db.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		log.Printf("Ошибка десериализации JSON: %v", err)
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "Ошибка десериализации JSON"})
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

	// Сохранение задачи в базу данных
	id, err := db.AddTask(&task)
	if err != nil {
		log.Printf("Ошибка добавления задачи в базу данных: %v", err)
		writeJSON(w, http.StatusInternalServerError, map[string]string{"error": "Ошибка добавления задачи в базу данных"})
		return
	}

	// Возвращение идентификатора созданной задачи
	writeJSON(w, http.StatusOK, map[string]string{"id": strconv.FormatInt(id, 10)})
}

// checkDate проверяет корректность даты и корректирует её при необходимости
func checkDate(task *db.Task) error {
	now := time.Now()

	// Установка сегоднящнего числа при отсутствии даты
	if task.Date == "" {
		task.Date = now.Format(DateFormat)
		return nil
	}

	// Проверка корректности формата даты
	t, err := time.Parse(DateFormat, task.Date)
	if err != nil {
		return err
	}

	// Если сегодня больше указанной даты
	if afterNow(now, t) {
		if len(task.Repeat) == 0 {
			// Повторения нет - сегодняшнее число
			task.Date = now.Format(DateFormat)
		} else {
			// иначе - следующая дата
			next, err := NextDate(now, task.Date, task.Repeat)
			if err != nil {
				return err
			}
			task.Date = next
		}
	}

	return nil
}

// writeJSON записывает данные в формате JSON в ответ
func writeJSON(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
