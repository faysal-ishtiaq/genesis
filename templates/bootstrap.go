package templates

// BootstrapTemplate is the template for bootstrapping services
var BootstrapTemplate string = `package {{.Reference}}

import "github.com/jinzhu/gorm"

// InitDB perform database migrations
func InitDB(db *gorm.DB) {
	{{range $model := .Models}}
	db.AutoMigrate(&{{$model.Name}}{})
	{{end}}
}

// Init{{.Name}}API initializes the API to bind with routes
func Init{{.Name}}API(db *gorm.DB) {{.Name}}API {
	{{range $model := .Models}}
	{{$model.Reference}}Repository := Provide{{$model.Name}}Repository(db)
	{{end}}
	{{.Reference}}Service := Provide{{.Name}}Service({{.GenerateArgsFromModels "Repository"}})
	{{.Reference}}API := Provide{{.Name}}API({{.Reference}}Service)
	return {{.Reference}}API
}
`
