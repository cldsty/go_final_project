package db

import (
	"database/sql"
	"os"

	_ "modernc.org/sqlite"
)

var db *sql.DB

const schema = `
CREATE TABLE scheduler (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    date CHAR(8) NOT NULL DEFAULT "",
    title VARCHAR(255) NOT NULL DEFAULT "",
    comment TEXT NOT NULL DEFAULT "",
    repeat VARCHAR(128) NOT NULL DEFAULT ""
);

CREATE INDEX idx_scheduler_date ON scheduler(date);
`

func Init(dbFile string) error {
	// Проверка существования файла БД
	_, err := os.Stat(dbFile)
	var install bool
	if err != nil {
		install = true
	}

	// Соединение с БД
	db, err = sql.Open("sqlite", dbFile)
	if err != nil {
		return err
	}

	// Создание файла БД, если его не было
	if install {
		_, err = db.Exec(schema)
		if err != nil {
			return err
		}
	}

	return nil
}

func GetDB() *sql.DB {
	return db
}
