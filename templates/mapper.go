package templates

// MapperTemplate is the blueprint for mappers of DTOs of the API
var MapperTemplate string = `package {{.Reference}}
{{range $model := .Models}}
func To{{$model.Name}}({{$model.Reference}}DTO {{$model.Name}}DTO) {{$model.Name}} {
	return {{$model.Name}}{
	}
}

func To{{$model.Name}}DTO({{$model.Reference}} {{$model.Name}}) {{$model.Name}}DTO {
	return {{$model.Name}}DTO{
	}
}

func To{{$model.Name}}DTOCollection({{$model.Reference}}s []{{$model.Name}}) []{{$model.Name}}DTO {
	{{$model.Reference}}DTOCollection := make([]{{$model.Name}}DTO, len({{$model.Reference}}s))

	for idx, item := range {{$model.Reference}}s {
		{{$model.Reference}}DTOCollection[idx] = To{{$model.Name}}DTO(item)
	}

	return {{$model.Reference}}DTOCollection
}
{{end}}`
