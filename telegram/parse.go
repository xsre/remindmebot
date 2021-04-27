package telegram

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/hako/durafmt"
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

// parseMessage parses the message for a command
func parseMessage(message tgbotapi.Message) (_ command, err error) {
	split := strings.Split(message.Text, " ")

	switch split[0]{
	case "/add":
		var c = CreateReminder{
			Id: generateId(),
			CreatedBy: message.From.ID,
			ChatId: message.Chat.ID,
		}

		// Parse Time
		// TODO find (or write) library that parses more than hours
		parsedTime, err := durafmt.ParseString(split[1])
		if err != nil {
			return nil, errors.New(ErrTimeParseErr)
		}
		c.Time = parsedTime.Duration()

		// Parse Item
		c.Item = strings.Join(split[2:], " ")

		return c, err
	case "/delete":
		// Create Delete Command
		var c = DeleteReminder {
			Id: split[1],
			CreatedBy: message.From.ID,
			ChatId: message.Chat.ID,
		}

		return c, err
	}

	return nil, err
}

// generateId generates an all caps letter ID like: DDHBN or KQOEU
func generateId() string {
	var letters = strings.Split("ABCDEFGHIJKLMNOPQRSTUVWXYZ", "")
	var s []string
	for i:=0; i<IDLength; i++ {
		s = append(s, letters[rand.Intn(len(letters))])
	}

	return strings.Join(s, "")
}