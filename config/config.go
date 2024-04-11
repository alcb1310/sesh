package config

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"github.com/joshmedeski/sesh/dir"
	"github.com/pelletier/go-toml/v2"
)

type (
	DefaultSessionConfig struct {
		StartupScript  string `toml:"startup_script"`
		StartupCommand string `toml:"startup_command"`
		Tmuxp          string `toml:"tmuxp"`
		Tmuxinator     string `toml:"tmuxinator"`
	}

	SessionConfig struct {
		Name string `toml:"name"`
		Path string `toml:"path"`
		DefaultSessionConfig
	}

	Icons struct {
		Tmux   string `toml:"tmux"`
		Config string `toml:"config"`
		Zoxide string `toml:"zoxide"`
	}

	Config struct {
		DefaultSessionConfig DefaultSessionConfig `toml:"default_session"`
		Icons                Icons                `toml:"icons"`
		SessionConfigs       []SessionConfig      `toml:"session"`
		ImportPaths          []string             `toml:"import"`
	}

	ConfigDirectoryFetcher interface {
		GetUserConfigDir() (string, error)
	}

	DefaultConfigDirectoryFetcher struct{}
)

var _ ConfigDirectoryFetcher = (*DefaultConfigDirectoryFetcher)(nil)

func (d *DefaultConfigDirectoryFetcher) GetUserConfigDir() (string, error) {
	switch runtime.GOOS {
	case "darwin":
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		return path.Join(homeDir, ".config"), nil
	default:
		return os.UserConfigDir()
	}
}

func parseConfigFromFile(configPath string, config *Config) error {
	file, err := os.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("error reading config file %s: %v", configPath, err)
	}
	err = toml.Unmarshal(file, config)
	if err != nil {
		return fmt.Errorf("error parsing toml %s: %v", file, err)
	}

	for i, session := range config.SessionConfigs {
		fullPath := dir.FullPath(session.Path)
		config.SessionConfigs[i].Path = fullPath
	}

	if len(config.ImportPaths) > 0 {
		for _, path := range config.ImportPaths {
			importConfig := Config{}
			importConfigPath := dir.FullPath(path)
			if err := parseConfigFromFile(importConfigPath, &importConfig); err != nil {
				return fmt.Errorf("error parsing config from import file path %s: %v", path, err)
			}
			config.SessionConfigs = append(config.SessionConfigs, importConfig.SessionConfigs...)
		}
	}
	return nil
}

// TODO: add error handling (return error)
func ParseConfigFile(fetcher ConfigDirectoryFetcher) Config {
	config := Config{}
	configDir, err := fetcher.GetUserConfigDir()
	if err != nil {
		return config
	}
	configPath := filepath.Join(configDir, "sesh", "sesh.toml")
	parseConfigFromFile(configPath, &config)
	return config
}
