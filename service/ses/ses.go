package ses

import "github.com/BiswajitThakur/toy-ses/session"

type Destination struct {
	CcAddresses []string
	ToAddresses []string
}

type Message struct {
	Subject string
	Body    string
}

type SendEmailInput struct {
	Destination Destination
	Message     Message
	// sender email addr
	Source string
}

func New(sess session.Session) *Client {
	// TODO
	return &Client{}
}
