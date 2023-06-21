package state

import (
	"eng-bot/state/events"

	"github.com/mymmrac/telego"
)

type AddingTranslation struct {
	word string
}

func (state AddingTranslation) Handle(ctx *context, event string, data interface{}) State {
	if event == events.Message {
		msg := data.(*telego.Message)
		println("AddingTranslation")
		word := ctx.store.NewWord(state.word, msg.Text, msg.Chat.ID)
		ctx.store.Save(word)
		println("Saved")

		return AddingWordSuccess{}
	}
	return state
}

func (state AddingTranslation) Name() string {
	return "AddingTranslation"
}
