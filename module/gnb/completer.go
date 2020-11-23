package gnb

import (
	"strings"

	"github.com/c-bata/go-prompt"
)

var GNBSuggestion = []prompt.Suggest{
	{Text: "ue", Description: "Manage registration and deregistration of UEs"},
	{Text: "pdu-session", Description: "Manage PDU sessions"},
	{Text: "qos", Description: "Apply DSCP PHB to PDU sessions"},
	{Text: "exit", Description: "Exit the gNB module"},
}

var PHBSuggestion = []prompt.Suggest{
	{Text: "ue", Description: "Manage registration and deregistration of UEs"},
	{Text: "pdu", Description: "Manage PDU sessions"},
	{Text: "qos", Description: "Apply DSCP PHB to PDU sessions"},
	{Text: "exit", Description: "Exit the gNB module"},
}

var supiSuggestion = &[]prompt.Suggest{}
var plmnSuggestion = &[]prompt.Suggest{}
var snssaiSuggestion = &[]prompt.Suggest{}

func completerPDU(in prompt.Document) []prompt.Suggest {
	return nil
}

func completerQOS(in prompt.Document) []prompt.Suggest {
	return nil
}

func completerUE(in prompt.Document) []prompt.Suggest {
	return nil
}

func CompleterGNB(in prompt.Document) []prompt.Suggest {
	a := in.TextBeforeCursor()
	var split = strings.Split(a, " ")
	w := in.GetWordBeforeCursor()
	if len(split) > 1 {
		var v = split[0]
		if v == "pdu" {
			return completerPDU(in)
		}
		if v == "qos" {
			return completerQOS(in)
		}
		if v == "ue" {
			return completerUE(in)
		}
		return prompt.FilterHasPrefix(GNBSuggestion, v, true)
	}
	return prompt.FilterHasPrefix(GNBSuggestion, w, true)
}
