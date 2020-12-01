package api

import (
	"encoding/hex"
	"errors"
	"fmt"
	"free5gc-cli/lib/nas"
	"free5gc-cli/lib/nas/nasMessage"
	"free5gc-cli/lib/nas/nasTestpacket"
	"free5gc-cli/lib/nas/nasType"
	"free5gc-cli/lib/nas/security"
	"free5gc-cli/lib/ngap"
	"free5gc-cli/lib/ngap/ngapType"
	"free5gc-cli/logger"
	"net"
	"strconv"
	"time"

	"github.com/ishidawataru/sctp"

	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
)

var amfConn *sctp.SCTPConn
var upfConn *net.UDPConn

func checkAmfConnection() error {
	if amfConn == nil {
		amfC, err := ConnectToAmf(APIConfig.Configuration.AmfInterface.IPv4Addr,
			APIConfig.Configuration.NGRANInterface.IPv4Addr,
			APIConfig.Configuration.AmfInterface.Port,
			APIConfig.Configuration.NGRANInterface.Port)
		if err != nil {
			return err
		}
		amfConn = amfC
		return nil
	}
	return nil
}

func TestPing(sourceIp string, destinationIp string) error {
	gtpHdr, err := hex.DecodeString("32ff00340000000100000000")
	if err != nil {
		logger.GNBLog.Errorln("Error decoding GTP Header")
		return err
	}
	icmpData, err := hex.DecodeString("8c870d0000000000101112131415161718191a1b1c1d1e1f202122232425262728292a2b2c2d2e2f3031323334353637")
	if err != nil {
		logger.GNBLog.Errorln("Error decoding ICMP Data")
		return err
	}

	ipv4hdr := ipv4.Header{
		Version:  4,
		Len:      20,
		Protocol: 1,
		Flags:    0,
		TotalLen: 48,
		TTL:      64,
		Src:      net.ParseIP(sourceIp).To4(),
		Dst:      net.ParseIP(destinationIp).To4(),
		ID:       1,
	}
	checksum := CalculateIpv4HeaderChecksum(&ipv4hdr)
	ipv4hdr.Checksum = int(checksum)

	v4HdrBuf, err := ipv4hdr.Marshal()
	if err != nil {
		logger.GNBLog.Errorln("Error Marshaling IP Header")
		return err
	}
	tt := append(gtpHdr, v4HdrBuf...)

	m := icmp.Message{
		Type: ipv4.ICMPTypeEcho, Code: 0,
		Body: &icmp.Echo{
			ID: 12394, Seq: 1,
			Data: icmpData,
		},
	}
	b, err := m.Marshal(nil)
	if err != nil {
		logger.GNBLog.Errorln("Error ICMP Payload")
		return err
	}
	b[2] = 0xaf
	b[3] = 0x88
	_, err = upfConn.Write(append(tt, b...))
	if err != nil {
		logger.GNBLog.Errorln("Error sending ICMP to UPF")
		return err
	}
	return nil

}

func Registration(ueId string) (*RanUeContext, error) {

	var n int
	var sendMsg []byte
	var recvMsg = make([]byte, 2048)

	// RAN connect to AMF
	err := checkAmfConnection()

	if err != nil {
		logger.GNBLog.Errorln("Error connecting to the AMF")
		return nil, err
	}

	// send NGSetupRequest Msg
	sendMsg, err = GetNGSetupRequest([]byte("\x00\x01\x02"), 24, "free5gc")
	_, err = amfConn.Write(sendMsg)
	if err != nil {
		logger.GNBLog.Errorln("Error sending NGSetup")
		return nil, err
	}

	// receive NGSetupResponse Msg
	n, err = amfConn.Read(recvMsg)
	ngapPdu, err := ngap.Decoder(recvMsg[:n])
	if err != nil {
		logger.GNBLog.Errorln("Error decoding NGAP")
		return nil, err
	}

	// New UE
	// ue := test.NewRanUeContext("imsi-2089300007487", 1, security.AlgCiphering128NEA2, security.AlgIntegrity128NIA2)
	ue := NewRanUeContext(ueId, 1, security.AlgCiphering128NEA0, security.AlgIntegrity128NIA2)
	ue.AmfUeNgapId = 1
	ue.AuthenticationSubs = GetAuthSubscription(APIConfig.Configuration.Security.K,
		APIConfig.Configuration.Security.OPC,
		APIConfig.Configuration.Security.OP)

	suci, err := supiToSuci(ueId)
	if err != nil {
		logger.GNBLog.Errorln("Error encoding SUPI to SUCI")
		return nil, err
	}

	// send InitialUeMessage(Registration Request)(imsi- 02 08 93 00 00 74 87)
	mobileIdentity5GS := nasType.MobileIdentity5GS{
		Len:    12, // suci
		Buffer: suci,
	}

	ueSecurityCapability := ue.GetUESecurityCapability()
	registrationRequest := nasTestpacket.GetRegistrationRequest(
		nasMessage.RegistrationType5GSInitialRegistration, mobileIdentity5GS, nil, ueSecurityCapability, nil, nil, nil)
	sendMsg, err = GetInitialUEMessage(ue.RanUeNgapId, registrationRequest, "")
	if err != nil {
		logger.GNBLog.Errorln("Error building Initial UE Message")
		return nil, err
	}
	_, err = amfConn.Write(sendMsg)
	if err != nil {
		logger.GNBLog.Errorln("Error sending Initial UE Message")
		return nil, err
	}

	// receive NAS Authentication Request Msg
	n, err = amfConn.Read(recvMsg)
	if err != nil {
		logger.GNBLog.Errorln("Error reading NAS Authentication Request Msg")
		return nil, err
	}
	ngapPdu, err = ngap.Decoder(recvMsg[:n])
	if err != nil {
		logger.GNBLog.Errorln("Error decoding NAS Authentication Request Msg")
		return nil, err
	}
	if ngapPdu.Present != ngapType.NGAPPDUPresentInitiatingMessage {
		logger.GNBLog.Errorln("Error No NGAP Initiating Message received.")
		return nil, err
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
		return nil, err
	}
	_, err = amfConn.Write(sendMsg)
	if err != nil {
		logger.GNBLog.Errorln("Error sending NAS UplinkNASTransport")
		return nil, err
	}

	// receive NAS Security Mode Command Msg
	n, err = amfConn.Read(recvMsg)
	if err != nil {
		logger.GNBLog.Errorln("Error reading NAS Security Mode Command Msg")
		return nil, err
	}
	ngapPdu, err = ngap.Decoder(recvMsg[:n])
	if err != nil {
		logger.GNBLog.Errorln("Error decoding NAS Security Mode Command Msg")
		return nil, err
	}
	nasPdu = GetNasPdu(ue, ngapPdu.InitiatingMessage.Value.DownlinkNASTransport)
	if nasPdu.GmmHeader.GetMessageType() != nas.MsgTypeSecurityModeCommand {
		logger.GNBLog.Errorln("No Security Mode Command received. Message: " + strconv.Itoa(int(nasPdu.GmmHeader.GetMessageType())))
		return nil, errors.New("No Security Mode Command received. Message: " + strconv.Itoa(int(nasPdu.GmmHeader.GetMessageType())))
	}

	// send NAS Security Mode Complete Msg
	registrationRequestWith5GMM := nasTestpacket.GetRegistrationRequest(nasMessage.RegistrationType5GSInitialRegistration,
		mobileIdentity5GS, nil, ueSecurityCapability, ue.Get5GMMCapability(), nil, nil)
	pdu = nasTestpacket.GetSecurityModeComplete(registrationRequestWith5GMM)
	pdu, err = EncodeNasPduWithSecurity(ue, pdu, nas.SecurityHeaderTypeIntegrityProtectedAndCipheredWithNew5gNasSecurityContext, true, true)
	if err != nil {
		logger.GNBLog.Errorln("Error encoding NAS PDU with Security")
		return nil, err
	}
	sendMsg, err = GetUplinkNASTransport(ue.AmfUeNgapId, ue.RanUeNgapId, pdu)
	if err != nil {
		logger.GNBLog.Errorln("Error sending NAS PDU with Security")
		return nil, err
	}
	_, err = amfConn.Write(sendMsg)

	// receive ngap Initial Context Setup Request Msg
	n, err = amfConn.Read(recvMsg)
	if err != nil {
		logger.GNBLog.Errorln("Error reading ngap Initial Context Setup Request Msg")
		return nil, err
	}
	ngapPdu, err = ngap.Decoder(recvMsg[:n])
	if err != nil {
		logger.GNBLog.Errorln("Error decoding ngap Initial Context Setup Request Msg")
		return nil, err
	}
	if ngapPdu.Present != ngapType.NGAPPDUPresentInitiatingMessage ||
		ngapPdu.InitiatingMessage.ProcedureCode.Value != ngapType.ProcedureCodeInitialContextSetup {
		logger.GNBLog.Errorln("Error No InitialContextSetup received")
		return nil, errors.New("Error No InitialContextSetup received")
	}

	// send ngap Initial Context Setup Response Msg
	sendMsg, err = GetInitialContextSetupResponse(ue.AmfUeNgapId, ue.RanUeNgapId)
	_, err = amfConn.Write(sendMsg)
	if err != nil {
		logger.GNBLog.Errorln("Error sending ngap Initial Context Setup Response Msg")
		return nil, err
	}

	// send NAS Registration Complete Msg
	pdu = nasTestpacket.GetRegistrationComplete(nil)
	pdu, err = EncodeNasPduWithSecurity(ue, pdu, nas.SecurityHeaderTypeIntegrityProtectedAndCiphered, true, false)
	if err != nil {
		logger.GNBLog.Errorln("Error encoding NAS Registration Complete Msg with Security")
		return nil, err
	}

	sendMsg, err = GetUplinkNASTransport(ue.AmfUeNgapId, ue.RanUeNgapId, pdu)
	if err != nil {
		logger.GNBLog.Errorln("Error building NAS Registration Complete Msg with Security")
		return nil, err
	}

	_, err = amfConn.Write(sendMsg)
	if err != nil {
		logger.GNBLog.Errorln("Error sending NAS Registration Complete Msg with Security")
		return nil, err
	}

	time.Sleep(1 * time.Second)

	return ue, nil

}

func DeRegistration(ue *RanUeContext) error {

	var n int
	var sendMsg []byte
	var recvMsg = make([]byte, 2048)

	// send NAS Deregistration Request (UE Originating)
	mobileIdentity5GS := nasType.MobileIdentity5GS{
		Len:    11, // 5g-guti
		Buffer: []uint8{0x02, 0x02, 0xf8, 0x39, 0xca, 0xfe, 0x00, 0x00, 0x00, 0x00, 0x01},
	}
	pdu := nasTestpacket.GetDeregistrationRequest(nasMessage.AccessType3GPP, 0, 0x04, mobileIdentity5GS)
	pdu, err := EncodeNasPduWithSecurity(ue, pdu, nas.SecurityHeaderTypeIntegrityProtectedAndCiphered, true, false)
	if err != nil {
		logger.GNBLog.Errorln("Error encoding Deregistration NAS with Security")
		return err
	}
	sendMsg, err = GetUplinkNASTransport(ue.AmfUeNgapId, ue.RanUeNgapId, pdu)
	if err != nil {
		logger.GNBLog.Errorln("Error building Deregistration NAS")
		return err
	}
	_, err = amfConn.Write(sendMsg)
	if err != nil {
		logger.GNBLog.Errorln("Error sending Deregistration NAS to AMF")
		return err
	}

	time.Sleep(500 * time.Millisecond)

	// receive Deregistration Accept
	n, err = amfConn.Read(recvMsg)
	if err != nil {
		logger.GNBLog.Errorln("Error reading AMF Deregistration Accept")
		return err
	}
	ngapPdu, err := ngap.Decoder(recvMsg[:n])
	if err != nil {
		logger.GNBLog.Errorln("Error decoding AMF Deregistration Accept")
		return err
	}

	if ngapPdu.Present != ngapType.NGAPPDUPresentInitiatingMessage ||
		ngapPdu.InitiatingMessage.ProcedureCode.Value != ngapType.ProcedureCodeDownlinkNASTransport {
		logger.GNBLog.Errorln("Error No DownlinkNASTransport received")
		return errors.New("Error No DownlinkNASTransport received")
	}

	nasPdu := GetNasPdu(ue, ngapPdu.InitiatingMessage.Value.DownlinkNASTransport)
	if nasPdu == nil {
		logger.GNBLog.Errorln("Error NAS PDU is nil")
		return errors.New("Error NAS PDU is nil")
	}

	if nasPdu.GmmMessage == nil {
		logger.GNBLog.Errorln("Error GMM Message is nil")
		return errors.New("Error GMM Message is nil")
	}

	if nasPdu.GmmHeader.GetMessageType() != nas.MsgTypeDeregistrationAcceptUEOriginatingDeregistration {
		logger.GNBLog.Errorln("Error Received wrong GMM message")
		return errors.New("Error Received wrong GMM message")
	}

	// receive ngap UE Context Release Command
	n, err = amfConn.Read(recvMsg)
	if err != nil {
		logger.GNBLog.Errorln("Error reading UE Context Release Command")
		return err
	}

	ngapPdu, err = ngap.Decoder(recvMsg[:n])
	if err != nil {
		logger.GNBLog.Errorln("Error decoding UE Context Release Command")
		return err
	}

	if ngapPdu.Present != ngapType.NGAPPDUPresentInitiatingMessage ||
		ngapPdu.InitiatingMessage.ProcedureCode.Value != ngapType.ProcedureCodeUEContextRelease {
		logger.GNBLog.Errorln("Error No UEContextReleaseCommand received")
		return errors.New("Error No UEContextReleaseCommand received")
	}

	// send ngap UE Context Release Complete
	sendMsg, err = GetUEContextReleaseComplete(ue.AmfUeNgapId, ue.RanUeNgapId, nil)
	if err != nil {
		logger.GNBLog.Errorln("Error building ngap UE Context Release Complete")
		return err
	}

	_, err = amfConn.Write(sendMsg)
	if err != nil {
		logger.GNBLog.Errorln("Error sending ngap UE Context Release Complete")
		return err
	}

	time.Sleep(time.Second * 1)
	return nil

}

func PDUSessionRequest(ue *RanUeContext, snssai string, sessionId uint8, dnn string) error {

	fmt.Println(snssai)
	fmt.Println(sessionId)
	fmt.Println(dnn)

	var n int
	var sendMsg []byte
	var recvMsg = make([]byte, 2048)

	err := checkAmfConnection()
	if err != nil {
		logger.GNBLog.Errorln("Error connecting to the AMF")
		return err
	}

	sNssai := *convertSnssai(snssai)

	pdu := nasTestpacket.GetUlNasTransport_PduSessionEstablishmentRequest(sessionId, nasMessage.ULNASTransportRequestTypeInitialRequest, dnn, &sNssai)
	pdu, err = EncodeNasPduWithSecurity(ue, pdu, nas.SecurityHeaderTypeIntegrityProtectedAndCiphered, true, false)
	if err != nil {
		logger.GNBLog.Errorln("Error encoding PduSessionEstablishmentRequest")
		return err
	}

	sendMsg, err = GetUplinkNASTransport(ue.AmfUeNgapId, ue.RanUeNgapId, pdu)
	_, err = amfConn.Write(sendMsg)
	if err != nil {
		logger.GNBLog.Errorln("Error sending PduSessionEstablishmentRequest")
		return err
	}

	// receive 12. NGAP-PDU Session Resource Setup Request(DL nas transport((NAS msg-PDU session setup Accept)))
	n, err = amfConn.Read(recvMsg)
	ngapPdu, err := ngap.Decoder(recvMsg[:n])
	if err != nil {
		logger.GNBLog.Errorln("Error decoding NGAP-PDU Session Resource Setup Request")
		return err
	}
	if ngapPdu.Present != ngapType.NGAPPDUPresentInitiatingMessage ||
		ngapPdu.InitiatingMessage.ProcedureCode.Value != ngapType.ProcedureCodePDUSessionResourceSetup {
		logger.GNBLog.Errorln("Error No PDUSessionResourceSetup received")
		return err
	}

	// send 14. NGAP-PDU Session Resource Setup Response
	sendMsg, err = GetPDUSessionResourceSetupResponse(ue.AmfUeNgapId, ue.RanUeNgapId, APIConfig.Configuration.NGRANInterface.IPv4Addr)
	if err != nil {
		logger.GNBLog.Errorln("Error encoding NGAP-PDU Session Resource Setup Response")
		return err
	}
	_, err = amfConn.Write(sendMsg)
	if err != nil {
		logger.GNBLog.Errorln("Error decoding NGAP-PDU Session Resource Setup Response")
		return err
	}

	time.Sleep(time.Second * 1)
	return nil

}

func PDUSessionRelease(ue *RanUeContext, snssai string, sessionId uint8, dnn string) error {

	var sendMsg []byte

	err := checkAmfConnection()
	if err != nil {
		logger.GNBLog.Errorln("Error connecting to the AMF")
		return err
	}

	sNssai := *convertSnssai(snssai)

	// Send Pdu Session Establishment Release Request
	pdu := nasTestpacket.GetUlNasTransport_PduSessionReleaseRequest(sessionId)
	pdu, err = EncodeNasPduWithSecurity(ue, pdu, nas.SecurityHeaderTypeIntegrityProtectedAndCiphered, true, false)
	if err != nil {
		logger.GNBLog.Errorln("Error encoding NPdu Session Establishment Release Request")
		return err
	}
	sendMsg, err = GetUplinkNASTransport(ue.AmfUeNgapId, ue.RanUeNgapId, pdu)
	if err != nil {
		logger.GNBLog.Errorln("Error building NPdu Session Establishment Release Request")
		return err
	}
	_, err = amfConn.Write(sendMsg)
	if err != nil {
		logger.GNBLog.Errorln("Error sending NPdu Session Establishment Release Request")
		return err
	}

	time.Sleep(1000 * time.Millisecond)
	// send N2 Resource Release Ack(PDUSession Resource Release Response
	sendMsg, err = GetPDUSessionResourceReleaseResponse(ue.AmfUeNgapId, ue.RanUeNgapId)
	if err != nil {
		logger.GNBLog.Errorln("Error building N2 Resource Release Ack(PDUSession Resource Release Response")
		return err
	}
	_, err = amfConn.Write(sendMsg)
	if err != nil {
		logger.GNBLog.Errorln("Error sending N2 Resource Release Ack(PDUSession Resource Release Response")
		return err
	}

	// wait 10 ms
	time.Sleep(1000 * time.Millisecond)

	//send N1 PDU Session Release Ack PDU session release complete
	pdu = nasTestpacket.GetUlNasTransport_PduSessionReleaseComplete(sessionId, nasMessage.ULNASTransportRequestTypeInitialRequest, dnn, &sNssai)
	pdu, err = EncodeNasPduWithSecurity(ue, pdu, nas.SecurityHeaderTypeIntegrityProtectedAndCiphered, true, false)
	if err != nil {
		logger.GNBLog.Errorln("Error encoding N1 PDU Session Release Ack PDU session release complete")
		return err
	}
	sendMsg, err = GetUplinkNASTransport(ue.AmfUeNgapId, ue.RanUeNgapId, pdu)
	if err != nil {
		logger.GNBLog.Errorln("Error building N1 PDU Session Release Ack PDU session release complete")
		return err
	}
	_, err = amfConn.Write(sendMsg)
	if err != nil {
		logger.GNBLog.Errorln("Error sending N1 PDU Session Release Ack PDU session release complete")
		return err
	}

	return nil

}
