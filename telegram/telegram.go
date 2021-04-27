package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"remindmebot/database"
	"time"
)

// MessageHandler is an abstraction over tgbotapi to be able to implement functions on it.
type MessageHandler struct {
	bot *tgbotapi.BotAPI
}

// NewMessageHandler creates a new MessageHandler
func NewMessageHandler (b *tgbotapi.BotAPI) MessageHandler {
	return MessageHandler{
		bot: b,
	}
}

// Init initializes new bot with token
func Init(token string) (bot *tgbotapi.BotAPI, err error){
	bot, err = tgbotapi.NewBotAPI(token)
	if err != nil {
		return bot, err
	}
	return bot, err
}


func (m MessageHandler) ListenToTelegram () {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := m.bot.GetUpdatesChan(u)
	if err != nil {
		log.Println(err)
	}

	for update := range updates {
		if update.Message == nil {
			continue
		}

		command, err := parseMessage(*update.Message)
		if err != nil {
			fmt.Println(err)
			continue
		}


		var response string

		response = command.ResponseMessage()

		err = command.Execute()
		if err != nil {
			fmt.Println(err)
			response = "Something went wrong, please try again!"
		}

		//msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		//msg.ReplyToMessageID = update.Message.MessageID

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)

		_, err = m.bot.Send(msg)
		if err != nil {
			log.Println(err)
		}
	}

}

func (m MessageHandler) ListenToDB () {
	for {
		sr, err := database.CheckReminders()
		if err != nil {
			log.Println(err)
			continue
		}

		if len(sr) == 0 {
			continue
		}

		for _, r := range sr {

			err = database.DeleteReminder(r.Id)
			if err != nil {
				log.Println(err)
				continue
			}

			fmtmessage := r.FormatMessage()
			_, err = m.bot.Send(tgbotapi.NewMessage(r.ChatId, fmtmessage))
			if err != nil {
				log.Println(err)
			}

			log.Println(fmtmessage)
		}

		time.Sleep(time.Second * 1)
	}
}
