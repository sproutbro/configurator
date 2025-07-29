package internal

type Configurator interface {
	GetConfig(key string) (value string, found bool)
}

type ParserInterface interface {
	Parser(t string) (map[string]interface{}, error)
	JSON(t string) (map[string]interface{}, error)
	ENV(t string) (map[string]interface{}, error)
}
