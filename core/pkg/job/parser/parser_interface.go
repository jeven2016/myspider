package parser

type ParseParams struct {
	Url     string
	SiteKey string
	Payload any
}

type ParseResult struct {
	Payload any
}

type Parser interface {
	Parse(params *ParseParams) *ParseResult
}
