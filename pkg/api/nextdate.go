package api

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// NextDate вычисляет следующую дату повторения задачи
func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	// Проверка наличия правила повторения
	if repeat == "" {
		return "", fmt.Errorf("пустое правило повторения")
	}

	// Парсинг исходной даты
	date, err := time.Parse(DateFormat, dstart)
	if err != nil {
		return "", fmt.Errorf("некорректная дата dstart: %v", err)
	}

	// Разбитие правила повторения на части
	parts := strings.Split(strings.TrimSpace(repeat), " ")
	if len(parts) == 0 {
		return "", fmt.Errorf("некорректный формат правила повторения")
	}

	// Обработка правила повторения
	switch parts[0] {
	// Повторение через указанное количество дней
	case "d":
		if len(parts) != 2 {
			return "", fmt.Errorf("некорректный формат правила d: ожидается 'd <число>'")
		}

		interval, err := strconv.Atoi(parts[1])
		if err != nil {
			return "", fmt.Errorf("некорректное число дней: %v", err)
		}

		if interval <= 0 || interval > 400 {
			return "", fmt.Errorf("интервал дней должен быть от 1 до 400, получено: %d", interval)
		}

		for {
			date = date.AddDate(0, 0, interval)
			if afterNow(date, now) {
				break
			}
		}

	// Ежегодное повторение
	case "y":

		for {
			date = date.AddDate(1, 0, 0)
			if afterNow(date, now) {
				break
			}
		}

	default:
		return "", fmt.Errorf("неподдерживаемый формат правила повторения: %s", parts[0])
	}

	return date.Format(DateFormat), nil
}

// afterNow проверяет, больше ли первая дата второй, не учитывая время
func afterNow(date, now time.Time) bool {

	dateNormalized := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	nowNormalized := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	return dateNormalized.After(nowNormalized)
}

// nextDateHandler обрабатывает HTTP запросы для расчета следующей даты
func nextDateHandler(w http.ResponseWriter, r *http.Request) {
	// Получение параметров через FormValue
	nowStr := r.FormValue("now")
	dateStr := r.FormValue("date")
	repeatStr := r.FormValue("repeat")

	// Если параметр now не определен, берется текущая дата
	if nowStr == "" {
		nowStr = time.Now().Format(DateFormat)
	}

	// Проверка наличия даты
	if dateStr == "" {
		http.Error(w, "отсутствует обязательный параметр date", http.StatusBadRequest)
		return
	}

	// Парсинг текущего времени
	now, err := time.Parse(DateFormat, nowStr)
	if err != nil {
		http.Error(w, "некорректный формат даты now", http.StatusBadRequest)
		return
	}

	// Вычисление следующей даты
	nextDate, err := NextDate(now, dateStr, repeatStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Возвращение результата в формате 20060102
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte(nextDate))
}
