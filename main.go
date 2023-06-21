package main

import (
	"eng-bot/keyboard"
	"eng-bot/state"
	"eng-bot/state/events"
	"eng-bot/store"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

func main() {
	loadEnv()
	db := store.New(&store.Options{DbPath: os.Getenv("DB_FILENAME")})
	fsm := state.NewFSM(db)

	bot := newBot()

	me, err := bot.GetMe()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Bot user: %+v\n", me.FirstName)

	updates, _ := bot.UpdatesViaLongPolling(nil)
	handler, err := th.NewBotHandler(bot, updates)
	if err != nil {
		fmt.Println(err)
	}

	// lastMessage := db.LastMessage()
	// fmt.Printf("Last message: %+v\n", lastMessage.Text)

	// handler.HandleMessage(func(bot *telego.Bot, message telego.Message) {
	// 	// Get chat ID from the message
	// 	chatID := tu.ID(message.Chat.ID)

	// 	msg := tu.Message(
	// 		tu.ID(chatID.ID),
	// 		"Hello World",
	// 	)

	// 	// Copy sent messages back to the user
	// 	_, err := bot.SendMessage(msg)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	db.Save(&state.Message{ChatId: message.Chat.ID, Text: message.Text})

	// 	fmt.Printf("Message: %+v\n", message.Text)
	// })

	defer handler.Stop()
	defer bot.StopLongPolling()

	// handleAddWord := func(bot *telego.Bot, message telego.Message) {
	// 	fsm.Handle(events.AddDescription, message.Text)
	// }

	handler.HandleCallbackQuery(func(bot *telego.Bot, query telego.CallbackQuery) {
		fsm.Handle(events.AddWord, nil)
		_, _ = bot.SendMessage(tu.Message(tu.ID(query.Message.Chat.ID), "Write word that you want to add"))
	}, th.AnyCallbackQueryWithMessage(), th.CallbackDataEqual(keyboard.AddWord))

	handler.Handle(func(bot *telego.Bot, update telego.Update) {
		// Send message
		_, _ = bot.SendMessage(tu.Message(
			tu.ID(update.Message.Chat.ID),
			fmt.Sprintf("Hello %s! I'll help you to remember english words.", update.Message.From.FirstName),
		).WithReplyMarkup(keyboard.CreateWords()))
	}, th.CommandEqual("start"))

	handler.HandleMessage(func(bot *telego.Bot, message telego.Message) {
		fsm.Handle(events.Message, &message)
		if fsm.CurrentState().Name() == "AddingWord" {
			_, _ = bot.SendMessage(tu.Message(tu.ID(message.Chat.ID), "Write translation"))
		} else if fsm.CurrentState().Name() == "AddingWordSuccess" {
			word := db.LastWord()
			_, _ = bot.SendMessage(tu.Message(tu.ID(message.Chat.ID), fmt.Sprintf("Word %s added!", word.Word)))
			fsm.Handle(events.Reset, nil)
		}
	})

	handler.Start()

}

func newBot() *telego.Bot {
	bot, err := telego.NewBot(os.Getenv("TOKEN"))
	if err != nil {
		log.Fatal(err)
	}
	return bot
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
