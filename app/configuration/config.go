package configuration

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func ParseConfig(configPath string) (Config, error) {
	var config Config

	// read config file
	configFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		return config, fmt.Errorf("failed to read config: %w", err)
	}

	// parse config file
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		return config, fmt.Errorf("failed to parse config: %w", err)
	}

	return config, nil
}

type Config struct {
	Monitors     map[MonitorName]Monitor         `yaml:"monitors"`
	Modes        map[ModeName]Mode               `yaml:"modes"`
	Touchscreens map[TouchscreenName]Touchscreen `yaml:"touchscreens"`
	Layouts      []Layout                        `yaml:"layout"`
}

type MonitorName string

type Monitor struct {
	Edid        string          `yaml:"edid"`
	Mode        ModeName        `yaml:"mode"`
	Touchscreen TouchscreenName `yaml:"touchscreen"`
}

type ModeName string

type Mode struct {
	Width  int    `yaml:"width"`
	Height int    `yaml:"height"`
	Line   string `yaml:"line"`
}

type TouchscreenName string

type Touchscreen struct {
}

type Layout struct {
	Skip     bool     `yaml:"skip"`
	Criteria Criteria `yaml:"criteria"`
}

type Criteria struct {
	Monitors map[MonitorName]struct{} `yaml:"monitors"`
}

type LayoutPrimaryMonitor struct {
	Monitor  MonitorName         `yaml:"monitor"`
	Commands EnvironmentCommands `yaml:"commands"`
}

type LayoutAuxiliaryMonitor struct {
	Monitor  MonitorName         `yaml:"monitor"`
	Position LayoutPosition      `yaml:"position"`
	Commands EnvironmentCommands `yaml:"commands"`
}

type LayoutDefaultMonitor struct {
	Position LayoutPosition      `yaml:"position"`
	Commands EnvironmentCommands `yaml:"commands"`
}

type LayoutPosition struct {
	X string `yaml:"x"`
	Y string `yaml:"y"`
}

type EnvironmentCommands struct {
}
