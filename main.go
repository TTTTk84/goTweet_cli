package main

import (
	"log"

	"github.com/TTTTk84/twitter_cli/goTweet_cli"
	"github.com/c-bata/go-prompt"
	"github.com/joho/godotenv"
)

func env_load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("err .env loading ", err)
	}
}

func main() {
  env_load()

	c, err := goTweet_cli.NewCompleter()
	if err != nil {
		log.Fatal(err)
	}

	p := prompt.New(
		c.Executor,
		c.Completer,
		prompt.OptionTitle("goTweet_cli"),
		prompt.OptionPrefix("goTweet_cli >"),
	)

	p.Run()
}
