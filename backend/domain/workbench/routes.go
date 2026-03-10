package workbench

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.RouterGroup, domain *Domain) {
	group := router.Group("/workbench")
	{
		group.GET("/overview", domain.ApiOverview)
		group.POST("/actions/:id/transition", domain.ApiTransition)
		group.GET("/actions/:id/events", domain.ApiEvents)
	}
}
