package main

import (
	"eng-bot/clients"
	"eng-bot/keyboard"
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
	bot := newBot()
	users := clients.New(db, bot)

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

	defer handler.Stop()
	defer bot.StopLongPolling()

	handler.HandleCallbackQuery(func(bot *telego.Bot, query telego.CallbackQuery) {
		user := users.GetOrAdd(query.Message.Chat.ID)
		user.State.Handle(events.AddWord, nil)
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
		user := users.GetOrAdd(message.Chat.ID)
		user.State.Handle(events.Message, &message)
		currentState := user.State.CurrentState().Name()
		if currentState == "AddingTranslation" {
			_, _ = bot.SendMessage(tu.Message(tu.ID(message.Chat.ID), "Write translation"))
		} else if currentState == "AddingWordSuccess" {
			word := db.LastWord()
			_, _ = bot.SendMessage(tu.Message(tu.ID(message.Chat.ID), fmt.Sprintf("Word %s added!", word.Word)))
			user.State.Handle(events.Reset, nil)
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
