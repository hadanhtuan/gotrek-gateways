package apiSearch

import (
	"github.com/gin-gonic/gin"
	pkg "github.com/hadanhtuan/go-sdk"
)

func InitRoute(router *gin.RouterGroup, app *pkg.App) error {
	searchController := app.Handler[app.Config.GRPC.SearchServicePort].(*SearchController)

	searchGroup := router.Group("/search")

	searchGroup.POST("/property", searchController.Search)
	searchGroup.POST("/property/ip", searchController.SearchByIP)

	return nil
}
