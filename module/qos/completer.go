package qos

import (
	"github.com/c-bata/go-prompt"
)

// list
// add
// delete

var QOSSuggestion = []prompt.Suggest{
	{Text: "add", Description: "Mark packet with DSCP based on packet match"},
	{Text: "delete", Description: "Delete a match"},
	{Text: "list", Description: "List all rules"},
	{Text: "exit", Description: "Exit the QoS module"},
}

// PHBSuggestion list all the PHB defined by RFC 2597, RFC 2598, RFC 3246,
var PHBSuggestion = []prompt.Suggest{
	{Text: "cs1", Description: "Apply CS1 PHB with DSCP value 001000"},
	{Text: "cs2", Description: "Apply CS2 PHB with DSCP value 010000"},
	{Text: "cs3", Description: "Apply CS3 PHB with DSCP value 011000"},
	{Text: "cs4", Description: "Apply CS4 PHB with DSCP value 100000"},
	{Text: "cs5", Description: "Apply CS5 PHB with DSCP value 101000"},
	{Text: "cs6", Description: "Apply CS6 PHB with DSCP value 110000"},
	{Text: "cs7", Description: "Apply CS7 PHB with DSCP value 111000"},

	{Text: "af11", Description: "Apply AF11 PHB with DSCP value 001010"},
	{Text: "af12", Description: "Apply AF12 PHB with DSCP value 001100"},
	{Text: "af13", Description: "Apply AF13 PHB with DSCP value 001110"},
	{Text: "af21", Description: "Apply AF21 PHB with DSCP value 010010"},
	{Text: "af22", Description: "Apply AF22 PHB with DSCP value 010100"},
	{Text: "af23", Description: "Apply AF23 PHB with DSCP value 010110"},
	{Text: "af31", Description: "Apply AF31 PHB with DSCP value 011010"},
	{Text: "af32", Description: "Apply AF32 PHB with DSCP value 011100"},
	{Text: "af33", Description: "Apply AF33 PHB with DSCP value 011110"},
	{Text: "af41", Description: "Apply AF41 PHB with DSCP value 100010"},
	{Text: "af42", Description: "Apply AF42 PHB with DSCP value 100100"},
	{Text: "af43", Description: "Apply AF43 PHB with DSCP value 100110"},

	{Text: "be", Description: "Apply GetWordBeforeCursor PHB with DSCP value 000000"},

	{Text: "ef", Description: "Apply EF with DSCP value 101110"},
}

func CompleterQOS(in prompt.Document) []prompt.Suggest {
	w := in.GetWordBeforeCursor()
	return prompt.FilterHasPrefix(QOSSuggestion, w, true)

}
