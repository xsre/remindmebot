package database

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"time"
)

var db *sql.DB

const (
	dbFlatFile = "sqlite.db"
)

type SelectResponse struct {
	Item string
	CreatedBy string
	Id string
	ChatId int64
}

// Init initializes the database
func Init() error {
	log.Println("Initializing DB")

	_, err := os.Stat(dbFlatFile)
	if os.IsNotExist(err) {
		log.Println("creating sqlite db flat file")
		_, err := os.Create(dbFlatFile)
		if err != nil {
			log.Println(err)
			return err
		}
	}

	db, err = sql.Open("sqlite3", dbFlatFile)
	if err != nil {
		log.Println(err)
		return err
	}

	log.Println("creating reminder table")
	err = createReminderTable()
	if err != nil {
		log.Println(err)
	}

	return err

}

// createReminderTable creates the reminder table in the db if it hadn't been created already
func createReminderTable() error {
	createTableSQL := `CREATE TABLE reminder (
			"id" TEXT,
			"endsOn" INT,
			"item" TEXT,
			"createdBy" INT,
            "chatId" INT
		);`

	statement, err := db.Prepare(createTableSQL)
	if err != nil {
		// doing it this way instead of err.Error() to avoid possible nil pointer dereference
		if fmt.Sprintf("%s", err) == "table reminder already exists" {
			return nil
		}
	}
	if err != nil {
		return err
	}

	_, err = statement.Exec()
	return err
}

// InsertReminder inserts a reminder into the database
func InsertReminder(id string, endsOn int64, item string, createdBy int, chatId int64) error {
	insertSQL := `INSERT INTO reminder (id, endsOn, item, createdBy, ChatId) VALUES (?, ?, ?, ?, ?)`
	statement, err := db.Prepare(insertSQL)
	if err != nil {
		return err
	}

	_, err = statement.Exec(id, endsOn, item, createdBy, chatId)
	return err
}

// CheckReminders looks trough the database if there are any reminders that have expired
func CheckReminders() (sr []SelectResponse, err error) {
	t := time.Now().Unix()
	selectSQL := `SELECT item, createdBy, chatId, id FROM reminder WHERE endsOn <= ?`
	rows, err := db.Query(selectSQL, t)
	if err != nil {
		return sr, err
	}

	for rows.Next() {
		var r SelectResponse
		err = rows.Scan(&r.Item, &r.CreatedBy, &r.ChatId, &r.Id)
		if err != nil {
			log.Println(err)
			continue
		}

		sr = append(sr, r)
	}

	return sr, err
}

// DeleteReminder deletes a reminder from the database based on its ID
func DeleteReminder(id string) error {
	deleteSQL := `DELETE FROM reminder WHERE id IS ?`
	statement, err := db.Prepare(deleteSQL)
	if err != nil {
		return err
	}

	_, err = statement.Exec(id)

	return err
}

func (r SelectResponse) FormatMessage() string {
	return fmt.Sprintf("Reminder: '%s'", r.Item)
}