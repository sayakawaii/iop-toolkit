/*
# ------------------------------------------------------------
# -- libraryRouter.go
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
func LibraryRouterInit(router *gin.Engine) {
	libraryRouter(router)
}

// user router
func libraryRouter(engine *gin.Engine) {
	var group = engine.Group("/api/library")
	{
		con := &controller.AppController{}
		group.POST("/search", con.LibrarySearch)
		group.POST("/addRecord", con.LibraryAddRecord)
		group.POST("/updateRecord", con.LibraryAddRecord)
		group.POST("/deleteRecord", con.LibraryDeleteRecord)
		group.GET("/customers", con.LibraryGetCustomers)
		group.GET("/vendors", con.LibraryGetVendors)
		group.POST("/addCustomer", con.LibraryAddCustomer)
		group.POST("/addVendor", con.LibraryAddVendor)
		group.GET("/download", con.LibraryDownload)
	}
}
