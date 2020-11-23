package config

import (
	"github.com/pelletier/go-toml"
	"io/ioutil"
)

type PipelineStep struct {
	Exec []string
}

type Rule struct {
	Name     string
	Action   string
	Globs    []string
	Pipeline []PipelineStep
}

type Config struct {
	actionmap map[string][]*Rule
	Rules     []Rule
}

func (cfg *Config) RulesForAction(action string) ([]*Rule, bool) {
	rules, ok := cfg.actionmap[action]
	return rules, ok
}

func ReadConfigFile(path string) (*Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return NewConfig(data)
}

func NewConfig(data []byte) (*Config, error) {
	cfg := &Config{}
	err := toml.Unmarshal(data, cfg)
	if err != nil {
		return nil, err
	}
	cfg.actionmap = make(map[string][]*Rule)
	for i, rule := range cfg.Rules {
		cfg.actionmap[rule.Action] = append(cfg.actionmap[rule.Action], &cfg.Rules[i])
	}
	return cfg, nil
}
