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
