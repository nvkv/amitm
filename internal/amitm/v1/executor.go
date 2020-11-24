package amitm

import (
	"9fans.net/go/acme"
	"git.sr.ht/~nvkv/amitm/internal/config/v1"

	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
)

func Match(rules []*config.Rule, event acme.LogEvent) []*config.Rule {
	var toApply []*config.Rule

	for _, rule := range rules {
		for _, glob := range rule.Globs {
			ok, _ := filepath.Match(glob, filepath.Base(event.Name))
			if ok {
				toApply = append(toApply, rule)
				break
			}
		}
	}
	return toApply
}

func Apply(rule *config.Rule, event acme.LogEvent) ([]byte, error) {
	if event.Op != rule.Action {
		return nil, fmt.Errorf(
			"action mismatch. Can't apply rule for %s to operation %s on file %s",
			rule.Action,
			event.Op,
			event.Name,
		)
	}

	var output []byte

	for _, step := range rule.Pipeline {
		if len(step.Exec) > 0 {
			prog := step.Exec[0]
			origArgs := step.Exec[1:]
			args := make([]string, len(origArgs))
			copy(args, origArgs)

			for i, arg := range args {
				args[i] = strings.Replace(arg, "$file", event.Name, -1)
			}

			cmd := exec.Command(prog, args...)
			out, err := cmd.CombinedOutput()
			output = append(output, out...)
			if err != nil {
				return output, err
			}
		}
	}
	w, err := acme.Open(event.ID, nil)
	if err != nil {
		return output, err
	}
	_ = w.Ctl("get")
	return output, nil
}
