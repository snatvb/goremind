package keyboard

import (
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
)

const (
	AddWord    string = "AddWord"
	RemoveWord        = "RemoveWord"
)

func CreateWords() *telego.InlineKeyboardMarkup {
	return tu.InlineKeyboard(
		tu.InlineKeyboardRow(
			tu.InlineKeyboardButton("Add word").WithCallbackData(AddWord),
			tu.InlineKeyboardButton("Remove word").WithCallbackData(RemoveWord),
		))
}

// keyboard := tu.Keyboard(
// 	tu.KeyboardRow( // Row 1
// 		// Column 1
// 		tu.KeyboardButton("Button"),

// 		// Column 2, `with` method
// 		tu.KeyboardButton("Poll Regular").
// 			WithRequestPoll(tu.PollTypeRegular()),
// 	),
// 	tu.KeyboardRow( // Row 2
// 		// Column 1, `with` method
// 		tu.KeyboardButton("Contact").WithRequestContact(),

// 		// Column 2, `with` method
// 		tu.KeyboardButton("Location").WithRequestLocation(),
// 	),
// ).WithResizeKeyboard().WithInputFieldPlaceholder("Select something")
