package state

import (
	"goreminder/keyboard"
	"goreminder/state/events"

	tu "github.com/mymmrac/telego/telegoutil"
)

type Idle struct {
}

func (state Idle) Handle(_ *Context, event string, data interface{}) State {
	switch event {
	case events.AddWord:
		return AddingWord{}
	case events.WordList:
		return WordsView{}
	case events.RemoveWord:
		return RemoveWord{}
	default:
		return state
	}
}

func (state Idle) OnEnter(fsm *FSM, context *Context, from State) {
	context.Bot.SendMessage(tu.Message(
		tu.ID(context.ChatId),
		"Choose action",
	).WithReplyMarkup(keyboard.WordsControls()))
}

func (s Idle) Name() string {
	return "Idle"
}
