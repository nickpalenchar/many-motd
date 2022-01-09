package data

type Message struct {
	Body   string
	Author string
}

func NewMessage(body string, author string) *Message {
	m := &Message{
		Body:   body,
		Author: author,
	}
	return m
}
