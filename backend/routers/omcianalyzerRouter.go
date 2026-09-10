/*
# ------------------------------------------------------------
# -- omcianalyzerRouter.go
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
func OmcianalyzerRouterInit(router *gin.Engine) {
	omcianalyzerRouter(router)
}

// user router
func omcianalyzerRouter(engine *gin.Engine) {
	var group = engine.Group("/api/omcianalyzer")
	{
		con := &controller.AppController{}
		group.POST("/request", con.OmciAnalyzerRequest)
		group.POST("/minio", con.OmciAnalyzerMinio)
		group.POST("/progress", con.GetOmciAnalyzerProgressData)
		group.POST("/history", con.GetOmciAnalyzerHistoryData)
		group.GET("/diagram", con.GetOmciAnalyzerDiagramData)
		group.GET("/plantuml", con.GetOmciAnalyzerPlantumlData)
		group.GET("/onus", con.GetOmciAnalyzerOnusData)
		group.GET("/omci", con.GetOmciAnalyzerOmciData)
		group.GET("/counters", con.GetOmciAnalyzerCounters)
		group.GET("/yangboards", con.GetOmci2YangBoards)
		group.POST("/generate", con.GenerateOmci2YangConfig)
	}
}
