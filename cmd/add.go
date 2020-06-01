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
	"path"
	"strings"

	"github.com/genesis/genesis/generables"
	"github.com/genesis/genesis/utils"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a new service to your API",
	Long: `This command generates the following file structure in a directory named after your service name.
service-name:
|- model.go
|- dto.go
|- mapper.go
|- repository.go
|- service.go
|- api.go
|- route.go
|- bootstrap.go
	`,
	Run: func(cmd *cobra.Command, args []string) {
		serviceName, err := cmd.Flags().GetString("service")

		if err != nil {
			fmt.Println("Error:", err)
		}

		serviceName = strings.Title(serviceName)

		fmt.Printf("Generating service: %s\n", serviceName)

		models, err := cmd.Flags().GetStringSlice("models")

		if err != nil {
			fmt.Println("Error:", err)
		}

		servicePath, err := generateService(serviceName, models)

		if err != nil {
			fmt.Println("Error:", err)
		}

		fmt.Printf("Your service is ready at %s!\n", servicePath)

	},
}

func init() {
	addCmd.Flags().StringP("service", "s", "", "Name of the Service to add to the API")
	addCmd.MarkFlagRequired("service")
	addCmd.Flags().StringSliceP("models", "m", []string{}, "List of model names as a comma separated string. Model names will be converted to title case")

	rootCmd.AddCommand(addCmd)
}

func generateService(serviceName string, models []string) (string, error) {
	wd, err := os.Getwd()

	if err != nil {
		return "", err
	}

	service := generables.Service{
		Name:      serviceName,
		Reference: utils.FirstLower(serviceName),
	}
	service.AbsolutePath = path.Join(wd, service.Reference)

	for _, modelName := range models {
		model := generables.Model{
			Name:      strings.Title(modelName),
			Reference: utils.FirstLower(modelName),
		}
		service.Models = append(service.Models, model)
	}

	err = service.Generate()

	if err != nil {
		return "", err
	}
	return service.AbsolutePath, nil
}
