package state

import (
	"eng-bot/state/events"
)

type Idle struct {
}

func (state Idle) Handle(_ *context, event string, data interface{}) State {
	switch event {
	case events.AddWord:
		return AddingWord{}
	default:
		return state
	}
}

func (s Idle) Name() string {
	return "Idle"
}
