package u32

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func runIptables(args ...string) error {
	cmd := exec.Command("iptables", args...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	err := cmd.Run()
	if nil != err {
		return err
	}
	return nil
}

// U32 struct
type U32 struct {
	Protocols []Protocol
	Length    int
	Matches   string
	DSCP      uint8
}

// BuildPacket build Packet Headers
func (u32 *U32) BuildPacket() {
	offset := &Offset{Offset: 0, U32Offset: ""}
	for i := 0; i < len(u32.Protocols); i++ {
		u32.Protocols[i].SetOffset(offset)
		u32.Protocols[i].MoveOffset(offset)
	}
}

// BuildMatches build Matches
func (u32 *U32) BuildMatches() string {
	var matches []string
	for i := 0; i < len(u32.Protocols); i++ {
		match := u32.Protocols[i].BuildMatches()
		if match != "" {
			matches = append(matches, match)
		}
	}
	u32.Matches = strings.Join(matches, " && ")
	return u32.Matches
}

// Run run the iptables command and add the rule to mangle/POSTROUTING
func (u32 *U32) Run() error {
	cmd := fmt.Sprintf(`"%s"`, u32.BuildMatches())
	dscp := fmt.Sprintf("%d", u32.DSCP)
	fmt.Println(cmd, dscp)
	err := runIptables("-t", "mangle", "-A", "POSTROUTING", "-m", "u32", "--u32", cmd, "-j", "DSCP", "--set-dscp", dscp)
	return err
}

// Flush delete all the rules of the mangle table
func (u32 *U32) Flush() error {
	err := runIptables("-t", "mangle", "-F")
	return err
}

// NewU32 returns a new U32 Struct
func NewU32(protocols *[]Protocol, dscp uint8) *U32 {
	var u32 = U32{Protocols: *protocols, DSCP: dscp}
	u32.BuildPacket()
	return &u32
}
