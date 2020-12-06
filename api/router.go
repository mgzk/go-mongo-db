package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func Router() *gin.Engine {
	router := gin.Default()

	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	router.GET("/peoples", handleGetAll)
	router.GET("/peoples/:id", handleGet)
	router.POST("/peoples", handlePost)
	router.PUT("/peoples/:id", handlePut)
	router.DELETE("/peoples/:id", handleDelete)

	return router
}
