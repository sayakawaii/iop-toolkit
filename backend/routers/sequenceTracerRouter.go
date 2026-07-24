/*
# ------------------------------------------------------------
# -- sequenceTracerRouter.go
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
func SequenceTracerRouterInit(router *gin.Engine) {
	sequenceTracerRouter(router)
}

// user router
func sequenceTracerRouter(engine *gin.Engine) {
	var group = engine.Group("/api/sequencetracer")
	{
		con := &controller.AppController{}
		group.POST("/request", con.SequenceTracerRequest)
	}
}
