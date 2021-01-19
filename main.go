package main

import (
	"github.com/c-bata/go-prompt"
)


func main() {
	p := prompt.New(
		executor,
		completer,
		prompt.OptionTitle("gotweet_cli"),
		prompt.OptionPrefix("gotweet_cli >"),
	)

	p.Run()
}
