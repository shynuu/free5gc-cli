package subscriber

import (
	"fmt"
	"free5gc-cli/module/subscriber/api"
	"strings"

	"github.com/c-bata/go-prompt"
)

func removeIndex(s []prompt.Suggest, index int, length int) []prompt.Suggest {
	if index == length-1 {
		return append(s[:index-1])
	}
	return append(s[:index], s[index+1:]...)
}

func executorRegister(in string) {
	// api.TestData()
}

func executorConfiguration(in string) {
	s := strings.TrimSpace(in)
	if s == "configuration reload" {
		Reload()
	}
}

func executorUser(in string) {
	cmd := strings.Split(strings.TrimSpace(in), " ")

	if len(cmd) < 2 {
		return
	}

	if cmd[1] == "list" {

		subs := api.GetSubscribers()
		var l []prompt.Suggest
		fmt.Println(fmt.Sprintf("Found %d user\n------------------------", len(subs)))
		for i := 0; i < len(subs); i++ {
			l = append(l, prompt.Suggest{Text: subs[i].UeId + "/" + subs[i].PlmnID,
				Description: "Remove " + subs[i].UeId + " from plmn " + subs[i].PlmnID})
			fmt.Println(fmt.Sprintf("%s %s", subs[i].UeId, subs[i].PlmnID))
		}
		supiSuggestion = &l
		return
	}

	if cmd[1] == "flush" {
		subs := api.GetSubscribers()
		if len(subs) == 0 {
			fmt.Println("No user to remove")
			return
		}
		fmt.Println(fmt.Sprintf("Removing %d user\n------------------------", len(subs)))
		for i := 0; i < len(subs); i++ {
			api.DeleteSubscriberByID(subs[i].UeId, subs[i].PlmnID)
			fmt.Println(fmt.Sprintf("Removing %s %s from user", subs[i].UeId, subs[i].PlmnID))
		}
		supiSuggestion = &[]prompt.Suggest{}
		return
	}

	if cmd[1] == "remove" && len(cmd) > 2 {
		tmp := strings.Split(cmd[2], "/")
		if len(tmp) == 2 {
			fmt.Println(fmt.Sprintf("Removing user %s from %s\n------------------------", tmp[0], tmp[1]))
			api.DeleteSubscriberByID(tmp[0], tmp[1])
			for i := 0; i < len(*supiSuggestion); i++ {
				if fmt.Sprintf("%s/%s", tmp[0], tmp[1]) == (*supiSuggestion)[i].Text {
					t := removeIndex(*supiSuggestion, i, len(*supiSuggestion))
					supiSuggestion = &t
				}
			}

		}
		return
	}
}

// Executor parse CLI
func Executor(in string) {

	if strings.HasPrefix(in, "configuration") {
		executorConfiguration(in)
	}

	if strings.HasPrefix(in, "user") {
		executorUser(in)
	}

	// register configuration-file.yaml
	// add imsi-207839393934
	if strings.HasPrefix(in, "register") {
		executorRegister(strings.TrimSpace(strings.ReplaceAll(in, "register", "")))
	}

}
