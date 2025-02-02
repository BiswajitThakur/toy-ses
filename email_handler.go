package main

import (
	"log"
	"strings"

	"github.com/BiswajitThakur/toy-ses/service/ses"
	"github.com/BiswajitThakur/toy-ses/session"
	"github.com/gin-gonic/gin"
)

const Sender = "sender@example.com"

func EmailHandler() gin.HandlerFunc {

	sess, err := session.NewSession(&session.Config{
		Region: "us-west-2",
	})

	if err != nil {
		log.Fatal(err)
	}

	svc := ses.New(sess)

	return func(ctx *gin.Context) {
		if strings.HasPrefix(ctx.Request.URL.Path, "/api/email") {
			ctx.Set("api.email", svc)
			ctx.Set("api.email.sender", Sender)
		}
		ctx.Next()
	}
}
