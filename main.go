package main

import (
	"os"
	"remindmebot/database"
	"remindmebot/telegram"
	"sync"
)

var (
	TOKEN = os.Getenv("TOKEN")
)

func main() {

	err := database.Init()
	if err != nil {
		panic(err)
	}

	bot, err := telegram.Init(TOKEN)
	if err != nil {
		panic(err)
	}

	// Message Handler to listen to commands from Telegram bot
	messageListener1 := telegram.NewMessageHandler(bot)
	go messageListener1.ListenToTelegram()


	// Message Handler to send message to the bot
	messageListener2 := telegram.NewMessageHandler(bot)
	go messageListener2.ListenToDB()


	// makes sure program doesn't exit
	s := sync.WaitGroup{}
	s.Add(1)
	s.Wait()

}
