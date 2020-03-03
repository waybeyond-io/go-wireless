package wireless

import (
	"os"
	"path"
	"path/filepath"
)

// Interfaces is a shortcut to the best known method for gathering the wireless
// interfaces from the current system
var Interfaces = SysFSInterfaces

// DefaultInterface will return the default wireless interface, being the first
// one returned from the Interfaces method
func DefaultInterface() (string, bool) {
	ifs := Interfaces()
	if len(ifs) == 0 {
		return "", false
	}

	return ifs[0], true
}

// SysFSInterfaces returns the wireless interfaces found in the SysFS (/sys/class/net)
func SysFSInterfaces() []string {
	s := []string{}
	base := "/sys/class/net"
	matches, _ := filepath.Glob(path.Join(base, "*"))

	//  look for the wireless folder in each interfces directory to determine if it is a wireless device
	for _, iface := range matches {
		if _, err := os.Stat(path.Join(iface, "wireless")); err == nil {
			s = append(s, path.Base(iface))
		}
	}

	return s
}
