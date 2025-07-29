package configurator

import (
	"encoding/json"
	"log"
)

type Configurator interface {
	GetConfig(key string) (value string, found bool)
}

type ConfigRequest struct {
	Key string // 키 값
}

type ConfigResponse struct {
	Value string // 설정 값
	Found bool   // 해당 설정 존재 여부
}

// GetConfig response Key - request - value found
func GetConfig(key string) (string, bool) {
	cfg, err := validateKEY(key)
	if err != nil {
		log.Println(err)
		return "", false
	}

	resBytes, _ := json.Marshal(cfg)
	return string(resBytes), true
}
