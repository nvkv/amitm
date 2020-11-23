package config

import (
	"fmt"
	"testing"
)

func TestReadConfig(t *testing.T) {
	fixture := []byte(`
[[rules]]
name = "Golang"
globs = ["*.go"]
action = "put"

[[rules.pipeline]]
exec = ["go", "fmt", "$file"]

[[rules.pipeline]]
exec = ["echo", "'Done formatting $file'"]

[[rules]]
name = "Terraform"
globs = ["*.tf", "*.tfvars"]
action = "put"

[[rules.pipeline]]
exec = ["terraform13", "fmt", "$file"]`)

	cfg, err := NewConfig(fixture)
	for k, rs := range cfg.actionmap {
		fmt.Printf("%s =>\n", k)
		for _, r := range rs {
			fmt.Printf("\t%s\n", r.Name)
		}
	}
	if err != nil || cfg == nil {
		t.Errorf("Can't read %s file: %s", fixture, err)
	}
}
