package config

import (
	"fmt"
	"testing"
)

func TestReadConfig(t *testing.T) {
	cfg, err := ReadConfig("../../../examples/amitm.toml")
	fmt.Printf("cfg: %+v\n", cfg.actionmap)
	if err != nil {
		t.Errorf("Can't read example.toml file: %s", err)
	}
}
