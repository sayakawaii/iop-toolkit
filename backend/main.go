package main

import (
	"fmt"
	"omciAnalyzer/controller"
	"omciAnalyzer/dao"
	"omciAnalyzer/global"

	// "omciAnalyzer/models"
	"omciAnalyzer/routers"
	// "omciAnalyzer/service/taskDistribute"
	"omciAnalyzer/utils"

	"github.com/gin-contrib/sessions"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	//set run mode
	gin.SetMode(global.AppConf.Server.RunMode)

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://127.0.0.1:5173"}
	// config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true
	config.MaxAge = 60 * 60 * 24

	r := gin.Default()
	//disable cache, should only used in debug mode
	// r.Use(middleware.NoCache())
	store := cookie.NewStore([]byte("secret111"))
	fmt.Println("store:", store)
	r.Use(sessions.Sessions("userid", store))

	//set router
	r.Use(cors.New(config))
	routers.RouterInit(r)
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"error": "not found"})
	})
	// The SPA is served by the frontend's nginx, which also reverse-proxies
	// /api and /uploads here. This process used to serve a second copy of the
	// built SPA as well, which meant every frontend change had to be rebuilt
	// into backend/static by hand -- and when that was forgotten, the two
	// entry points silently disagreed about what the UI was.
	r.Static("/uploads", "./static/uploads")

	r.Run(":" + global.AppConf.Server.HttpPort)
}

func init() {
	yamlFile := `./config/config.yaml`
	err := global.InitConf(yamlFile)
	if err != nil {
		utils.Log("InitConf failed")
	}
	err = global.InitMysql()
	if err != nil {
		utils.Log("InitMysql failed")
	} else {
		// err = dao.MysqlAutoMigrate()
		// if err != nil {
		// 	utils.Log("MysqlAutoMigrate failed: ", err)
		// }
		dao.MysqlStartupRepair()
	}
	err = global.InitKafka()
	if err != nil {
		utils.Log("InitKafka failed")
	} else {
		controller.InitConsumers()
	}
}
