package keyboard

import (
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

const (
	AddWord    string = "AddWord"
	RemoveWord        = "RemoveWord"
	WordList          = "WordList"
	Back              = "Back"
	Forgot            = "Forgot"
)

func WordsControls() *telego.InlineKeyboardMarkup {
	return tu.InlineKeyboard(
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton("Add word").WithCallbackData(AddWord),
			tu.InlineKeyboardButton("Remove word").WithCallbackData(RemoveWord),
		), tu.InlineKeyboardRow(
			tu.InlineKeyboardButton("Words").WithCallbackData(WordList),
		))
}

func BackButton() telego.InlineKeyboardButton {
	return tu.InlineKeyboardButton("Back").WithCallbackData(Back)
}

func BackOnly() *telego.InlineKeyboardMarkup {
	return tu.InlineKeyboard(
		tu.InlineKeyboardRow(
			BackButton(),
		))
}

func ForgotButton() telego.InlineKeyboardButton {
	return tu.InlineKeyboardButton("Forgot").WithCallbackData(Forgot)
}

func ForgotOnly() *telego.InlineKeyboardMarkup {
	return tu.InlineKeyboard(
		tu.InlineKeyboardRow(
			ForgotButton(),
		))
}
