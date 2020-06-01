package templates

// RouteTemplate is the blueprint for the routes of the API
var RouteTemplate string = `package {{.Reference}}

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func BindRoutes(db *gorm.DB, router *gin.Engine) {
	{{.Reference}}API := Init{{.Name}}API(db)
	{{.Reference}}Routes := router.Group("/{{.Reference}}")
	{{range $model := .Models}}
	{{$.Reference}}Routes.GET("/{{$model.Reference}}s", {{$.Reference}}API.FindAll{{$model.Name}})
	{{$.Reference}}Routes.GET("/{{$model.Reference}}s/:id", {{$.Reference}}API.Find{{$model.Name}}ByID)
	{{$.Reference}}Routes.POST("/{{$model.Reference}}s", {{$.Reference}}API.Create{{$model.Name}})
	{{$.Reference}}Routes.PUT("/{{$model.Reference}}s/:id", {{$.Reference}}API.Update{{$model.Name}})
	{{$.Reference}}Routes.DELETE("/{{$model.Reference}}s/:id", {{$.Reference}}API.Delete{{$model.Name}})
	{{end}}
}`
