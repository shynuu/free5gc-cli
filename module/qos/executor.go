package qos

import (
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

// mark --set-dscp af12 --destination-ip 172.16.10.3 --source-ip 172.16.10.3 --teid 000 --protocol tcp --destination-port 23 --source-port 23

func executorMark(in string) {
	cmd := strings.Split(strings.TrimSpace(in), " ")
	var sourcePort uint16 = 0
	var matchSourcePort bool = false

	l := len(cmd)
	if l < 2 {
		return
	}

	if l > 15 {
		return
	}

	logger.GNBLog.Infoln("Generating the U32 command")

	dscps := cmd[2]
	dscp := QOSMAP[dscps]
	destinationIP := cmd[4]
	sourceIP := cmd[6]
	teids := cmd[8]
	teid, err := stringToUint32(teids)
	if err != nil {
		logger.GNBLog.Errorln("Impossible to convert TEID")
		return
	}

	protocols := cmd[10]
	protocol := PROTOMAP[protocols]
	destinationPorts := cmd[12]
	destinationPort, err := stringToUint16(destinationPorts)
	if err != nil {
		logger.GNBLog.Errorln("Impossible to convert destination port")
		return
	}

	if l > 13 {
		sourcePorts := cmd[14]
		sourcePort, err = stringToUint16(sourcePorts)
		if err != nil {
			logger.GNBLog.Errorln("Impossible to convert source port")
			return
		}
		matchSourcePort = true
	}

	var packet = []u32.Protocol{
		&u32.IPV4Header{
			Source:      sourceIP,
			Destination: destinationIP,
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
			TEID: teid,
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
	U32.BuildCommand()

}

func executorConfiguration(in string) {
	s := strings.TrimSpace(in)
	if s == "configuration reload" {
		Reload()
		return
	}
}

// Executor parse CLI
func Executor(in string) {

	if strings.HasPrefix(in, "configuration") {
		executorConfiguration(in)
	}

	if strings.HasPrefix(in, "mark") {
		executorMark(in)
	}

	return

}
