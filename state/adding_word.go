package state

import (
	"eng-bot/state/events"

	"github.com/mymmrac/telego"
)

type AddingWord struct {
}

func (state AddingWord) Handle(_ *context, event string, data interface{}) State {
	if event == events.Message {
		msg := data.(*telego.Message)
		return AddingTranslation{word: msg.Text}
	}

	return state
}

func (state AddingWord) Name() string {
	return "AddingWord"
}
