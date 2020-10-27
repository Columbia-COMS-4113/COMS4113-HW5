package state

type Node interface {

	Equals(other Node) bool

	MessageHandler(message Message)

	Attribute() interface{}

}
