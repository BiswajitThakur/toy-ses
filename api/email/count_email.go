package email

import (
	"net/http"

	"github.com/BiswajitThakur/toy-ses/service/ses"
	"github.com/gin-gonic/gin"
)

func CountEmail(ctx *gin.Context) {
	svc := ctx.MustGet("api.email").(*ses.Client)
	data := map[string]interface{}{
		"success_count": svc.TotalSuccessCount,
		"faild_count":   svc.TotalFaildCount,
	}
	ctx.JSONP(http.StatusOK, data)
}
