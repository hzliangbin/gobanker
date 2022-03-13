package healthz

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Check(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}