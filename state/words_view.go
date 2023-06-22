package state

import (
	"goreminder/keyboard"
	"goreminder/state/events"

	tu "github.com/mymmrac/telego/telegoutil"
)

type WordsView struct {
}

func (state WordsView) Handle(_ *Context, event string, data interface{}) State {
	if event == events.Back {
		return Idle{}
	}
	return state
}

func (state WordsView) OnEnter(fsm *FSM, context *Context, from State) {
	words := context.Store.GetWords(context.ChatId, 0, 100)
	wordsTextsMsg := "Words:\n"
	for _, word := range words {
		wordsTextsMsg += word.Word + "\n"
	}
	context.Bot.SendMessage(tu.Message(
		tu.ID(context.ChatId),
		wordsTextsMsg,
	).WithReplyMarkup(keyboard.BackOnly()))
}

func (s WordsView) Name() string {
	return "WordsView"
}
