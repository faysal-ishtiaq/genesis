package utils

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

// GenerateFromTemplate generates file in a given directory with given name from given file. Returns error if there is any
func GenerateFromTemplate(directory string, filename string, fileTemplate string, values interface{}) error {
	if !strings.Contains(filename, ".") {
		filename = fmt.Sprintf("%s.go", filename)
	}

	fileToGenerate, err := os.Create(fmt.Sprintf("%s/%s", directory, filename))

	if err != nil {
		return err
	}

	defer fileToGenerate.Close()

	generatedTemplate := template.Must(template.New(filename).Parse(fileTemplate))

	err = generatedTemplate.Execute(fileToGenerate, values)

	return err
}
