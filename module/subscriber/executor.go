package subscriber

import (
	"fmt"
	"free5gc-cli/logger"
	"free5gc-cli/module/subscriber/api"
	"strconv"
	"strings"

	"github.com/c-bata/go-prompt"
)

func registerSubscriber(ueid string, plmn string) {

}

func removeIndex(s []prompt.Suggest, index int, length int) []prompt.Suggest {
	if length == 1 {
		return []prompt.Suggest{}
	}
	if index == length-1 {
		return append(s[:index-1])
	}
	return append(s[:index], s[index+1:]...)
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
		logger.SubscriberLog.Infoln(fmt.Sprintf("Found %d subscriber", len(subs)))
		for i := 0; i < len(subs); i++ {
			l = append(l, prompt.Suggest{Text: subs[i].UeId + "/" + subs[i].PlmnID,
				Description: "Remove " + subs[i].UeId + " from plmn " + subs[i].PlmnID})
			logger.SubscriberLog.Infoln(fmt.Sprintf("Found subscriber with %s from plmn %s", subs[i].UeId, subs[i].PlmnID))
		}
		supiSuggestion = &l
		return
	}

	if cmd[1] == "flush" {
		subs := api.GetSubscribers()
		if len(subs) == 0 {
			logger.SubscriberLog.Infoln("No subscriber to remove")
			return
		}
		logger.SubscriberLog.Infoln(fmt.Sprintf("Removing %d subscriber", len(subs)))
		for i := 0; i < len(subs); i++ {
			api.DeleteSubscriberByID(subs[i].UeId, subs[i].PlmnID)
			logger.SubscriberLog.Infoln(fmt.Sprintf("Removing subscriber %s from plmn %s", subs[i].UeId, subs[i].PlmnID))
		}
		supiSuggestion = &[]prompt.Suggest{}
		return
	}

	if cmd[1] == "delete" && len(cmd) > 2 {
		tmp := strings.Split(cmd[2], "/")
		if len(tmp) == 2 {
			logger.SubscriberLog.Infoln(fmt.Sprintf("Removing subscriber %s from plmn %s", tmp[0], tmp[1]))
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

	if cmd[1] == "register" {

		if len(cmd) >= 6 {
			supi := cmd[3]
			plmn := cmd[5]

			sub := api.GetSubscriberByID(supi, plmn)
			if sub.AuthenticationSubscription.SequenceNumber != "" {
				logger.SubscriberLog.Infoln(fmt.Sprintf("Existing subscriber with supi %s in PLMN %s\n---", supi, plmn))
				return
			}
			logger.SubscriberLog.Infoln(fmt.Sprintf("Register new subscriber %s in PLMN %s\n---", supi, plmn))
			api.PostSubscriberByID(supi, plmn, *SubsDataConfig)

			sugg := prompt.Suggest{Text: supi + "/" + plmn, Description: "Remove " + supi + " from plmn " + plmn}
			l := append(*supiSuggestion, sugg)
			supiSuggestion = &l

		}
		return

	}

	if cmd[1] == "random" {

		if len(cmd) >= 6 {
			rng := cmd[3]
			plmn := cmd[5]

			n, err := strconv.Atoi(rng)
			if err != nil {
				logger.SubscriberLog.Errorln(fmt.Sprintf("Error parsing the range"))
				return
			}

			logger.SubscriberLog.Infoln(fmt.Sprintf("Generating %s subscribers from plmn %s", rng, plmn))
			logger.SubscriberLog.Infoln(fmt.Sprintf("Generated supi would be of the form ismi-%sXXXXXXXX", plmn))

			i := 0
			j := 0
			for i < n {
				supi := fmt.Sprintf("imsi-%s%08d", plmn, j)
				sub := api.GetSubscriberByID(supi, plmn)
				if sub.AuthenticationSubscription.SequenceNumber == "" {
					logger.SubscriberLog.Infoln(fmt.Sprintf("Register new subscriber with supi %s in PLMN %s", supi, plmn))
					api.PostSubscriberByID(supi, plmn, *SubsDataConfig)
					sugg := prompt.Suggest{Text: supi + "/" + plmn, Description: "Remove " + supi + " from plmn " + plmn}
					l := append(*supiSuggestion, sugg)
					supiSuggestion = &l
					i++
				}
				j++
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

}
