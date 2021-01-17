package main

import (
	"github.com/c-bata/go-prompt"
)


func main() {
	p := prompt.New(
		executor,
		completer,
		prompt.OptionTitle("gotweet"),
		prompt.OptionPrefix("gotweet >"),
	)

	p.Run()
}
