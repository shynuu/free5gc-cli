package qos

import (
	"strings"

	"github.com/c-bata/go-prompt"
)

// list
// add
// delete

// QOSSuggestion suggestions
var QOSSuggestion = []prompt.Suggest{
	{Text: "mark", Description: "Mark packet with DSCP based on packet match"},
	{Text: "flush", Description: "Flush iptables table mangle"},
	// {Text: "list", Description: "List all rules"},
	{Text: "configuration", Description: "Manage the module configuration"},
	{Text: "exit", Description: "Exit the QoS module"},
}

var configurationSuggestion = []prompt.Suggest{
	{Text: "reload", Description: "Reload the QoS configuration module"},
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

// IPSuggestion holds the ip of the module
var IPSuggestion = &[]prompt.Suggest{}

// RulesSuggestion holds the rules of the module
var RulesSuggestion = &[]prompt.Suggest{}

// TEIDSuggestion holds the TEID of the packets
var TEIDSuggestion = &[]prompt.Suggest{}

// PortSuggestion holds the TEID of the packets
var PortSuggestion = &[]prompt.Suggest{}

// completerConfiguration
func completerConfiguration(in prompt.Document) []prompt.Suggest {
	a := in.GetWordBeforeCursor()
	a = strings.TrimSpace(a)
	d := in.TextBeforeCursor()
	if len(strings.Split(d, " ")) > 2 {
		return []prompt.Suggest{}
	}
	return prompt.FilterHasPrefix(configurationSuggestion, a, true)
}

// qos flush
func completerFlush(in prompt.Document) []prompt.Suggest {
	return []prompt.Suggest{}
}

// mark --dscp 10 --source-ip 10.10.0.1 --destination-ip 10.10.0.1 --teid 00101010 --protocol tcp/udp --source-port 80 --destination-port 2000
func completerMark(in prompt.Document) []prompt.Suggest {
	a := in.GetWordBeforeCursor()
	a = strings.TrimSpace(a)
	d := strings.Split(in.TextBeforeCursor(), " ")
	if d[1] == "--set-phb" {
		l := len(d)

		if l > 15 {
			return []prompt.Suggest{}
		}

		if l == 3 {
			return prompt.FilterHasPrefix(PHBSuggestion, a, true)
		}

		if l == 4 {
			return prompt.FilterHasPrefix([]prompt.Suggest{
				{Text: "--destination-ip", Description: "Specify the destination IP of the outer packet"},
			}, a, true)
		}

		if l == 5 {
			return prompt.FilterHasPrefix(*IPSuggestion, a, true)
		}

		if l == 6 {
			return prompt.FilterHasPrefix([]prompt.Suggest{
				{Text: "--source-ip", Description: "Specify the source IP of the outer packet"},
			}, a, true)
		}

		if l == 7 {
			return prompt.FilterHasPrefix(*IPSuggestion, a, true)
		}

		if l == 8 {
			return prompt.FilterHasPrefix([]prompt.Suggest{
				{Text: "--teid", Description: "Specify the TEID of the GTP-U tunnel"},
			}, a, true)
		}

		if l == 9 {
			return prompt.FilterHasPrefix(*TEIDSuggestion, a, true)
		}

		if l == 10 {
			return prompt.FilterHasPrefix([]prompt.Suggest{
				{Text: "--protocol", Description: "Specify the protocol of the the inner packet"},
			}, a, true)
		}

		if l == 11 {
			return prompt.FilterHasPrefix([]prompt.Suggest{
				{Text: "tcp", Description: ""},
				{Text: "udp", Description: ""},
			}, a, true)
		}

		if l == 12 {
			return prompt.FilterHasPrefix([]prompt.Suggest{
				{Text: "--destination-port", Description: "Specify the destination port of the the inner packet"},
			}, a, true)
		}

		if l == 13 {
			return prompt.FilterHasPrefix(*PortSuggestion, a, true)
		}

		if l == 14 {
			return prompt.FilterHasPrefix([]prompt.Suggest{
				{Text: "--source-port", Description: "Specify the source port of the the inner packet"},
			}, a, true)
		}

		if l == 15 {
			return prompt.FilterHasPrefix(*PortSuggestion, a, true)
		}

		return []prompt.Suggest{}
	}

	return prompt.FilterHasPrefix([]prompt.Suggest{
		{Text: "--set-phb", Description: "Specify the DSCP field to be applied"},
	}, a, true)
}

// CompleterQOS shows the QOS module commands
func CompleterQOS(in prompt.Document) []prompt.Suggest {
	a := in.TextBeforeCursor()
	var split = strings.Split(a, " ")
	w := in.GetWordBeforeCursor()
	if len(split) > 1 {
		var v = split[0]
		if v == "mark" {
			return completerMark(in)
		}
		if v == "flush" {
			return completerFlush(in)
		}
		if v == "configuration" {
			return completerConfiguration(in)
		}
		return prompt.FilterHasPrefix(QOSSuggestion, v, true)
	}
	return prompt.FilterHasPrefix(QOSSuggestion, w, true)

}
