package executor

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
		var applicable = false
		for _, glob := range rule.Globs {
			ok, _ := filepath.Match(glob, filepath.Base(event.Name))
			if ok {
				applicable = true
				break
			}
		}
		if applicable {
			toApply = append(toApply, rule)
		}
	}
	return toApply
}

func Apply(rule *config.Rule, op, file string) ([]byte, error) {
	if op != rule.Action {
		return nil, fmt.Errorf(
			"action mismatch. Can't apply rule for %s to operation %s on file %s",
			rule.Action,
			op,
			file,
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
				args[i] = strings.Replace(arg, "$file", file, -1)
			}

			cmd := exec.Command(prog, args...)
			out, err := cmd.CombinedOutput()
			output = append(output, out...)
			if err != nil {
				return output, err
			}
		}
	}
	return output, nil
}
