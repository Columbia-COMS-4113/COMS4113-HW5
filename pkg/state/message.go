package state

type Address string

type Message interface {

	From() Address

	To() Address

}
