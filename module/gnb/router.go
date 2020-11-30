package gnb

import (
	"free5gc-cli/logger"
	"os"
	"os/exec"

	"github.com/songgao/water"
)

// BUFFERSIZE is
const BUFFERSIZE = 1500

// MTU is 1500 - IPV4 - UDP - GTP
const MTU = "1400"

func runIP(args ...string) {
	cmd := exec.Command("/sbin/ip", args...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	if nil != err {
		logger.GNBLog.Errorln("Error running /sbin/ip:", err)
	}
}

func InitTUN() {

	config := water.Config{
		DeviceType: water.TUN,
	}
	config.Name = GNBConfig.Configuration.TUN

	iface, err := water.New(config)
	if nil != err {
		logger.GNBLog.Errorln("Unable to allocate TUN interface:", err)
	}

	logger.GNBLog.Infoln("Interface allocated:", iface.Name())
	// set interface parameters
	runIP("sudo", "link", "set", "dev", iface.Name(), "mtu", MTU)
	runIP("sudo", "addr", "add", GNBConfig.Configuration.GTPInterface.Ipv4, "dev", iface.Name())
	runIP("sudo", "link", "set", "dev", iface.Name(), "up")

}
