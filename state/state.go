package state

import (
	"eng-bot/store"
	"fmt"
)

type context struct {
	store *store.Store
}

type State interface {
	Handle(context *context, event string, data interface{}) State
	Name() string
}

type FSM struct {
	context *context
	state   State
}

func NewFSM(store *store.Store) *FSM {
	return &FSM{
		context: &context{
			store: store,
		},
		state: Idle{},
	}
}

func (fsm *FSM) Handle(event string, data interface{}) {
	if event == "Reset" {
		fsm.state = Idle{}
	} else {
		fsm.state = fsm.state.Handle(fsm.context, event, data)
	}
	fmt.Printf("New state: %+v\n", fsm.state.Name())
}

func (fsm *FSM) CurrentState() State {
	return fsm.state
}
