package state

type State struct {
	Servers map[Address]Node

	Clients map[Address]Node

	Network []Message
}

func (state *State) Timer(addr Address) *Timer {
	return nil
}

func (state *State) Send(meg Message) {
	return
}
