package state

import (
	"goreminder/utils"
	"log"

	tu "github.com/mymmrac/telego/telegoutil"
)

type AddingWordSuccess struct {
	Translations []string
}

func (state AddingWordSuccess) Handle(_ *Context, event string, data interface{}) State {
	return state
}

func (state AddingWordSuccess) OnEnter(fsm *FSM, context *Context, from State) {
	log.Printf("AddingWordSuccess")
	if addingTranslation, ok := from.(AddingTranslation); ok {
		context.Bot.SendMessage(tu.MessageWithEntities(
			tu.ID(context.ChatId),
			utils.Prepend(
				utils.StringsToEntities(state.Translations),
				tu.Entity("Word "),
				tu.Entity(addingTranslation.word).Code(),
				tu.Entity(" added!"),
				tu.Entity(".\nTranslation is: "),
			)...,
		))
	} else {
		context.Bot.SendMessage(tu.MessageWithEntities(
			tu.ID(context.ChatId),
			tu.Entity("Error: ").Bold(),
			tu.Entity("Adding translation failed in last part, but word has been added."),
		))
	}
	fsm.Handle("Reset", nil)
}

func (s AddingWordSuccess) Name() string {
	return "AddingWordSuccess"
}
