package job

type ParseParams struct {
	Url string
}

type Parser interface {
	Parse(params *ParseParams) []string
}
