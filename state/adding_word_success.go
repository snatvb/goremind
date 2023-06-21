package state

import (
	"fmt"

	tu "github.com/mymmrac/telego/telegoutil"
)

type AddingWordSuccess struct {
}

func (state AddingWordSuccess) Handle(_ *Context, event string, data interface{}) State {
	return state
}

func (state AddingWordSuccess) OnEnter(fsm *FSM, context *Context, from State) {
	println("AddingWordSuccess")
	if addingTranslation, ok := from.(AddingTranslation); ok {
		context.Bot.SendMessage(tu.Message(tu.ID(context.ChatId), fmt.Sprintf("Word %s added!", addingTranslation.word)))
	}
	fsm.Handle("Reset", nil)
}

func (s AddingWordSuccess) Name() string {
	return "AddingWordSuccess"
}
