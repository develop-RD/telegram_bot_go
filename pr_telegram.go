package main

import (
	"database/sql"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	_ "github.com/lib/pq"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("1302311660:AAH9Cs-4Ln_z4cdYx0k5UzHWKHo9u9YwAwM")
	if err != nil {
		panic(err)
	}

	bot.Debug = true

	// Connect to the PostgreSQL database
	connStr := "user=myuser dbname=mydb password=mypassword sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Create a new UpdateConfig struct with an offset of 0. Offsets are used
	// to make sure Telegram knows we've handled previous values and we don't
	// need them repeated.
	updateConfig := tgbotapi.NewUpdate(0)

	// Tell Telegram we should wait up to 30 seconds on each request for an
	// update. This way we can get information just as quickly as making many
	// frequent requests without having to send nearly as many.
	updateConfig.Timeout = 30

	// Start polling Telegram for updates.
	updates := bot.GetUpdatesChan(updateConfig)

	// Let's go through each update that we're getting from Telegram.
	for update := range updates {
		// Telegram can send many types of updates depending on what your Bot
		// is up to. We only want to look at messages for now, so we can
		// discard any other updates.
		if update.Message == nil {
			continue
		}

		// Create a new MessageConfig. We don't have text yet,
		// so we leave it empty.
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		msg.Text = "Hi chel"
		switch {
		case update.Message.Text == "/help":
			// Запрос списка команд
			msg.Text = " Вот список команд:\n"
			msg.Text = msg.Text + "/add - это добавление\n"
		}
		// Okay, we're sending our message off! We don't care about the message
		// we just sent, so we'll discard it.
		//	if _, err := bot.Send(msg); err != nil {
		// Note that panics are a bad way to handle errors. Telegram can
		// have service outages or network errors, you should retry sending
		// messages or more gracefully handle failures.
		//	panic(err)
		//}

		// send db data
		// Perform a sample query
		//		rows, err := db.Query("SELECT * FROM cinema_rom")
		//	if err != nil {
		//		panic(err)
		//	}
		//	defer rows.Close()

		// Iterate through the results
		//for rows.Next() {
		//	var id int
		//	var name string
		//	var allmessage string
		//	var id_cinema int
		//	if err := rows.Scan(&id, &name, &id_cinema); err != nil {
		//		panic(err)
		//	}
		//	fmt.Printf("ID: %d, Name: %s ID_cinema %d\n",
		//		id, name, id_cinema)
		//	allmessage = name
		//	msg.Text = allmessage
		// Okay, we're sending our message off! We don't care about the message
		// we just sent, so we'll discard it.
		if _, err := bot.Send(msg); err != nil {
			// Note that panics are a bad way to handle errors. Telegram can
			// have service outages or network errors, you should retry sending
			// messages or more gracefully handle failures.
			panic(err)
		}
		//}
	}
}
