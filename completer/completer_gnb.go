package completer

import "github.com/c-bata/go-prompt"

var GNBSuggestion = []prompt.Suggest{
	{Text: "pdu", Description: "Manage the PDU sessions established"},
	{Text: "qos", Description: "Manage the QoS of the PDU session"},
	{Text: "exit", Description: "Exit the module"},
}
