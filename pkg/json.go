package pkg

import (
	"encoding/json"
)

// RenderJsonIndent pretty renders a JSON configuration.
func RenderJsonIndent(cfg interface{}) (string, error) {
	b, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return "", err
	}
	return string(b), nil
}
