package main

import (
	"log"

	"9fans.net/go/acme"
	"git.sr.ht/~nvkv/amitm/internal/config/v1"
	"git.sr.ht/~nvkv/amitm/internal/executor/v1"
)

func main() {
	l, err := acme.Log()
	if err != nil {
		log.Fatal(err)
	}

	config, err := config.ReadConfigFile("./examples/amitm.toml")
	if err != nil {
		log.Fatal(err)
	}

	for {
		event, err := l.Read()
		if err != nil {
			log.Fatal(err)
		}

		rules, ok := config.RulesForAction(event.Op)
		matched := executor.Match(rules, event)
		if ok {
			for _, rule := range matched {
				out, err := executor.Apply(rule, event.Op, event.Name)
				log.Printf("%s: %s\n", rule.Name, string(out))
				if err != nil {
					log.Printf("error: %s\n", err)
				}
			}
		}
	}
}
