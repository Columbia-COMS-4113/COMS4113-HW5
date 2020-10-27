package pingpong

import (
	"coms4113/hw5/pkg/state"
)

type Server struct {
	state state.State
	ServerAttribute
}

type ServerAttribute struct {
	counter int
}

func (s *Server) Equals(other state.Node) bool {
	panic("implement me")
}

func (s *Server) MessageHandler(message state.Message) {
	ping, ok := message.(*PingMessage)
	if !ok {
		return
	}

	s.counter++

	pong := &PongMessage{
		from: ping.To(),
		to:   ping.From(),
		Id:   0,
	}

	s.state.Send(pong)
}

func (s *Server) Attribute() interface{} {
	return s.ServerAttribute
}

