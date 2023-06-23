package state

import (
	"goreminder/keyboard"
	"goreminder/state/events"
	"log"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	"github.com/steebchen/prisma-client-go/runtime/types"
)

type RemoveWord struct {
}

func (state RemoveWord) Handle(context *Context, event string, data interface{}) State {
	if event == events.Back {
		return Idle{}
	}

	if event == events.Message {
		msg := data.(*telego.Message)
		word := msg.Text
		log.Printf("Removing word: %s", word)
		success := context.Store.RemoveWord(word, types.BigInt(context.ChatId))
		if success {
			context.Bot.SendMessage(tu.Message(
				tu.ID(context.ChatId),
				"Word removed successfully",
			).WithReplyMarkup(keyboard.BackOnly()))
		} else {
			log.Printf("FAILURE Removing word: %s", word)
			context.Bot.SendMessage(tu.Message(
				tu.ID(context.ChatId),
				"Word not found",
			).WithReplyMarkup(keyboard.BackOnly()))
		}
	}

	return state
}

func (state RemoveWord) OnEnter(fsm *FSM, context *Context, from State) {
	context.Bot.SendMessage(tu.Message(
		tu.ID(context.ChatId),
		"Write a word that you want to remove",
	).WithReplyMarkup(keyboard.BackOnly()))
}

func (s RemoveWord) Name() string {
	return "WordsView"
}
