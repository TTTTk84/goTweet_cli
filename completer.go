package main

import (
	"strings"

	"github.com/c-bata/go-prompt"
)


func completer(d prompt.Document) []prompt.Suggest{
		if d.TextBeforeCursor() == "" {
			return []prompt.Suggest{}
		}
		//args := strings.Split(d.TextBeforeCursor(), " ")
		w := d.GetWordBeforeCursor()

		if strings.HasPrefix(w, "-"){
			return mainOptionCompleter()
		}

		return []prompt.Suggest{}
}

func mainOptionCompleter() []prompt.Suggest{
		s := []prompt.Suggest{
			{Text: "-post", Description: "-post"},
			{Text: "-search", Description: "-search"},
		}

		return s
}
