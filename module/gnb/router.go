package gnb

import (
	"fmt"
	"free5gc-cli/logger"
	"net"
	"os"
	"os/exec"
	"sync"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/songgao/water"
)

// BUFFERSIZE of the packet
const BUFFERSIZE = 1500

// MTU is 1500 - IPV4 - UDP - GTP
const MTU = "1400"

func runIP(args ...string) {
	cmd := exec.Command("ip", args...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	if nil != err {
		logger.GNBLog.Errorln("Error running ip command:", err)
	}
	return
}

var gtpRouter *GTPRouter

// GTPRouter provides the functionnality to encapsulate and desencapsulate packet using GTP protocol
type GTPRouter struct {
	GNB        *GNB
	UpfConn    *net.UDPConn
	Iface      *water.Interface
	IfaceMutex *sync.Mutex
	UpfAddress *net.UDPAddr
}

// NewRouter build a new router
func NewRouter(upfIP string, upfPort int, gnbIP string, gnbPort int, subnet string, gnb *GNB) (*GTPRouter, error) {

	config := water.Config{
		DeviceType: water.TUN,
	}
	config.Name = GNBConfig.Configuration.TUN

	iface, err := water.New(config)
	if nil != err {
		logger.GNBLog.Errorln("Unable to allocate TUN interface:", err)
	}

	logger.GNBLog.Infoln("TUN Interface allocated:", iface.Name())
	// set interface parameters
	runIP("link", "set", "dev", iface.Name(), "mtu", MTU)
	runIP("addr", "add", GNBConfig.Configuration.GTPInterface.Ipv4, "dev", iface.Name())
	runIP("link", "set", "dev", iface.Name(), "up")

	var GNBAddr = net.UDPAddr{IP: net.ParseIP(gnbIP), Port: gnbPort}
	// var UPFAddr = net.UDPAddr{IP: net.ParseIP(upfIP), Port: upfPort}

	// Connect to the UPF
	upfAddress, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", upfIP, upfPort))
	if err != nil {
		logger.GNBLog.Errorln("Impossible to resolve UPF address")
		logger.GNBLog.Errorln(err)
		return nil, err
	}
	upfConn, err := net.ListenUDP("udp", &GNBAddr)
	if err != nil {
		logger.GNBLog.Errorln("Impossible to Dial UPF")
		return nil, err
	}

	runIP("route", "add", fmt.Sprintf("%s/16", subnet), "via", gnbIP)

	var m1 sync.Mutex

	var gtpRouter = GTPRouter{
		GNB:        gnb,
		UpfConn:    upfConn,
		Iface:      iface,
		IfaceMutex: &m1,
		UpfAddress: upfAddress,
	}
	return &gtpRouter, nil

}

// Close the connection with the UPF and Tun interface
func (r *GTPRouter) Close() {
	r.UpfConn.Close()
	r.Iface.Close()
}

// Encapsulate the packet using GTP protocol
func (r *GTPRouter) Encapsulate() {

	// Read the incoming packet on the tun interface
	// Encapsulate the packet with GTP
	// Write it to the socket with the UPF
	packet := make([]byte, BUFFERSIZE)
	var ipv4 layers.IPv4
	var gtp layers.GTPv1U
	buf := gopacket.NewSerializeBuffer()
	opts := gopacket.SerializeOptions{} // See SerializeOptions for more details.
	parser := gopacket.NewDecodingLayerParser(layers.LayerTypeIPv4, &ipv4)
	decoded := []gopacket.LayerType{}

	for {
		// read the packet coming from the TUN interface
		n, err := r.Iface.Read(packet)
		if err != nil {
			logger.GNBLog.Errorln("Error reading the TUN interface input")
			panic("Impossible to read the TUN interface")
		}
		// build the ipv4 header
		err = parser.DecodeLayers(packet[:n], &decoded)
		if len(decoded) > 0 {
			// find the teid
			teid, err := r.GNB.GetTEID(ipv4.SrcIP)
			if err == nil {
				gtp = layers.GTPv1U{
					Version:       0x01,
					TEID:          teid,
					MessageType:   0xFF,
					MessageLength: uint16(n),
				}
				err = gtp.SerializeTo(buf, opts)
				if err != nil {
					logger.GNBLog.Errorln("Error Serializing the packet Layers")
					break
				}
				pkt := append(buf.Bytes(), packet[:n]...)
				n, err = r.UpfConn.WriteToUDP(pkt, r.UpfAddress)
				buf.Clear()
			}
		}
	}

}

// Desencapsulate the GTP packet: remove the GTP headers and route the packet
func (r *GTPRouter) Desencapsulate() {

	// Read the packet coming from the socket
	// Desencapsulate the packet and remove GTP Header
	// Write the answer to the TUN interface

	buf := make([]byte, BUFFERSIZE)
	var gtp layers.GTPv1U
	parser := gopacket.NewDecodingLayerParser(layers.LayerTypeGTPv1U, &gtp)
	decoded := []gopacket.LayerType{}

	for {
		n, err := r.UpfConn.Read(buf)
		if err != nil {
			logger.GNBLog.Errorln("Error reading the UPF incoming packet")
			panic("Error reading the UPF incoming packet")
		}

		err = parser.DecodeLayers(buf[:n], &decoded)
		if len(decoded) > 0 {
			r.Iface.Write(gtp.LayerPayload())
		}

	}

}
