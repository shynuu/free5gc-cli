package custom

import "free5gc-cli/lib/openapi/models"

type SubsData struct {
	PlmnID                            string                                   `json:"plmnID" yaml:"plmnID"`
	UeId                              string                                   `json:"ueId" yaml:"ueId"`
	AuthenticationSubscription        models.AuthenticationSubscription        `json:"AuthenticationSubscription" yaml:"AuthenticationSubscription"`
	AccessAndMobilitySubscriptionData models.AccessAndMobilitySubscriptionData `json:"AccessAndMobilitySubscriptionData" yaml:"AccessAndMobilitySubscriptionData"`
	SessionManagementSubscriptionData models.SessionManagementSubscriptionData `json:"SessionManagementSubscriptionData" yaml:"SessionManagementSubscriptionData"`
	SmfSelectionSubscriptionData      models.SmfSelectionSubscriptionData      `json:"SmfSelectionSubscriptionData" yaml:"SmfSelectionSubscriptionData"`
	AmPolicyData                      models.AmPolicyData                      `json:"AmPolicyData" yaml:"AmPolicyData"`
	SmPolicyData                      models.SmPolicyData                      `json:"SmPolicyData" yaml:"SmPolicyData"`
}

type SubsListIE struct {
	PlmnID string `json:"plmnID" yaml:"plmnID"`
	UeId   string `json:"ueId" yaml:"ueId"`
}
