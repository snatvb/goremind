package state

import (
	"eng-bot/state/events"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

type AddingWord struct {
}

func (state AddingWord) Handle(ctx *context, event string, data interface{}) State {
	if event == events.Message {
		msg := data.(*telego.Message)
		hasWord := ctx.store.HasWord(msg.Text, msg.Chat.ID)
		if hasWord {
			ctx.bot.SendMessage(tu.Message(tu.ID(msg.Chat.ID), "Word already exists"))
			return state
		}
		return AddingTranslation{word: msg.Text}
	}

	return state
}

func (state AddingWord) Name() string {
	return "AddingWord"
}
