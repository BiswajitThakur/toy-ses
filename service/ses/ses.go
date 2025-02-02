package ses

import "github.com/BiswajitThakur/toy-ses/session"

type Destination struct {
	CcAddresses []string
	ToAddresses []string
}

type Message struct {
	Subject string `json:"subject" binding:"required"`
	Body    string `json:"body" binding:"required"`
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
