package api

import (
	"free5gc-cli/lib/nas"
	"free5gc-cli/lib/nas/nasMessage"
	"free5gc-cli/lib/nas/nasTestpacket"
	"free5gc-cli/lib/nas/nasType"
	"free5gc-cli/lib/nas/security"
	"free5gc-cli/lib/ngap"
	"free5gc-cli/logger"
)

func Registration(ueId string, plmn string, sst int, sd string) bool {

	var n int
	var sendMsg []byte
	var recvMsg = make([]byte, 2048)

	// RAN connect to AMF
	conn, err := ConnectToAmf(APIConfig.Configuration.AmfInterface.IPv4Addr,
		APIConfig.Configuration.NGRANInterface.IPv4Addr,
		APIConfig.Configuration.AmfInterface.Port,
		APIConfig.Configuration.NGRANInterface.Port)
	if err != nil {
		logger.GNBLog.Errorln("Error connecting to the AMF")
		return false
	}

	// send NGSetupRequest Msg
	sendMsg, err = GetNGSetupRequest([]byte("\x00\x01\x02"), 24, "free5gc")
	_, err = conn.Write(sendMsg)
	if err != nil {
		logger.GNBLog.Errorln("Error sending NGSetup")
		return false
	}

	// receive NGSetupResponse Msg
	n, err = conn.Read(recvMsg)
	ngapPdu, err := ngap.Decoder(recvMsg[:n])
	if err != nil {
		logger.GNBLog.Errorln("Error decoding NGAP")
		return false
	}

	// New UE
	// ue := test.NewRanUeContext("imsi-2089300007487", 1, security.AlgCiphering128NEA2, security.AlgIntegrity128NIA2)
	ue := NewRanUeContext(ueId, 1, security.AlgCiphering128NEA0, security.AlgIntegrity128NIA2)
	ue.AmfUeNgapId = 1
	ue.AuthenticationSubs = GetAuthSubscription(APIConfig.Configuration.Security.K,
		APIConfig.Configuration.Security.OPC,
		APIConfig.Configuration.Security.OP)

	// send InitialUeMessage(Registration Request)(imsi-2089300007487)
	mobileIdentity5GS := nasType.MobileIdentity5GS{
		Len:    12, // suci
		Buffer: []uint8{0x01, 0x02, 0xf8, 0x39, 0xf0, 0xff, 0x00, 0x00, 0x00, 0x00, 0x47, 0x78},
	}

	ueSecurityCapability := ue.GetUESecurityCapability()
	registrationRequest := nasTestpacket.GetRegistrationRequest(
		nasMessage.RegistrationType5GSInitialRegistration, mobileIdentity5GS, nil, ueSecurityCapability, nil, nil, nil)
	sendMsg, err = GetInitialUEMessage(ue.RanUeNgapId, registrationRequest, "")
	if err != nil {
		logger.GNBLog.Errorln("Error building Initial UE Message")
		return false
	}
	_, err = conn.Write(sendMsg)
	if err != nil {
		logger.GNBLog.Errorln("Error sending Initial UE Message")
		return false
	}

	// receive NAS Authentication Request Msg
	n, err = conn.Read(recvMsg)
	if err != nil {
		logger.GNBLog.Errorln("Error reading NAS Authentication Request Msg")
		return false
	}
	ngapPdu, err = ngap.Decoder(recvMsg[:n])
	if err != nil {
		logger.GNBLog.Errorln("Error decoding NAS Authentication Request Msg")
		return false
	}

	// Calculate for RES*
	nasPdu := GetNasPdu(ue, ngapPdu.InitiatingMessage.Value.DownlinkNASTransport)
	rand := nasPdu.AuthenticationRequest.GetRANDValue()
	resStat := ue.DeriveRESstarAndSetKey(ue.AuthenticationSubs, rand[:], APIConfig.Configuration.Security.NetworkName)

	// send NAS Authentication Response
	pdu := nasTestpacket.GetAuthenticationResponse(resStat, "")
	sendMsg, err = GetUplinkNASTransport(ue.AmfUeNgapId, ue.RanUeNgapId, pdu)
	if err != nil {
		logger.GNBLog.Errorln("Error building NAS UplinkNASTransport")
		return false
	}
	_, err = conn.Write(sendMsg)
	if err != nil {
		logger.GNBLog.Errorln("Error sending NAS UplinkNASTransport")
		return false
	}

	// receive NAS Security Mode Command Msg
	n, err = conn.Read(recvMsg)
	if err != nil {
		logger.GNBLog.Errorln("Error reading NAS Security Mode Command Msg")
		return false
	}
	ngapPdu, err = ngap.Decoder(recvMsg[:n])
	if err != nil {
		logger.GNBLog.Errorln("Error decoding NAS Security Mode Command Msg")
		return false
	}
	nasPdu = GetNasPdu(ue, ngapPdu.InitiatingMessage.Value.DownlinkNASTransport)

	// send NAS Security Mode Complete Msg
	registrationRequestWith5GMM := nasTestpacket.GetRegistrationRequest(nasMessage.RegistrationType5GSInitialRegistration,
		mobileIdentity5GS, nil, ueSecurityCapability, ue.Get5GMMCapability(), nil, nil)
	pdu = nasTestpacket.GetSecurityModeComplete(registrationRequestWith5GMM)
	pdu, err = EncodeNasPduWithSecurity(ue, pdu, nas.SecurityHeaderTypeIntegrityProtectedAndCipheredWithNew5gNasSecurityContext, true, true)
	if err != nil {
		logger.GNBLog.Errorln("Error encoding NAS PDU with Security")
		return false
	}
	sendMsg, err = GetUplinkNASTransport(ue.AmfUeNgapId, ue.RanUeNgapId, pdu)
	if err != nil {
		logger.GNBLog.Errorln("Error sending NAS PDU with Security")
		return false
	}
	_, err = conn.Write(sendMsg)

	// receive ngap Initial Context Setup Request Msg
	n, err = conn.Read(recvMsg)
	if err != nil {
		logger.GNBLog.Errorln("Error reading ngap Initial Context Setup Request Msg")
		return false
	}
	ngapPdu, err = ngap.Decoder(recvMsg[:n])
	if err != nil {
		logger.GNBLog.Errorln("Error decoding ngap Initial Context Setup Request Msg")
		return false
	}

	// send ngap Initial Context Setup Response Msg
	sendMsg, err = GetInitialContextSetupResponse(ue.AmfUeNgapId, ue.RanUeNgapId)
	_, err = conn.Write(sendMsg)
	if err != nil {
		logger.GNBLog.Errorln("Error sending ngap Initial Context Setup Response Msg")
		return false
	}

	// send NAS Registration Complete Msg
	pdu = nasTestpacket.GetRegistrationComplete(nil)
	pdu, err = EncodeNasPduWithSecurity(ue, pdu, nas.SecurityHeaderTypeIntegrityProtectedAndCiphered, true, false)
	if err != nil {
		logger.GNBLog.Errorln("Error encoding NAS Registration Complete Msg with Security")
		return false
	}

	sendMsg, err = GetUplinkNASTransport(ue.AmfUeNgapId, ue.RanUeNgapId, pdu)
	if err != nil {
		logger.GNBLog.Errorln("Error building NAS Registration Complete Msg with Security")
		return false
	}

	_, err = conn.Write(sendMsg)
	if err != nil {
		logger.GNBLog.Errorln("Error sending NAS Registration Complete Msg with Security")
		return false
	}

	return true

}

func DeRegistration(ueId string, plmn string, sst int, sd string) {

}

func PDUSessionRequest(ueId string, plmn string, sst int, sd string, sessionId int) {

	// RAN connect to UPF
	// upfConn, err := ConnectToUpf(APIConfig.Configuration.GTPInterface.IPv4Addr,
	// 	APIConfig.Configuration.UpfInterface.IPv4Addr,
	// 	APIConfig.Configuration.GTPInterface.Port,
	// 	APIConfig.Configuration.UpfInterface.Port)s

}

func PDUSessionRelease(ueId string, plmn string, sst int, sd string, sessionId int) {

}
