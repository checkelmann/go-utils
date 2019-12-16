package utils

import (
	"github.com/keptn/go-utils/pkg/configuration-service/models"
	"gopkg.in/yaml.v2"
)

type SLIConfig struct {
	ServiceLevelObjectives map[string]string `json:"serviceLevelObjectives" yaml:"serviceLevelObjectives"`
}

// GetSLIConfiguration retrieves the SLI configuration for a service, considering SLI configs on stage and project level
func (r *ResourceHandler) GetSLIConfiguration(project string, stage string, service string, resourceURI string) (map[string]string, error) {

	var res *models.Resource
	var err error
	SLIs := make(map[string]string)

	if project != "" {
		res, err = r.GetProjectResource(project, resourceURI)
		if err != nil {
			if err.Error() != "resource not found" {
				return nil, err
			}
		}
		SLIs, _ = addResourceContentToSLIMap(SLIs, res)
	}

	if project != "" && stage != "" {
		res, err = r.GetStageResource(project, stage, resourceURI)
		if err != nil {
			if err.Error() != "resource not found" {
				return nil, err
			}
		}
		SLIs, _ = addResourceContentToSLIMap(SLIs, res)
	}

	if project != "" && stage != "" && service != "" {
		res, err = r.GetServiceResource(project, stage, service, resourceURI)
		if err != nil {
			if err.Error() != "resource not found" {
				return nil, err
			}
		}
		SLIs, _ = addResourceContentToSLIMap(SLIs, res)
	}

	return SLIs, err
}

func addResourceContentToSLIMap(SLIs map[string]string, resource *models.Resource) (map[string]string, error) {
	if resource != nil {
		sliConfig := SLIConfig{}
		err := yaml.Unmarshal([]byte(resource.ResourceContent), &sliConfig)
		if err != nil {
			return nil, err
		}

		for key, value := range sliConfig.ServiceLevelObjectives {
			SLIs[key] = value
		}

	}
	return SLIs, nil
}
