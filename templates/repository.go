package templates

// RepositoryTemplate is the blueprint for service repositories of the API
var RepositoryTemplate string = `package {{.Reference}}

import (
	"github.com/jinzhu/gorm"
)


{{range $model := .Models}}
type {{$model.Name}}Repository struct {
	DB *gorm.DB
}

func Provide{{$model.Name}}Repository(DB *gorm.DB) {{$model.Name}}Repository {
	return {{$model.Name}}Repository{
		DB: DB,
	}
}

func ({{$model.Reference}}Repo *{{$model.Name}}Repository) FindAll() []{{$model.Name}} {
	var {{$model.Reference}}s []{{$model.Name}}
	{{$model.Reference}}Repo.DB.Find(&{{$model.Reference}}s)

	return {{$model.Reference}}s
}

func ({{$model.Reference}}Repo *{{$model.Name}}Repository) FindByID(id uint) {{$model.Name}} {
	var {{$model.Reference}} {{$model.Name}}
	{{$model.Reference}}Repo.DB.First(&{{$model.Reference}}, id)

	return {{$model.Reference}}
}

func ({{$model.Reference}}Repo *{{$model.Name}}Repository) Save({{$model.Reference}} {{$model.Name}}) {{$model.Name}} {
	{{$model.Reference}}Repo.DB.Save(&{{$model.Reference}})

	return {{$model.Reference}}
}

func ({{$model.Reference}}Repo *{{$model.Name}}Repository) Delete({{$model.Reference}} {{$model.Name}}) {
	{{$model.Reference}}Repo.DB.Delete(&{{$model.Reference}})
}
{{end}}
`
