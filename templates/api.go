package templates

// APITemplate is the blueprint for API of service
var APITemplate string = `package {{.Reference}}

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type {{.Name}}API struct {
	{{.Name}}Service {{.Name}}Service
}

func Provide{{.Name}}API({{.Reference}} {{.Name}}Service) {{.Name}}API {
	return {{.Name}}API{
		{{.Name}}Service: {{.Reference}},
	}
}
{{range $model := .Models}}
func ({{$.Reference}} *{{$.Name}}API) FindAll{{$model.Name}}(c *gin.Context) {
	{{$model.Reference}}s := {{$.Reference}}.{{$.Name}}Service.FindAll{{$model.Name}}()

	c.JSON(http.StatusOK, gin.H{
		"{{$model.Reference}}s": To{{$model.Name}}DTOCollection({{$model.Reference}}s),
	})
}

func ({{$.Reference}} *{{$.Name}}API) Find{{$model.Name}}ByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	{{$model.Reference}} := {{$.Reference}}.{{$.Name}}Service.Find{{$model.Name}}ByID(uint(id))

	if {{$model.Reference}} == ({{$model.Name}}{})  {
		c.Status(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"{{$model.Reference}}": To{{$model.Name}}DTO({{$model.Reference}}),
	})
}

func ({{$.Reference}} *{{$.Name}}API) Create{{$model.Name}}(c *gin.Context) {
	var {{$model.Reference}}DTO {{$model.Name}}DTO
	err := c.ShouldBindJSON(&{{$model.Reference}}DTO)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}

	created{{$model.Name}} := {{$.Reference}}.{{$.Name}}Service.Save{{$model.Name}}(To{{$model.Name}}({{$model.Reference}}DTO))

	c.JSON(http.StatusOK, gin.H{"{{$model.Reference}}": To{{$model.Name}}DTO(created{{$model.Name}})})
}

func ({{$.Reference}} *{{$.Name}}API) Update{{$model.Name}}(c *gin.Context) {
	var {{$model.Reference}}DTO {{$model.Name}}DTO
	err := c.ShouldBindJSON(&{{$model.Reference}}DTO)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	{{$model.Reference}} := {{$.Reference}}.{{$.Name}}Service.Find{{$model.Name}}ByID(uint(id))
	if {{$model.Reference}} == ({{$model.Name}}{}) {
		c.Status(http.StatusBadRequest)
		return
	}

	{{$.Reference}}.{{$.Name}}Service.Save{{$model.Name}}({{$model.Reference}})

	c.Status(http.StatusOK)
}

func ({{$.Reference}} *{{$.Name}}API) Delete{{$model.Name}}(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	{{$model.Reference}} := {{$.Reference}}.{{$.Name}}Service.Find{{$model.Name}}ByID(uint(id))
	if {{$model.Reference}} == ({{$model.Name}}{}) {
		c.Status(http.StatusBadRequest)
		return
	}

	{{$.Reference}}.{{$.Name}}Service.Delete{{$model.Name}}({{$model.Reference}})

	c.Status(http.StatusOK)
}
{{end}}`
