package generables

import (
	"github.com/genesis/genesis/templates"
	"github.com/genesis/genesis/utils"
	"github.com/spf13/viper"
)

// Application struct represents a project
type Application struct {
	Name         string
	DBEngine     string
	AbsolutePath string
}

// Create creates a project
func (a Application) Create() error {
	if err := a.GenerateEnv(); err != nil {
		return err
	}

	if err := a.GenerateMain(); err != nil {
		return err
	}

	return nil
}

// DefaultDBPort returns the default port for chosen db engine
func (a Application) DefaultDBPort() string {
	switch a.DBEngine {
	case "mysql":
		return "3306"
	case "postgres":
		return "5432"
	case "mssql":
		return "1433"
	}
	return ""
}

// DBURL returns the db url for gorm
func (a Application) DBURL() string {
	switch a.DBEngine {
	case "sqlite":
		return `viper.Get("DB_PATH")`
	case "mysql":
		if viper.Get("DB_HOST") != "" {
			return `fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", viper.Get("DB_USERNAME"), viper.Get("DB_PASSWORD"), viper.Get("DB_HOST"), viper.Get("DB_NAME"))`
		}
		return `fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", viper.Get("DB_USERNAME"), viper.Get("DB_PASSWORD"), viper.Get("DB_NAME"))`
	case "mssql":
		return `fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", viper.Get("DB_USERNAME"), viper.Get("DB_PASSWORD"), viper.Get("DB_HOST"), viper.Get("DB_PORT"), viper.Get("DB_NAME"))`
	case "postgres":
		return `fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", viper.Get("DB_HOST"), viper.Get("DB_PORT"), viper.Get("DB_USERNAME"), viper.Get("DB_NAME"), viper.Get("DB_PASSWORD"))`
	}
	return ""

}

// GenerateEnv generates .env file for the application
func (a Application) GenerateEnv() error {
	return utils.GenerateFromTemplate(
		a.AbsolutePath,
		".env",
		templates.EnvTemplate,
		a)
}

// GenerateMain generates main.go file for the application
func (a Application) GenerateMain() error {
	return utils.GenerateFromTemplate(
		a.AbsolutePath,
		"main",
		templates.MainTemplate,
		a)
}
