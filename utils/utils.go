package utils

import (
	"github.com/shirou/gopsutil/v3/host"
)

type HostInfo struct {
	Os     string
	Distro string
	Arch   string
}

func GetHostInfo() HostInfo {
	info, _ := host.Info()

	return HostInfo{
		Os:     info.OS,
		Distro: info.Platform,
		Arch:   info.KernelArch,
	}
}
