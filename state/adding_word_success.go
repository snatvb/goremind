package state

type AddingWordSuccess struct {
}

func (state AddingWordSuccess) Handle(_ *context, event string, data interface{}) State {
	return state
}

func (s AddingWordSuccess) Name() string {
	return "AddingWordSuccess"
}
