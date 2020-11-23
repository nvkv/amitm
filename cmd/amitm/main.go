package main

import (
	"fmt"
	"log"

	"9fans.net/go/acme"
	"git.sr.ht/~nvkv/amitm/internal/config/v1"
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
		if ok {
			fmt.Printf("%+v -> %+v", event, rules)
		}
	}
}
