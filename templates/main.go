package templates

// MainTemplate is the template for the generated API's main.go
var MainTemplate string = `
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	_ "github.com/jinzhu/gorm/dialects/{{.DBEngine}}"
)

//go:generate go mod init {{.Name}}
//go:generate go get "github.com/gin-gonic/gin"
//go:generate go get "github.com/jinzhu/gorm"
//go:generate go get "github.com/spf13/viper"

func initDB() *gorm.DB {
	db, err := gorm.Open({{if eq .DBEngine "sqlite"}}"sqlite3"{{else}}"{{.DBEngine}}"{{end}}, {{.DBURL}})
	if err != nil {
		panic(err)
	}

	return db
}

func main() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	db := initDB()
	defer db.Close()

	router := gin.Default()

	err := router.Run()
	if err != nil {
		panic(err)
	}
}`
