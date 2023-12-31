package main

import (
	"fmt"
	"goreminder/clients"
	"goreminder/keyboard"
	"goreminder/reminder"
	"goreminder/state/events"
	"goreminder/store"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
)

func main() {
	loadEnv()
	db := store.New()
	bot := newBot()
	users := clients.New(db, bot)
	remind := reminder.Reminder{
		Options: reminder.Options{
			Interval: 5 * time.Second,
		},
		Store:   db,
		Clients: users,
	}

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

	remind.Start()
	defer handler.Stop()
	defer bot.StopLongPolling()
	defer remind.Stop()

	fmt.Printf("Wait...\n")

	handler.HandleCallbackQuery(func(bot *telego.Bot, query telego.CallbackQuery) {
		user := users.GetOrAdd(query.Message.Chat.ID)
		user.State.Handle(events.AddWord, nil)
		_ = bot.AnswerCallbackQuery(tu.CallbackQuery(query.ID).WithText("Started adding word"))
	}, th.AnyCallbackQueryWithMessage(), th.CallbackDataEqual(keyboard.AddWord))

	subscribeKeyboard(handler, users, keyboard.Back, events.Back)
	subscribeKeyboard(handler, users, keyboard.WordList, events.WordList)
	subscribeKeyboard(handler, users, keyboard.RemoveWord, events.RemoveWord)
	subscribeKeyboard(handler, users, keyboard.Forgot, events.Forgot)

	handler.Handle(func(bot *telego.Bot, update telego.Update) {
		// Send message
		_, _ = bot.SendMessage(tu.Message(
			tu.ID(update.Message.Chat.ID),
			fmt.Sprintf("Hello %s! I'll help you to remember english words.", update.Message.From.FirstName),
		).WithReplyMarkup(keyboard.WordsControls()))
	}, th.CommandEqual("start"))

	handler.HandleMessage(func(bot *telego.Bot, message telego.Message) {
		user := users.GetOrAdd(message.Chat.ID)
		user.State.Handle(events.Message, &message)
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
	env := os.Getenv("ENV")

	if env == "development" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
}

func subscribeKeyboard(handler *th.BotHandler, users *clients.Clients, keyboardEvent string, stateEvent string) {
	handler.HandleCallbackQuery(func(bot *telego.Bot, query telego.CallbackQuery) {
		user := users.GetOrAdd(query.Message.Chat.ID)
		user.State.Handle(stateEvent, nil)
		_ = bot.AnswerCallbackQuery(tu.CallbackQuery(query.ID).WithText("Done"))
	}, th.AnyCallbackQueryWithMessage(), th.CallbackDataEqual(keyboardEvent))
}
