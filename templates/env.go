package templates

// EnvTemplate is the template for initializing project environment variables
var EnvTemplate string = `
DB_DRIVER={{.DBEngine}}
{{if eq .DBEngine "sqlite"}}
DB_PATH=./{{.Name}}.db
{{else}}
DB_NAME={{.Name}}
DB_USERNAME=root
DB_PASSWORD=
DB_HOST=localhost
DB_PORT={{.DefaultDBPort}}
{{end}}
`
