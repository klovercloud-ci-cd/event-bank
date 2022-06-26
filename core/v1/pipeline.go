package v1

import "errors"

// Pipeline Pipeline struct
type Pipeline struct {
	MetaData   PipelineMetadata    `json:"_metadata" yaml:"_metadata"`
	Option     PipelineApplyOption `json:"option" yaml:"option"`
	ApiVersion string              `json:"api_version" yaml:"api_version"`
	Name       string              `json:"name"  yaml:"name"`
	ProcessId  string              `json:"process_id" yaml:"process_id"`
	Label      map[string]string   `json:"label" yaml:"label"`
	Steps      []Step              `json:"steps" yaml:"steps"`
	Claim      int                 `json:"claim" yaml:"claim"`
}

// Validate validates pipeline.
func (pipeline Pipeline) Validate() error {
	if pipeline.ApiVersion == "" {
		return errors.New("api version is required")
	}
	if pipeline.Name == "" {
		return errors.New("pipeline name is required")
	}
	if pipeline.ProcessId == "" {
		return errors.New("pipeline process id is required")
	}

	for _, each := range pipeline.Steps {
		err := each.Validate()
		if err != nil {
			return err
		}
	}
	return nil
}
