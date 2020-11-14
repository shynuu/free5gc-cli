package factory

import (
	prompt "github.com/c-bata/go-prompt"
)

type Prompt struct {
	Title      string
	Prefix     string
	IsEnable   bool
	Suggestion []prompt.Suggest
}
