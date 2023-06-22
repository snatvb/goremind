package state

import (
	"goreminder/keyboard"
	"goreminder/state/events"
	"log"
	"time"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

type AddingTranslation struct {
	word string
}

func (state AddingTranslation) Handle(ctx *Context, event string, data interface{}) State {
	if event == events.Back {
		return AddingWord{}
	}

	if event == events.Message {
		msg := data.(*telego.Message)
		log.Printf("AddingTranslation: %s", msg.Text)
		if ctx.Store.HasWord(state.word, msg.Chat.ID) {
			log.Printf("DEBUG AddingTranslation: %s already exists", state.word)
			return state
		}

		now := time.Now()
		inHour := now.Add(time.Minute)
		success := ctx.Store.AddNewWord(state.word, msg.Text, msg.Chat.ID, inHour)
		if success {
			log.Printf("AddingTranslation: %s added", state.word)
			return AddingWordSuccess{}
		}

		log.Printf("FAILED AddingTranslation: %s", state.word)
		ctx.Bot.SendMessage(
			tu.Message(tu.ID(msg.Chat.ID), "Error adding word"),
		)
		return Idle{}
	}
	return state
}

func (state AddingTranslation) OnEnter(fsm *FSM, context *Context, from State) {
	context.Bot.SendMessage(tu.MessageWithEntities(
		tu.ID(context.ChatId),
		tu.Entity("Write translation for the word "),
		tu.Entity(state.word).Code(),
	).WithReplyMarkup(keyboard.BackOnly()))
}

func (state AddingTranslation) Name() string {
	return "AddingTranslation"
}
