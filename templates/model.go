package templates

// ModelTemplate is the blueprint for Models of the API
var ModelTemplate string = `package {{.Reference}}

import "github.com/jinzhu/gorm"

{{range $model := .Models}}
type {{$model.Name}} struct {
	gorm.Model
}
{{end}}`
