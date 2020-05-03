package rest

import (
	"log"
	"yuting/daily-pg/2020-04-25/handler"
	"yuting/daily-pg/2020-04-25/lib/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	// swagger
	_ "yuting/daily-pg/2020-04-25/docs"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func RunAPI(dbName string, conn string) error {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:  []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposeHeaders: []string{"Content-Length"},
	}))

	h, err := handler.NewHandler(dbName, conn)
	if err != nil {
		log.Panic(err)
	}
	router.GET("/", func(c *gin.Context) {
		c.String(200, "ok")
	})

	routerVersion := router.Group("/v1")
	routerVersion.Use(middleware.Wrapper())
	{
		teamRouter := routerVersion.Group("teams")
		{
			teamRouter.GET("/", h.GetAllTeams)
			teamRouter.POST("/", h.AddOneTeam)
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router.Run()
}
