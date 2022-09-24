package core

type Message struct {
	ID   string
	From User
	To   User
	Msg  string
}
