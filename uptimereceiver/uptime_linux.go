//go:build linux
// +build linux

package uptimereceiver

import (
	"golang.org/x/sys/unix"
	"time"
)

func getUptime() (time.Duration, error) {
	var info unix.Sysinfo_t
	if err := unix.Sysinfo(&info); err != nil {
		return time.Duration(0), err
	}
	return time.Duration(info.Uptime) * time.Second, nil
}
