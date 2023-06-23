package state

import (
	"fmt"
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
	words := context.Store.GetWords(context.ChatId, 0, 50)
	var answer []tu.MessageEntityCollection
	for _, word := range words {
		answer = append(answer, tu.Entity("Word: "))
		answer = append(answer, tu.Entity(word.Word).Code())
		answer = append(answer, tu.Entity("; Level: "))
		answer = append(answer, tu.Entity(fmt.Sprintf("%d\n", word.RememberLevel)).Bold())
	}
	context.Bot.SendMessage(tu.MessageWithEntities(
		tu.ID(context.ChatId),
		answer...,
	).WithReplyMarkup(keyboard.BackOnly()))

	// wordsTextsMsg := "Words:\n"
	// for _, word := range words {
	// 	wordsTextsMsg += fmt.Sprintf("%s; Level: %d \n", word.Word, word.RememberLevel)
	// }
	// context.Bot.SendMessage(tu.Message(
	// 	tu.ID(context.ChatId),
	// 	wordsTextsMsg,
	// ).WithReplyMarkup(keyboard.BackOnly()))
}

func (s WordsView) Name() string {
	return "WordsView"
}
