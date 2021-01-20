package main

import (
	"strings"

	"github.com/c-bata/go-prompt"
)


func completer(d prompt.Document) []prompt.Suggest{
		if d.TextBeforeCursor() == "" {
			return []prompt.Suggest{}
		}
		args := strings.Split(d.TextBeforeCursor(), " ")
		w := d.GetWordBeforeCursor()


		if strings.HasPrefix(w, "-"){
			return mainOptionCompleter(args)
		}

		return []prompt.Suggest{}
}

func mainOptionCompleter(args []string) []prompt.Suggest{
		firstOption := []prompt.Suggest{
			{Text: "-post", Description: "-post"},
			{Text: "-search", Description: "-search"},
		}
		searchOption := []prompt.Suggest{
			{Text: "-page", Description: "-page"},
			{Text: "-count", Description: "-count"},
			{Text: "-include_entities", Description: "-include_entities"},
		}

		if len(args) < 2 {
			return firstOption
		}

		switch args[0]{
		case "-post":
			//
		case "-search":
			return searchOption
		}

		return []prompt.Suggest{}
}
