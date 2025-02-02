package email

import (
	"net/http"

	"github.com/BiswajitThakur/toy-ses/service/ses"
	"github.com/gin-gonic/gin"
)

func SendEmail(ctx *gin.Context) {
	var msg ses.Message
	if ctx.ShouldBindJSON(&msg) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "faild"})
		return
	}
	svc := ctx.MustGet("api.email").(*ses.Client)
	input := ses.SendEmailInput{
		Destination: ses.Destination{
			CcAddresses: []string{},
			ToAddresses: []string{ctx.Param("email")},
		},
		Message: msg,
		Source:  ctx.MustGet("api.email.sender").(string),
	}
	_, err := svc.SendEmail(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "faild"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}
