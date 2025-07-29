package configurator

import (
	"path/filepath"
	"strings"
)

var configLsit = make(map[string]string)

func validateKEY(key string) (any, error) {

	filename, ok := configLsit[key]
	root, err := findProjectRoot()
	if !ok || err != nil {
		return "", errNOKEYErr
	}

	sp := strings.Split(filename, ".")
	if sp[1] == "env" {
		path := filepath.Join(root, "env", filename)
		return parseEnv(path)
	}

	switch filepath.Ext(filename) {
	case ".json":
		path := filepath.Join(root, "json", filename)
		return parseJSON(path)

	case ".yaml":
		path := filepath.Join(root, "yaml", filename)
		return parseYAML(path)
	}

	return nil, errNOKEYErr
}
