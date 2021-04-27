package telegram

import (
	"fmt"
	"remindmebot/database"
	"time"
)

// CreateReminder command
type CreateReminder struct {
	Id string
	Time time.Duration
	Item string
	CreatedBy int
	ChatId int64
}

// DeleteReminder command
type DeleteReminder struct {
	Id string
	CreatedBy int
	ChatId int64
}

type command interface {
	Execute() error
	ResponseMessage() string
}

func (c CreateReminder) Execute() (err error) {
	err = database.InsertReminder(c.Id, c.EndsOnUnix(), c.Item, c.CreatedBy, c.ChatId)
	return err
}

func (c DeleteReminder) Execute() (err error) {
	return err
}

func (c CreateReminder) ResponseMessage() string {
	return fmt.Sprintf("Set a timer on: '%v' for: '%v'", c.EndsOn(), c.Item)
}

func (c DeleteReminder) ResponseMessage() string {
	return ""
}

// Get DateTime when the reminder ends
func (c CreateReminder) EndsOn() string {
	added := time.Now().Add(c.Time)
	return added.Format(time.ANSIC)
}

// Get Unix time when reminder ends
func (c CreateReminder) EndsOnUnix() int64 {
	added := time.Now().Add(c.Time)
	return added.Unix()
}