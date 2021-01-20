package goTweet_cli

import (
	"strings"

	"github.com/c-bata/go-prompt"
)

type Completer struct{}

func NewCompleter() (*Completer, error){
	return &Completer{}, nil
}


func (c *Completer) Completer(d prompt.Document) []prompt.Suggest{
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
			{Text: "-tweet", Description: "simple tweet"},
			{Text: "-search", Description: "simple search user accounts"},
		}
		searchOption := []prompt.Suggest{
			{Text: "-page", Description: "Specifies the page of results to retrieve."},
			{Text: "-count", Description: "The number of potential user results to retrieve per page."},
			{Text: "-include_entities", Description: "The entities node will not be included in embedded Tweet objects when set to false "},
		}

		if len(args) < 2 {
			return firstOption
		}

		switch args[0]{
		case "-tweet":
			//
		case "-search":
			return searchOption
		}

		return []prompt.Suggest{}
}
