package filter

type CLIFilter struct {
	Json         bool `json:"json"`
	Periodically bool `json:"periodically"`
	Second       int  `json:"second"`
}
