# RemindMe Telegram Bot
Reminder bot in telegram.

This bot uses an SQLite flatfile.
But it can easily be adapted to use a HA SQL Database.

Get a token from [*BotFather*](https://telegram.me/botfather) -> 
`/newbot`

## Running the bot
```
cd remindmebot
export TOKEN="your bot token"
go mod tidy
go run .
```

## Commands
Due to limitations in golang's `time` package, the max unit is hours.
This means if you want to set a reminder for 2 days, you will need to use `48h`. This should be fixed soon.

`/add 35s sell dogecoin`

`/add 15m water plants`

`/add 5h walk dogs`