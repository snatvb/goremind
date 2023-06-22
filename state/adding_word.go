package state

import (
	"goreminder/keyboard"
	"goreminder/state/events"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

type AddingWord struct {
}

func (state AddingWord) Handle(ctx *Context, event string, data interface{}) State {
	if event == events.Back {
		return Idle{}
	}

	if event == events.Message {
		msg := data.(*telego.Message)
		hasWord := ctx.Store.HasWord(msg.Text, msg.Chat.ID)
		if !hasWord {
			return AddingTranslation{word: msg.Text}
		}
		ctx.Bot.SendMessage(tu.MessageWithEntities(
			tu.ID(msg.Chat.ID),
			tu.Entity("Word "),
			tu.Entity(msg.Text).Code(),
			tu.Entity(" already exists. Please, try another word."),
		).WithReplyMarkup(keyboard.BackOnly()))
		return state
	}

	return state
}

func (state AddingWord) OnEnter(fsm *FSM, context *Context, from State) {
	context.Bot.SendMessage(tu.Message(
		tu.ID(context.ChatId),
		"Write word that you want to add",
	).WithReplyMarkup(keyboard.BackOnly()))
}

func (state AddingWord) Name() string {
	return "AddingWord"
}
