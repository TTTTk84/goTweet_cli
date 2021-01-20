package main

import (
	"github.com/c-bata/go-prompt"
)


func main() {
	p := prompt.New(
		executor,
		completer,
		prompt.OptionTitle("goTweet_cli"),
		prompt.OptionPrefix("goTweet_cli >"),
	)

	p.Run()
}
