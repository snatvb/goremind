package state

import (
	"goreminder/db"
	"goreminder/keyboard"

	tu "github.com/mymmrac/telego/telegoutil"
)

type RemindWord struct {
	Word *db.WordModel
}

func (state RemindWord) Handle(ctx *Context, event string, data interface{}) State {
	return state
}

func (state RemindWord) OnEnter(fsm *FSM, context *Context, from State) {
	context.Bot.SendMessage(tu.MessageWithEntities(
		tu.ID(context.ChatId),
		tu.Entity("Write translation for the word "),
		tu.Entity(state.Word.Word).Code(),
	).WithReplyMarkup(keyboard.ForgotOnly()))
}

func (s RemindWord) Name() string {
	return "RemindWord"
}
