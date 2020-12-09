package gnb

import (
	"fmt"
	"free5gc-cli/lib/u32"
	"free5gc-cli/logger"
	"strconv"
	"strings"
)

func stringToUint32(number string) (uint32, error) {
	conv, err := strconv.Atoi(number)
	if err == nil {
		return uint32(conv), nil
	}
	return 0, err
}

func stringToUint16(number string) (uint16, error) {
	conv, err := strconv.Atoi(number)
	if err == nil {
		return uint16(conv), nil
	}
	return 0, err
}

func stringToUint8(number string) (uint8, error) {
	conv, err := strconv.Atoi(number)
	if err == nil {
		return uint8(conv), nil
	}
	return 0, err
}

func executorConfiguration(in string) {
	s := strings.TrimSpace(in)
	if s == "configuration reload" {
		Reload()
		return
	}
}

func executorUE(in string) {

	cmd := strings.Split(strings.TrimSpace(in), " ")
	l := len(cmd)

	if l < 2 || l > 4 {
		return
	}

	first := cmd[1]

	if first == "register" {
		if l < 4 {
			return
		}
		u := cmd[3]
		ueInfo := strings.Split(u, "/")
		logger.GNBLog.Infoln(fmt.Sprintf("Registering user %s on the network", ueInfo[0]))
		err := gnb.Register(ueInfo[0])
		if err != nil {
			return
		}
		logger.GNBLog.Infoln(fmt.Sprintf("Successfully register user %s on the network", ueInfo[0]))
		return
	}

	if first == "deregister" {
		if l < 4 {
			return
		}
		ue := cmd[3]
		logger.GNBLog.Infoln(fmt.Sprintf("De-Registering user %s on the network", ue))
		err := gnb.Deregister(ue)
		if err != nil {
			return
		}
		logger.GNBLog.Infoln(fmt.Sprintf("Successfully de-register user %s on the network", ue))
		return
	}

	return

}

func executorPDUSession(in string) {

	cmd := strings.Split(strings.TrimSpace(in), " ")
	l := len(cmd)

	if l < 2 {
		return
	}

	first := cmd[1]

	if first == "request" && l > 7 {
		logger.GNBLog.Infoln(fmt.Sprintf("Establishing PDU Session for user %s with snssai %s and dnn %s", cmd[3], cmd[5], cmd[7]))
		err := gnb.PDURequest(cmd[3], cmd[5], cmd[7])
		if err != nil {
			return
		}
		logger.GNBLog.Infoln(fmt.Sprintf("Successfully Established PDU Session for user %s with snssai %s and dnn %s", cmd[3], cmd[5], cmd[7]))
	}

	if first == "release" && l > 3 {
		cmd = strings.Split(cmd[3], "-")
		logger.GNBLog.Infoln(fmt.Sprintf("Releasing PDU Session for user %s with session %s", cmd[0], cmd[1]))
		err := gnb.PDURelease(cmd[0], cmd[1])
		if err != nil {
			return
		}
		logger.GNBLog.Infoln(fmt.Sprintf("Successfully Releasing PDU Session for user %s with session %s", cmd[0], cmd[1]))
	}

	if first == "qos" {
		executorQOS(in)
	}

	return

}

func executorQOS(in string) {

	// pdu-sessions qos add --set-dscp 12 --session imsmi/24 --protocol tcp/udp --destination-port 23 --source-port 23
	cmd := strings.Split(strings.TrimSpace(in), " ")
	l := len(cmd)

	second := cmd[2]

	if second == "add" {

		var sourcePort uint16 = 0
		var matchSourcePort bool = false
		if l > 13 {
			return
		}

		dscps := cmd[4]
		dscp := QOSMAP[dscps]

		split := strings.Split(cmd[6], "/")
		session, err := stringToUint8(split[1])
		if err != nil {
			logger.GNBLog.Errorln("Impossible to convert session")
			return
		}
		pduSession, err := gnb.GetPDUSession(split[0], session)
		if err != nil {
			logger.GNBLog.Errorln("Impossible to find PDU Session")
			return
		}

		protocols := cmd[8]
		protocol := PROTOMAP[protocols]

		destinationPorts := cmd[10]
		destinationPort, err := stringToUint16(destinationPorts)
		if err != nil {
			logger.GNBLog.Errorln("Impossible to convert destination port")
			return
		}

		if l > 11 {
			sourcePorts := cmd[12]
			sourcePort, err = stringToUint16(sourcePorts)
			if err != nil {
				logger.GNBLog.Errorln("Impossible to convert source port")
				return
			}
			matchSourcePort = true
		}

		var packet = []u32.Protocol{
			&u32.IPV4Header{
				Source:      GNBConfig.Configuration.GTPInterface.Ipv4,
				Destination: GNBConfig.Configuration.UpfInterface.IPv4Addr,
				Protocol:    u32.PROTO_UDP,
				Set: &u32.IPV4Fields{
					Source:      true,
					Destination: true,
					Protocol:    true,
				},
			},
			&u32.UDPHeader{
				SourcePort:      2152,
				DestinationPort: 2152,
				Set: &u32.UDPFields{
					SourcePort:      true,
					DestinationPort: true,
				},
			},
			&u32.GTPv1Header{
				TEID: pduSession.TEID,
				Set: &u32.GTPv1Fields{
					TEID: true,
				},
			},
		}

		var packet2 = []u32.Protocol{}
		if protocol == u32.PROTO_UDP {

			packet2 = []u32.Protocol{
				&u32.IPV4Header{
					Protocol: protocol,
					Set: &u32.IPV4Fields{
						Protocol: true,
					},
				},
				&u32.UDPHeader{
					SourcePort:      sourcePort,
					DestinationPort: destinationPort,
					Set: &u32.UDPFields{
						SourcePort:      matchSourcePort,
						DestinationPort: true,
					},
				},
			}
		} else {
			packet2 = []u32.Protocol{
				&u32.IPV4Header{
					Protocol: protocol,
					Set: &u32.IPV4Fields{
						Protocol: true,
					},
				},
				&u32.TCPHeader{
					SourcePort:      sourcePort,
					DestinationPort: destinationPort,
					Set: &u32.TCPFields{
						SourcePort:      matchSourcePort,
						DestinationPort: true,
					},
				},
			}
		}

		packet = append(packet, packet2...)

		var U32 = u32.NewU32(&packet, dscp)
		err = U32.Run()
		if err != nil {
			fmt.Println(err)
			logger.GNBLog.Errorln("Impossible to add the QoS rule to iptables")
			return
		}
		logger.GNBLog.Infoln(fmt.Sprintf("Successfully added QoS rule to for pdu session %d", session))
		return
	}

	if second == "flush" {
		var U32 = &u32.U32{}
		err := U32.Flush()
		if err != nil {
			logger.GNBLog.Errorln("Error flushing QoS Rules table")
		}
		logger.GNBLog.Infoln("Successfully flushed QoS Rules")
		return
	}

}

// Executor parse CLI
func Executor(in string) {

	if strings.HasPrefix(in, "configuration") {
		executorConfiguration(in)
	}

	if strings.HasPrefix(in, "user") {
		executorUE(in)
	}

	if strings.HasPrefix(in, "pdu-session") {
		executorPDUSession(in)
	}

}
