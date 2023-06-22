package state

import (
	"fmt"
	"goreminder/state/events"
	"goreminder/store"

	"github.com/mymmrac/telego"
)

type Context struct {
	Store  *store.Store
	Bot    *telego.Bot
	ChatId int64
}

type State interface {
	Handle(context *Context, event string, data interface{}) State
	Name() string
}

type EnterableState interface {
	OnEnter(fsm *FSM, context *Context, from State)
}

type FSM struct {
	context *Context
	state   State
}

func NewFSM(context *Context) *FSM {
	fmt.Printf("New FSM %d\n", context.ChatId)
	return &FSM{
		context: context,
		state:   Idle{},
	}
}

func (fsm *FSM) Handle(event string, data interface{}) {
	beforeUpdate := fsm.state

	if event == events.Reset {
		fsm.state = Idle{}
	} else {
		fsm.state = fsm.state.Handle(fsm.context, event, data)
	}

	if beforeUpdate != fsm.state {
		if enterableState, ok := fsm.state.(EnterableState); ok {
			fmt.Printf("Call enter: %+v\n", fsm.state.Name())
			enterableState.OnEnter(fsm, fsm.context, beforeUpdate)
		}
	}

	fmt.Printf("After handle state: %+v\n", fsm.state.Name())
}

func (fsm *FSM) CurrentState() State {
	return fsm.state
}
