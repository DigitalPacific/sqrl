package pkg

import (
	"fmt"
	"io/ioutil"

	"github.com/olebedev/config"
)

// ParseTextFile reads text from the given filename.
func ParseTextFile(filename string) (*config.Config, error) {
	cfg, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return ParseText(string(cfg))
}

// ParseText reads a text configuration from the given string.
func ParseText(cfg interface{}) (*config.Config, error) {
	parsed, err := parseText([]byte(fmt.Sprintf("%v", cfg)))
	return parsed, err
}

// parseText performs the real text parsing.
func parseText(cfg []byte) (*config.Config, error) {
	return &config.Config{Root: string(cfg)}, nil
}

// RenderText renders a text configuration.
func RenderText(cfg interface{}) (string, error) {
	return fmt.Sprintf("%v", cfg), nil
}
