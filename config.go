package rockset

// Config contains shared config parameters, common to the source and
// destination. If you don't need shared parameters you can entirely remove this
// file.
type Config struct {
	// Region of the Rockset API to use
	Region     string `json:"region" validate:"inclusion=us-west-2|us-east-1|eu-central-1" default:"us-west-2"`
	APIKey     string `json:"api_key" validate:"required"`
	Workspace  string `json:"workspace" validate:"required"`
	Collection string `json:"collection" validate:"required"`
}
