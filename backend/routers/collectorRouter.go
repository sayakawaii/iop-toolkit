/*
# ------------------------------------------------------------
# -- collectorRouter.go
# --
# -- Huang Minghe
# -- 2025-10-21
# ------------------------------------------------------------
*/

package routers

import (
	"omciAnalyzer/controller"

	"github.com/gin-gonic/gin"
)

// router configure
func CollectorRouterInit(router *gin.Engine) {
	collectorRouter(router)
}

// user router
func collectorRouter(engine *gin.Engine) {
	var group = engine.Group("/api/collector")
	{
		con := &controller.AppController{}
		group.POST("/connect", con.CollectorConnect)
		group.GET("/query", con.CollectorQuery)
		group.GET("/ws", con.CollectorWebSocket)
		group.GET("/loggerGet", con.CollectorLoggerGet)
		group.POST("/loggerSet", con.CollectorLoggerSet)
		group.GET("/loggerGetLogs", con.CollectorLoggerGetLogs)
	}
}
