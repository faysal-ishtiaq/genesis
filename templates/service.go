package templates

// ServiceTemplate is the blueprint for the services of the API
var ServiceTemplate string = `package {{.Reference}}

type {{.Name}}Service struct {
	{{range $model := .Models}}
	{{$model.Name}}Repository {{$model.Name}}Repository
	{{end}}
}

func Provide{{.Name}}Service({{.GenerateParamsFromModels "Repository"}}) {{.Name}}Service {
	return {{.Name}}Service{
		{{range $model := .Models}}
		{{$model.Name}}Repository: {{$model.Reference}},
		{{end}}
	}
}

{{range $model := .Models}}
func ({{$.Reference}}Service *{{$.Name}}Service) FindAll{{$model.Name}}() []{{$model.Name}} {
	return {{$.Reference}}Service.{{$model.Name}}Repository.FindAll()
}

func ({{$.Reference}}Service *{{$.Name}}Service) Find{{$model.Name}}ByID(id uint) {{$model.Name}} {
	return {{$.Reference}}Service.{{$model.Name}}Repository.FindByID(id)
}

func ({{$.Reference}}Service *{{$.Name}}Service) Save{{$model.Name}}({{$model.Reference}} {{$model.Name}}) {{$model.Name}} {
	{{$.Reference}}Service.{{$model.Name}}Repository.Save({{$model.Reference}})

	return {{$model.Reference}}
}

func ({{$.Reference}}Service *{{$.Name}}Service) Delete{{$model.Name}}({{$model.Reference}} {{$model.Name}}) {
	{{$.Reference}}Service.{{$model.Name}}Repository.Delete({{$model.Reference}})
}
{{end}}
`
