package main

import (
	_ "log"

	_ "github.com/BiswajitThakur/toy-ses/service/ses"
	_ "github.com/BiswajitThakur/toy-ses/session"
)

const (
	//	Sender           = "sender@example.com"
	Recipient string = "recipient@example.com"
	Subject          = "Demo Email"
)

func main() {
	/*
		sess, err := session.NewSession(&session.Config{
			Region: "us-west-2",
		})

		if err != nil {
			log.Fatal(err)
		}

		svc := ses.New(sess)
	*/
	/*
		input := ses.SendEmailInput{
			Destination: ses.Destination{
				CcAddresses: []string{},
				ToAddresses: []string{Recipient},
			},
			Message: ses.Message{
				Subject: Subject,
				Body:    "Body of the email",
			},
			Source: Sender,
		}
		result, err := svc.SendEmail(input)
		if err != nil {
			log.Fatal(err)
		}*/
	router := Server()
	router.Run(":8080")

	//fmt.Println("Email Sent to address: " + Recipient)
	//fmt.Println(result)
}
