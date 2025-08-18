package api

import (
	"net/http"
	"time"

	"go-final-project/pkg/db"
)

// doneTaskHandler обрабатывает POST-запросы для отметки задачи как выполненной
func doneTaskHandler(w http.ResponseWriter, r *http.Request) {
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

	// Удаление задач без повторения
	if task.Repeat == "" {
		err = db.DeleteTask(id)
		if err != nil {
			writeJSON(w, map[string]string{"error": "Ошибка удаления задачи"})
			return
		}
		writeJSON(w, map[string]interface{}{})
		return
	}

	// Рассчет следующей даты задачи
	now := time.Now()
	nextDate, err := NextDate(now, task.Date, task.Repeat)
	if err != nil {
		writeJSON(w, map[string]string{"error": "Ошибка расчета следующей даты"})
		return
	}

	// Обновление даты задачи
	err = db.UpdateDate(nextDate, id)
	if err != nil {
		writeJSON(w, map[string]string{"error": "Ошибка обновления даты задачи"})
		return
	}

	// Возвращение пустого JSON
	writeJSON(w, map[string]interface{}{})
}
