package api

import (
	"encoding/json"
	"free5gc-cli/lib/MongoDBLibrary"
	"free5gc-cli/lib/openapi/models"
	"free5gc-cli/logger"

	"go.mongodb.org/mongo-driver/bson"
)

// GetSubscribers returns subscriber list
func GetSubscribers() []SubsListIE {

	var subsList []SubsListIE = make([]SubsListIE, 0)
	amDataList := MongoDBLibrary.RestfulAPIGetMany(amDataColl, bson.M{})
	for _, amData := range amDataList {
		ueId := amData["ueId"]
		servingPlmnId := amData["servingPlmnId"]
		var tmp = SubsListIE{
			PlmnID: servingPlmnId.(string),
			UeId:   ueId.(string),
		}
		subsList = append(subsList, tmp)
	}

	return subsList
}

// GetSubscriberByID returns the subscriber by IMSI(ueId) and PlmnID(servingPlmnId)
func GetSubscriberByID(ueId string, servingPlmnId string) SubsData {

	var subsData SubsData

	filterUeIdOnly := bson.M{"ueId": ueId}
	filter := bson.M{"ueId": ueId, "servingPlmnId": servingPlmnId}

	authSubsDataInterface := MongoDBLibrary.RestfulAPIGetOne(authSubsDataColl, filterUeIdOnly)
	amDataDataInterface := MongoDBLibrary.RestfulAPIGetOne(amDataColl, filter)
	smDataDataInterface := MongoDBLibrary.RestfulAPIGetMany(smDataColl, filter)
	smfSelDataInterface := MongoDBLibrary.RestfulAPIGetOne(smfSelDataColl, filter)
	amPolicyDataInterface := MongoDBLibrary.RestfulAPIGetOne(amPolicyDataColl, filterUeIdOnly)
	smPolicyDataInterface := MongoDBLibrary.RestfulAPIGetOne(smPolicyDataColl, filterUeIdOnly)

	var authSubsData models.AuthenticationSubscription
	json.Unmarshal(mapToByte(authSubsDataInterface), &authSubsData)
	var amDataData models.AccessAndMobilitySubscriptionData
	json.Unmarshal(mapToByte(amDataDataInterface), &amDataData)

	smDataDataLength := len(smDataDataInterface)
	var smDataData []models.SessionManagementSubscriptionData
	smDataData = make([]models.SessionManagementSubscriptionData, smDataDataLength, smDataDataLength)
	for i, sm := range smDataDataInterface {
		json.Unmarshal(mapToByte(sm), &smDataData[i])
	}

	var smfSelData models.SmfSelectionSubscriptionData
	json.Unmarshal(mapToByte(smfSelDataInterface), &smfSelData)
	var amPolicyData models.AmPolicyData
	json.Unmarshal(mapToByte(amPolicyDataInterface), &amPolicyData)
	var smPolicyData models.SmPolicyData
	json.Unmarshal(mapToByte(smPolicyDataInterface), &smPolicyData)

	subsData = SubsData{
		PlmnID:                            servingPlmnId,
		UeId:                              ueId,
		AuthenticationSubscription:        authSubsData,
		AccessAndMobilitySubscriptionData: amDataData,
		SessionManagementSubscriptionData: smDataData,
		SmfSelectionSubscriptionData:      smfSelData,
		AmPolicyData:                      amPolicyData,
		SmPolicyData:                      smPolicyData,
	}

	return subsData

}

// PostSubscriberByID subscriber by IMSI(ueId) and PlmnID(servingPlmnId)
func PostSubscriberByID(ueId string, servingPlmnId string, subsData SubsData) {

	filterUeIdOnly := bson.M{"ueId": ueId}
	filter := bson.M{"ueId": ueId, "servingPlmnId": servingPlmnId}

	authSubsBsonM := toBsonM(subsData.AuthenticationSubscription)
	authSubsBsonM["ueId"] = ueId
	amDataBsonM := toBsonM(subsData.AccessAndMobilitySubscriptionData)
	amDataBsonM["ueId"] = ueId
	amDataBsonM["servingPlmnId"] = servingPlmnId
	smfSelSubsBsonM := toBsonM(subsData.SmfSelectionSubscriptionData)
	smfSelSubsBsonM["ueId"] = ueId
	smfSelSubsBsonM["servingPlmnId"] = servingPlmnId
	for _, sm := range subsData.SessionManagementSubscriptionData {
		smDataBsonM := toBsonM(sm)
		smDataSnssaiBsonM := toBsonM(sm.SingleNssai)
		filterSnssai := bson.M{"ueId": ueId, "servingPlmnId": servingPlmnId, "singleNssai": smDataSnssaiBsonM}
		smDataBsonM["ueId"] = ueId
		smDataBsonM["servingPlmnId"] = servingPlmnId
		MongoDBLibrary.RestfulAPIPost(smDataColl, filterSnssai, smDataBsonM)
	}
	amPolicyDataBsonM := toBsonM(subsData.AmPolicyData)
	amPolicyDataBsonM["ueId"] = ueId
	smPolicyDataBsonM := toBsonM(subsData.SmPolicyData)
	smPolicyDataBsonM["ueId"] = ueId

	MongoDBLibrary.RestfulAPIPost(authSubsDataColl, filterUeIdOnly, authSubsBsonM)
	MongoDBLibrary.RestfulAPIPost(amDataColl, filter, amDataBsonM)
	MongoDBLibrary.RestfulAPIPost(smfSelDataColl, filter, smfSelSubsBsonM)
	MongoDBLibrary.RestfulAPIPost(amPolicyDataColl, filterUeIdOnly, amPolicyDataBsonM)
	MongoDBLibrary.RestfulAPIPost(smPolicyDataColl, filterUeIdOnly, smPolicyDataBsonM)

}

// PatchSubscriberByID subscriber by IMSI(ueId) and PlmnID(servingPlmnId)
func PatchSubscriberByID(ueId string, servingPlmnId string) SubsData {
	logger.FreecliLog.Infoln("Updating Subscriber Data", ueId)

	var subsData SubsData

	filterUeIdOnly := bson.M{"ueId": ueId}
	filter := bson.M{"ueId": ueId, "servingPlmnId": servingPlmnId}

	authSubsBsonM := toBsonM(subsData.AuthenticationSubscription)
	authSubsBsonM["ueId"] = ueId
	amDataBsonM := toBsonM(subsData.AccessAndMobilitySubscriptionData)
	amDataBsonM["ueId"] = ueId
	amDataBsonM["servingPlmnId"] = servingPlmnId
	smDataBsonM := toBsonM(subsData.SessionManagementSubscriptionData)
	smDataBsonM["ueId"] = ueId
	smDataBsonM["servingPlmnId"] = servingPlmnId
	smfSelSubsBsonM := toBsonM(subsData.SmfSelectionSubscriptionData)
	smfSelSubsBsonM["ueId"] = ueId
	smfSelSubsBsonM["servingPlmnId"] = servingPlmnId
	amPolicyDataBsonM := toBsonM(subsData.AmPolicyData)
	amPolicyDataBsonM["ueId"] = ueId
	smPolicyDataBsonM := toBsonM(subsData.SmPolicyData)
	smPolicyDataBsonM["ueId"] = ueId

	MongoDBLibrary.RestfulAPIMergePatch(authSubsDataColl, filterUeIdOnly, authSubsBsonM)
	MongoDBLibrary.RestfulAPIMergePatch(amDataColl, filter, amDataBsonM)
	MongoDBLibrary.RestfulAPIMergePatch(smDataColl, filter, smDataBsonM)
	MongoDBLibrary.RestfulAPIMergePatch(smfSelDataColl, filter, smfSelSubsBsonM)
	MongoDBLibrary.RestfulAPIMergePatch(amPolicyDataColl, filterUeIdOnly, amPolicyDataBsonM)
	MongoDBLibrary.RestfulAPIMergePatch(smPolicyDataColl, filterUeIdOnly, smPolicyDataBsonM)

	return subsData

}

// DeleteSubscriberByID deletes a subscriber by IMSI(ueId) and PlmnID(servingPlmnId)
func DeleteSubscriberByID(ueId string, servingPlmnId string) {

	filterUeIdOnly := bson.M{"ueId": ueId}
	filter := bson.M{"ueId": ueId, "servingPlmnId": servingPlmnId}

	MongoDBLibrary.RestfulAPIDeleteOne(authSubsDataColl, filterUeIdOnly)
	MongoDBLibrary.RestfulAPIDeleteOne(amDataColl, filter)
	MongoDBLibrary.RestfulAPIDeleteMany(smDataColl, filter)
	MongoDBLibrary.RestfulAPIDeleteOne(smfSelDataColl, filter)
	MongoDBLibrary.RestfulAPIDeleteOne(amPolicyDataColl, filterUeIdOnly)
	MongoDBLibrary.RestfulAPIDeleteOne(smPolicyDataColl, filterUeIdOnly)

}

func TestData() {
	var subsData SubsData

	authSubsData := models.AuthenticationSubscription{
		AuthenticationManagementField: "8000",
		AuthenticationMethod:          "5G_AKA", // "5G_AKA", "EAP_AKA_PRIME"
		Milenage: &models.Milenage{
			Op: &models.Op{
				EncryptionAlgorithm: 0,
				EncryptionKey:       0,
				OpValue:             "c9e8763286b5b9ffbdf56e1297d0887b", // Required
			},
		},
		Opc: &models.Opc{
			EncryptionAlgorithm: 0,
			EncryptionKey:       0,
			OpcValue:            "981d464c7c52eb6e5036234984ad0bcf", // Required
		},
		PermanentKey: &models.PermanentKey{
			EncryptionAlgorithm: 0,
			EncryptionKey:       0,
			PermanentKeyValue:   "5122250214c33e723a5dd523fc145fc0", // Required
		},
		SequenceNumber: "16f3b3f70fc2",
	}

	amDataData := models.AccessAndMobilitySubscriptionData{
		Gpsis: []string{
			"msisdn-0900000000",
		},
		Nssai: &models.Nssai{
			DefaultSingleNssais: []models.Snssai{
				{
					Sd:  "010203",
					Sst: 1,
				},
				{
					Sd:  "112233",
					Sst: 1,
				},
			},
			SingleNssais: []models.Snssai{
				{
					Sd:  "010203",
					Sst: 1,
				},
				{
					Sd:  "112233",
					Sst: 1,
				},
			},
		},
		SubscribedUeAmbr: &models.AmbrRm{
			Downlink: "1000 Kbps",
			Uplink:   "1000 Kbps",
		},
	}

	smDataData := []models.SessionManagementSubscriptionData{{
		SingleNssai: &models.Snssai{
			Sst: 1,
			Sd:  "010203",
		},
		DnnConfigurations: map[string]models.DnnConfiguration{
			"internet": models.DnnConfiguration{
				PduSessionTypes: &models.PduSessionTypes{
					DefaultSessionType:  models.PduSessionType_IPV4,
					AllowedSessionTypes: []models.PduSessionType{models.PduSessionType_IPV4},
				},
				SscModes: &models.SscModes{
					DefaultSscMode:  models.SscMode__1,
					AllowedSscModes: []models.SscMode{models.SscMode__1},
				},
				SessionAmbr: &models.Ambr{
					Downlink: "1000 Kbps",
					Uplink:   "1000 Kbps",
				},
				Var5gQosProfile: &models.SubscribedDefaultQos{
					Var5qi: 9,
					Arp: &models.Arp{
						PriorityLevel: 8,
					},
					PriorityLevel: 8,
				},
			},
		},
	}}

	smfSelData := models.SmfSelectionSubscriptionData{
		SubscribedSnssaiInfos: map[string]models.SnssaiInfo{
			"01010203": {
				DnnInfos: []models.DnnInfo{
					{
						Dnn: "internet",
					},
				},
			},
			"01112233": {
				DnnInfos: []models.DnnInfo{
					{
						Dnn: "internet",
					},
				},
			},
		},
	}

	amPolicyData := models.AmPolicyData{
		SubscCats: []string{
			"free5gc",
		},
	}

	smPolicyData := models.SmPolicyData{
		SmPolicySnssaiData: map[string]models.SmPolicySnssaiData{
			"01010203": {
				Snssai: &models.Snssai{
					Sd:  "010203",
					Sst: 1,
				},
				SmPolicyDnnData: map[string]models.SmPolicyDnnData{
					"internet": {
						Dnn: "internet",
					},
				},
			},
			"01112233": {
				Snssai: &models.Snssai{
					Sd:  "112233",
					Sst: 1,
				},
				SmPolicyDnnData: map[string]models.SmPolicyDnnData{
					"internet": {
						Dnn: "internet",
					},
				},
			},
		},
	}

	servingPlmnId := "20893"
	ueId := "imsi-2089300007487"

	subsData = SubsData{
		PlmnID:                            servingPlmnId,
		UeId:                              ueId,
		AuthenticationSubscription:        authSubsData,
		AccessAndMobilitySubscriptionData: amDataData,
		SessionManagementSubscriptionData: smDataData,
		SmfSelectionSubscriptionData:      smfSelData,
		AmPolicyData:                      amPolicyData,
		SmPolicyData:                      smPolicyData,
	}

	PostSubscriberByID(ueId, servingPlmnId, subsData)
}
