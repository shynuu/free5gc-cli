package gnb

import (
	"strings"

	"github.com/c-bata/go-prompt"
)

// GNBSuggestion qos add --session 10 --protocol tcp --destination-port 80 --phb
var GNBSuggestion = []prompt.Suggest{
	{Text: "user", Description: "Manage registration and deregistration of UEs"},
	{Text: "pdu-session", Description: "Manage PDU sessions"},
	{Text: "configuration", Description: "Manage the configuration of the gNB module"},
	{Text: "exit", Description: "Exit the gNB module"},
}

var configurationSuggestion = []prompt.Suggest{
	{Text: "reload", Description: "Reload the configuration of the subscriber module"},
}

// ue register --supi imsi-20893XXXXXX00/20893
// ue deregister --supi imsi-20893XXXXXX00/20893
// ue list
var ueSuggestion = []prompt.Suggest{
	{Text: "register", Description: "Register a UE on the network"},
	{Text: "deregister", Description: "Deregister a UE of the network"},
	{Text: "list", Description: "List the registered UE"},
}

var pduSuggestion = []prompt.Suggest{
	{Text: "request", Description: "Establish a new PDU session for a registered UE"},
	{Text: "release", Description: "Release an existing PDU session"},
	{Text: "list", Description: "List the registered UE"},
	{Text: "qos", Description: "Apply DSCP PHB to PDU sessions"},
}

// PHBSuggestion list all the PHB defined by RFC 2597, RFC 2598, RFC 3246,
var PHBSuggestion = []prompt.Suggest{
	{Text: "be", Description: "Apply Best Effort PHB with DSCP value 000000"},

	{Text: "ef", Description: "Apply Expedited Forward PHB with DSCP value 101110"},

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
}

var UserSuggestion = &[]prompt.Suggest{}
var RegisteredSuggestion = &[]prompt.Suggest{}
var SnssaiSuggestion = &[]prompt.Suggest{}
var DNNSuggestion = &[]prompt.Suggest{}
var PDUSuggestion = &[]prompt.Suggest{}

// pdu-session list
// pdu-session request --user imsiXXXXX --snssai 01010101 --dnn internet
// pdu-session
// ===> ipv4, qos profile, sessionid
// pdu-session release --session <session_id>
// pdu-session
func completerPDU(in prompt.Document) []prompt.Suggest {
	a := in.GetWordBeforeCursor()
	a = strings.TrimSpace(a)
	d := strings.Split(in.TextBeforeCursor(), " ")
	l := len(d)

	if d[1] == "list" {
		return prompt.FilterHasPrefix(*SnssaiSuggestion, a, true)
	}

	if d[1] == "request" {
		if l == 3 {
			return prompt.FilterHasPrefix([]prompt.Suggest{
				{Text: "--user", Description: "Specify the user"},
			}, a, true)
		}
		if l == 4 {
			return prompt.FilterHasPrefix(*RegisteredSuggestion, a, true)
		}
		if l == 5 {
			return prompt.FilterHasPrefix([]prompt.Suggest{
				{Text: "--snssai", Description: "Specify the SNSSAI for the PDU Session"},
			}, a, true)
		}
		if l == 6 {
			return prompt.FilterHasPrefix(*SnssaiSuggestion, a, true)
		}
		if l == 7 {
			return prompt.FilterHasPrefix([]prompt.Suggest{
				{Text: "--dnn", Description: "Specify the data network associated to the SNSSAI"},
			}, a, true)
		}
		if l == 8 {
			return prompt.FilterHasPrefix(*DNNSuggestion, a, true)
		}
		if l > 8 {
			return []prompt.Suggest{}
		}
		return []prompt.Suggest{}
	}

	if d[1] == "release" {
		return prompt.FilterHasPrefix(*SnssaiSuggestion, a, true)
	}

	if d[1] == "qos" {
		return prompt.FilterHasPrefix(*SnssaiSuggestion, a, true)
	}

	return prompt.FilterHasPrefix(pduSuggestion, a, true)
}

func completerUE(in prompt.Document) []prompt.Suggest {
	a := in.GetWordBeforeCursor()
	a = strings.TrimSpace(a)
	d := strings.Split(in.TextBeforeCursor(), " ")
	if d[1] == "register" {
		l := len(d)
		if l > 2 && l < 4 {
			return prompt.FilterHasPrefix([]prompt.Suggest{
				{Text: "--user", Description: "Specify the user to register"},
			}, a, true)
		}
		if l > 3 && l < 5 {
			return prompt.FilterHasPrefix(*UserSuggestion, a, true)
		}
	}
	if d[1] == "deregister" {
		l := len(d)
		if l > 2 && l < 4 {
			return prompt.FilterHasPrefix([]prompt.Suggest{
				{Text: "--user", Description: "Specify the user to deregister"},
			}, a, true)
		}
		if l > 3 && l < 5 {
			return prompt.FilterHasPrefix(*RegisteredSuggestion, a, true)
		}
	}
	return prompt.FilterHasPrefix(ueSuggestion, a, true)
}

func completerConfiguration(in prompt.Document) []prompt.Suggest {
	a := in.GetWordBeforeCursor()
	a = strings.TrimSpace(a)
	d := in.TextBeforeCursor()
	if len(strings.Split(d, " ")) > 2 {
		return []prompt.Suggest{}
	}
	return prompt.FilterHasPrefix(configurationSuggestion, a, true)
}

func completerQOS(in prompt.Document) []prompt.Suggest {
	return nil
}

func CompleterGNB(in prompt.Document) []prompt.Suggest {
	a := in.TextBeforeCursor()
	var split = strings.Split(a, " ")
	w := in.GetWordBeforeCursor()
	if len(split) > 1 {
		var v = split[0]
		if v == "pdu-session" {
			return completerPDU(in)
		}
		if v == "qos" {
			return completerQOS(in)
		}
		if v == "user" {
			return completerUE(in)
		}
		if v == "configuration" {
			return completerConfiguration(in)
		}
		return prompt.FilterHasPrefix(GNBSuggestion, v, true)
	}
	return prompt.FilterHasPrefix(GNBSuggestion, w, true)

}
