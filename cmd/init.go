/*Package cmd is where the genesis commands reside

Copyright © 2020 Faysal Ishtiaq <f.i.rabby@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/genesis/genesis/generables"
	"github.com/genesis/genesis/utils"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a REST API Project",
	Long:  `Initializes a REST API project using gin-gonic as http handler and gorm as orm.`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		var appName, dbEngine string
		appName, err = cmd.Flags().GetString("app-name")
		if err != nil {
			fmt.Println("Error:", err)
		}

		dbEngine, err = cmd.Flags().GetString("db-engine")
		if err != nil {
			fmt.Println("Error:", err)
		}

		acceptedDbEngines := []string{"mysql", "pgsql", "sqlite", "mssql"}
		if !utils.Contains(acceptedDbEngines, dbEngine) {
			fmt.Printf("Error: %v for --db-engine (-d) value is not supported. \n", dbEngine)
			fmt.Println("Accepted values are:", strings.Join(acceptedDbEngines, ", "))
			os.Exit(1)
		}

		projectPath, err := initializeProject(appName, dbEngine)

		if err != nil {
			fmt.Println("Error: ", err)
		}

		fmt.Printf("Your application is ready at %s!\n", projectPath)
	},
}

func init() {
	initCmd.Flags().StringP("app-name", "n", "", "Package name for your API")
	initCmd.MarkFlagRequired("app-name")
	initCmd.Flags().StringP("db-engine", "e", "mysql", "Database Engine for your API. Accepted values are - mysql, pgsql, mssql, sqlite")

	rootCmd.AddCommand(initCmd)
}

func initializeProject(appName, dbEngine string) (string, error) {
	wd, err := os.Getwd()

	if err != nil {
		return "", err
	}

	app := &generables.Application{
		AbsolutePath: wd,
		Name:         appName,
		DBEngine:     dbEngine,
	}

	if err := app.Create(); err != nil {
		return "", err
	}

	return app.AbsolutePath, nil
}
