/*
# ------------------------------------------------------------
# -- downloadsRouter.go
# --
# -- Huang Minghe
# -- 2025-11-5
# ------------------------------------------------------------
*/

package routers

import (
	"omciAnalyzer/controller"

	"github.com/gin-gonic/gin"
)

// router configure
func DownloadsRouterInit(router *gin.Engine) {
	downloadsRouter(router)
}

// user router
func downloadsRouter(engine *gin.Engine) {
	var group = engine.Group("/api/downloads")
	{
		con := &controller.AppController{}
		group.GET("/list", con.DownloadsList)
		group.GET("/download", con.DownloadsDownload)
		group.POST("/zip", con.DownloadsZip)
		group.GET("/preview", con.DownloadsDownload)
	}
}
