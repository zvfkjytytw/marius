package mariuscommon

import (
	"fmt"
	"os"
	"path/filepath"
)

func ReadConfigFile(configFile string) ([]byte, error) {
	filename, err := filepath.Abs(configFile)
	if err != nil {
		return nil, fmt.Errorf("file %s not found: %w", configFile, err) //nolint // wraped higher
	}

	yamlConfig, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("fail read file %s: %w", configFile, err) //nolint // wraped higher
	}

	return yamlConfig, nil
}
