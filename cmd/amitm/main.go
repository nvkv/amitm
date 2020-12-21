package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"

	"9fans.net/go/acme"
	"github.com/nvkv/amitm/internal/amitm/v1"
	"github.com/nvkv/amitm/internal/config/v1"
)

func main() {
	var cpathptr = flag.String("config", "", `Path to the config file.
If not provided will default to $HOME/.amitm.toml`)

	flag.Parse()
	configPath := *cpathptr

	if len(configPath) == 0 {
		home := os.Getenv("HOME")
		if len(home) == 0 {
			log.Fatal("Amitm can't figure out config path. Consider to provide it using -config")
		}
		configPath = path.Join(home, ".amitm.toml")
	}

	l, err := acme.Log()
	if err != nil {
		log.Fatal(err)
	}

	config, err := config.ReadConfigFile(configPath)
	if err != nil {
		log.Fatal(fmt.Sprintf("Amitm can't read config file: %s", err))
	}

	for {
		event, err := l.Read()
		if err != nil {
			log.Fatal(err)
		}

		rules, ok := config.RulesForAction(event.Op)
		matched := amitm.Match(rules, event)
		if ok {
			for _, rule := range matched {
				out, err := amitm.Apply(rule, event)
				log.Printf("%s:\n%s", rule.Name, string(out))
				if err != nil {
					log.Printf("error: %s\n", err)
				}
			}
		}
	}
}
