/*
# ------------------------------------------------------------
# -- init.go
# --
# -- Huang Minghe
# -- 2025-11-5
# ------------------------------------------------------------
*/

package routers

import "github.com/gin-gonic/gin"

func RouterInit(router *gin.Engine) {
	OmcianalyzerRouterInit(router)
	DownloadsRouterInit(router)
	LibraryRouterInit(router)
	SequenceTracerRouterInit(router)
	CollectorRouterInit(router)
}
