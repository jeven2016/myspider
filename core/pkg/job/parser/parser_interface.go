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
	//TODO
	//Validate(params *ParseParams) // url could be blank
	Parse(params *ParseParams) *ParseResult
}
