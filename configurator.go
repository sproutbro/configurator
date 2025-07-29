package configurator

type Configurator interface {
	GetConfig(key string) (value string, found bool)
}

type ConfigResponse struct {
	Value string // 설정 값
	Found bool   // 해당 설정 존재 여부
}

// GetConfig response Key - request - value found
func GetConfig(key string) (string, bool) {
	switch key {
	case "dev":
		return "dev-config", true
	case "prod":
		return "prod-config", true
	default:
		return "", false
	}
}
