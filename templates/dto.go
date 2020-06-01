package templates

// DTOTemplate is the blueprint for data transfer objects of the API
var DTOTemplate string = `package {{.Reference}}

import "time"

{{range $model := .Models}}
type {{$model.Name}}DTO struct {
	ID    uint   ` + "`json:" + `"id,string,omitempty"` + "`" + `
}
{{end}}`
