package api

const database = "free5gc"

const authSubsDataColl = "subscriptionData.authenticationData.authenticationSubscription"
const amDataColl = "subscriptionData.provisionedData.amData"
const smDataColl = "subscriptionData.provisionedData.smData"
const smfSelDataColl = "subscriptionData.provisionedData.smfSelectionSubscriptionData"
const amPolicyDataColl = "policyData.ues.amData"
const smPolicyDataColl = "policyData.ues.smData"
const amf3gppAccessColl = "subscriptionData.contextData.amf3gppAccess"
const urilistColl = "urilist"
const nfProfileColl = "NfProfile"

// DatabaseCollectionList of DB
var DatabaseCollectionList = []string{
	authSubsDataColl,
	amDataColl,
	smDataColl,
	smfSelDataColl,
	amPolicyDataColl,
	smPolicyDataColl,
	amf3gppAccessColl,
	urilistColl,
	nfProfileColl,
}
