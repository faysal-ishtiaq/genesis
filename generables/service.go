package generables

import (
	"fmt"
	"strings"

	"github.com/genesis/genesis/templates"

	"github.com/genesis/genesis/utils"
)

// Service is a part of Application. This struct contains the atributes of a Service
type Service struct {
	AbsolutePath string
	Name         string
	Reference    string
	Models       []Model
}

// Model is data type of your service. A service handle one/multiple models via API or RPC
type Model struct {
	Name      string
	Reference string
}

// Generate generates a service for your application
func (s Service) Generate() error {
	if err := utils.MakeDir(s.AbsolutePath); err != nil {
		return err
	}

	if err := s.GenerateService(); err != nil {
		return err
	}
	if err := s.GenerateAPI(); err != nil {
		return err
	}
	if err := s.GenerateRoute(); err != nil {
		return err
	}
	if err := s.GenerateBootstrap(); err != nil {
		return err
	}
	if err := s.GenerateModel(); err != nil {
		return err
	}
	if err := s.GenerateDTO(); err != nil {
		return err
	}
	if err := s.GenerateMapper(); err != nil {
		return err
	}
	if err := s.GenerateRepository(); err != nil {
		return err
	}

	return nil
}

// GenerateArgsFromModels generate arguments for models for passing to repository factory or for similar usage
func (s Service) GenerateArgsFromModels(entity string) string {
	args := []string{}
	for _, model := range s.Models {
		args = append(args, fmt.Sprintf("%s%s", model.Reference, entity))
	}

	return strings.Join(args, ",")
}

// GenerateParamsFromModels generate arguments for models for passing to repository factory or for similar usage
func (s Service) GenerateParamsFromModels(entity string) string {
	args := []string{}
	for _, model := range s.Models {
		args = append(args, fmt.Sprintf("%s %s%s", model.Reference, model.Name, entity))
	}

	return strings.Join(args, ",")
}

// MultipleModels returns a boolean indicating if the service has multiple models
func (s Service) MultipleModels() bool {
	return len(s.Models) > 1
}

// GenerateService creates service file for a service
func (s Service) GenerateService() error {
	return utils.GenerateFromTemplate(
		s.AbsolutePath,
		"service",
		templates.ServiceTemplate,
		s)
}

// GenerateAPI creates common API endpoints for a service
func (s Service) GenerateAPI() error {
	return utils.GenerateFromTemplate(
		s.AbsolutePath,
		"api",
		templates.APITemplate,
		s)
}

// GenerateRoute creates routes for API endpoints of a service
func (s Service) GenerateRoute() error {
	return utils.GenerateFromTemplate(
		s.AbsolutePath,
		"route",
		templates.RouteTemplate,
		s)
}

// GenerateBootstrap creates bootstrap file for a service
func (s Service) GenerateBootstrap() error {
	return utils.GenerateFromTemplate(
		s.AbsolutePath,
		"bootstrap",
		templates.BootstrapTemplate,
		s)
}

// GenerateModel creates model file for a service
func (s Service) GenerateModel() error {
	return utils.GenerateFromTemplate(
		s.AbsolutePath,
		"model",
		templates.ModelTemplate,
		s)
}

// GenerateDTO creates dto file for a service
func (s Service) GenerateDTO() error {
	return utils.GenerateFromTemplate(
		s.AbsolutePath,
		"dto",
		templates.DTOTemplate,
		s)
}

// GenerateMapper creates dto file for a service
func (s Service) GenerateMapper() error {
	return utils.GenerateFromTemplate(
		s.AbsolutePath,
		"mapper",
		templates.MapperTemplate,
		s)
}

// GenerateRepository creates repository file for a service
func (s Service) GenerateRepository() error {
	return utils.GenerateFromTemplate(
		s.AbsolutePath,
		"repository",
		templates.RepositoryTemplate,
		s)
}
