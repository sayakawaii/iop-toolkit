package main

import (
	"fmt"
	"omciAnalyzer/controller"
	"omciAnalyzer/dao"
	"omciAnalyzer/global"
	"strings"

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
		if strings.HasPrefix(c.Request.URL.Path, "/api") {
			c.JSON(404, gin.H{"error": "not found"})
			return
		}
		c.File("./static/index.html")
	})
	r.Static("/assets", "./static/assets")
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
