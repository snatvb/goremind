package state

import (
	"goreminder/config"
	"goreminder/db"
	"goreminder/keyboard"
	"goreminder/state/events"
	"strings"
	"time"

	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	"github.com/steebchen/prisma-client-go/runtime/types"
)

type RemindWord struct {
	Word *db.WordModel
}

func (state RemindWord) Handle(ctx *Context, event string, data interface{}) State {
	if event == events.Forgot {
		state.handleForgot(ctx)
		return Idle{}
	}
	if event == events.Message {
		state.handleTranslation(ctx, data.(*telego.Message))
		return Idle{}
	}
	return state
}

// trim and equal strings
func equalStrings(a, b string) bool {
	return strings.EqualFold(strings.TrimSpace(a), strings.TrimSpace(b))
}

func (state RemindWord) handleTranslation(ctx *Context, msg *telego.Message) {
	translation := state.Word.Translate

	for _, v := range strings.Split(translation, ",") {
		translation = strings.TrimSpace(v)
		if equalStrings(msg.Text, translation) {
			state.handleCorrect(ctx)
			return
		}
	}

	state.handleIncorrect(ctx)
}

func (state RemindWord) handleIncorrect(ctx *Context) {
	rememberIn := config.Timings[0]
	nextRememberAt := time.Now().Add(rememberIn)
	ctx.Bot.SendMessage(tu.MessageWithEntities(
		tu.ID(ctx.ChatId),
		tu.Entity("Incorrect! I remind you in "),
		tu.Entity(rememberIn.String()).Code(),
		tu.Entity(".\n Correct answer: "),
		tu.Entity(state.Word.Translate).Code(),
	))

	res := ctx.Store.UpdateWord(
		state.Word,
		db.Word.RememberLevel.Set(0),
		db.Word.NextRemindAt.Set(types.DateTime(nextRememberAt)),
	)

	if !res {
		ctx.Bot.SendMessage(tu.MessageWithEntities(
			tu.ID(ctx.ChatId),
			tu.Entity("Error: ").Bold(),
			tu.Entity("Save updated word has been failed").Code(),
		))
	}
}

func (state RemindWord) handleCorrect(ctx *Context) {
	nextLevel := state.Word.RememberLevel + 1
	if nextLevel >= len(config.Timings) {
		ctx.Bot.SendMessage(tu.MessageWithEntities(
			tu.ID(ctx.ChatId),
			tu.Entity("🎉 Success! You have remembered the word! 🎊"),
		))

		res := ctx.Store.RemoveWord(state.Word.Word, state.Word.ChatID)
		if !res {
			ctx.Bot.SendMessage(tu.MessageWithEntities(
				tu.ID(ctx.ChatId),
				tu.Entity("Error: ").Bold(),
				tu.Entity("Removing word has been failed").Code(),
			))
		}
	} else {
		rememberIn := config.Timings[nextLevel]
		nextRememberAt := time.Now().Add(rememberIn)
		ctx.Bot.SendMessage(tu.MessageWithEntities(
			tu.ID(ctx.ChatId),
			tu.Entity("Success! I remind you in "),
			tu.Entity(rememberIn.String()).Code(),
			tu.Entity(".\n Correct answer: "),
			tu.Entity(state.Word.Translate).Code(),
		))

		res := ctx.Store.UpdateWord(
			state.Word,
			db.Word.RememberLevel.Set(nextLevel),
			db.Word.NextRemindAt.Set(types.DateTime(nextRememberAt)),
		)

		if !res {
			ctx.Bot.SendMessage(tu.MessageWithEntities(
				tu.ID(ctx.ChatId),
				tu.Entity("Error: ").Bold(),
				tu.Entity("Save updated word has been failed").Code(),
			))
		}
	}
}

func (state RemindWord) handleForgot(ctx *Context) {
	firstLevel := config.Timings[0]
	word := state.Word

	ctx.Bot.SendMessage(tu.MessageWithEntities(
		tu.ID(ctx.ChatId),
		tu.Entity("It's ok. You can try again later.\nCorrect answer: "),
		tu.Entity(state.Word.Translate).Code(),
	))

	res := ctx.Store.UpdateWord(word, db.Word.NextRemindAt.Set(
		types.DateTime(time.Now().Add(firstLevel)),
	))

	if !res {
		ctx.Bot.SendMessage(tu.MessageWithEntities(
			tu.ID(ctx.ChatId),
			tu.Entity("Error: ").Bold(),
			tu.Entity("Save updated word has been failed").Code(),
		))
	}
}

func (state RemindWord) OnEnter(fsm *FSM, context *Context, from State) {
	context.Bot.SendMessage(tu.MessageWithEntities(
		tu.ID(context.ChatId),
		tu.Entity("Write translation for the word "),
		tu.Entity(state.Word.Word).Code(),
		tu.Entity("."),
	).WithReplyMarkup(keyboard.ForgotOnly()))
}

func (s RemindWord) Name() string {
	return "RemindWord"
}
