package pingpong

import "coms4113/hw5/pkg/state"

type PingMessage struct {
	from state.Address
	to state.Address

	// value
	Id int
}

func (p *PingMessage) From() state.Address {
	return p.from
}

func (p *PingMessage) To() state.Address {
	return p.to
}

type PongMessage struct {
	from state.Address
	to state.Address

	// value
	Id int
}

func (p *PongMessage) From() state.Address {
	return p.from
}

func (p *PongMessage) To() state.Address {
	return p.to
}

