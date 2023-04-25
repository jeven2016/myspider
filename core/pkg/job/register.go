package job

var parserMap = make(map[string]Parser)

func RegisterParser(parserName string, parser Parser) {
	parserMap[parserName] = parser
}

func GetParser(parserName string) Parser {
	return parserMap[parserName]
}
