package config

import (
	"fmt"
	"testing"
)

func TestReadConfig(t *testing.T) {
	fixture := []byte(`
[[rules]]
name = "Testing tests"
glob = "test*.test"
action = "put"

[[rules.pipeline]]
exec = "test -t $file"`)

	cfg, err := NewConfig(fixture)
	fmt.Printf("cfg: %+v\n", cfg.actionmap)
	if err != nil || cfg == nil {
		t.Errorf("Can't read %s file: %s", fixture, err)
	}
}
