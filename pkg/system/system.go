// Copyright © 2022 Alibaba Group Holding Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package system

import (
	"runtime"
	"strings"

	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

type InfoStat struct {
	Host   HostInfo
	Memory MemoryInfo
	Net    []Network
	Disk   []DiskInfo
}

type HostInfo struct {
	Arch                 string
	HostName             string
	OS                   string // ex: freebsd, linux
	OSDistribution       string // ex: ubuntu, centos
	OSDistributionFamily string // ex: debian, rhel
	OSVersion            string // version of the complete OS
	Kernel               string // "uname -r"
	NumCPU               int
}

type Network struct {
	HardwareAddr string
	IPAddress    []string
	Name         string
}

type DiskInfo struct {
	Name       string
	FsType     string
	MountPoint string
}

type MemoryInfo struct {
	// Total amount of RAM on this system,Kib size.
	Total       uint64
	Available   uint64
	Used        uint64
	UsedPercent float64
}

func Info() (*InfoStat, error) {
	hostInfo, err := GetHostInfo()
	if err != nil {
		return nil, err
	}

	diskInfo, err := GetDiskInfo()
	if err != nil {
		return nil, err
	}

	netInfo, err := GetNetInfo()
	if err != nil {
		return nil, err
	}

	memInfo, err := GetMemoryInfo()
	if err != nil {
		return nil, err
	}
	return &InfoStat{
		Host:   hostInfo,
		Disk:   diskInfo,
		Net:    netInfo,
		Memory: memInfo,
	}, nil
}

func GetHostInfo() (HostInfo, error) {
	ret, err := host.Info()
	if err != nil {
		return HostInfo{}, err
	}
	return HostInfo{
		NumCPU:               runtime.NumCPU(),
		OS:                   ret.OS,
		Arch:                 NormalizeArch(ret.KernelArch),
		HostName:             ret.Hostname,
		Kernel:               ret.KernelVersion,
		OSDistribution:       ret.Platform,
		OSDistributionFamily: ret.PlatformFamily,
		OSVersion:            ret.PlatformVersion,
	}, nil
}

func GetDiskInfo() ([]DiskInfo, error) {
	var di []DiskInfo

	ps, err := disk.Partitions(false)
	if err != nil {
		return nil, err
	}
	for _, p := range ps {
		di = append(di, DiskInfo{
			Name:       p.Device,
			FsType:     p.Fstype,
			MountPoint: p.Mountpoint,
		})
	}
	return di, nil
}

func GetNetInfo() ([]Network, error) {
	var netInfo []Network
	infos, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	for _, i := range infos {
		if i.HardwareAddr == "" {
			continue
		}
		var ips []string
		for _, ip := range i.Addrs {
			ips = append(ips, ip.Addr)
		}

		netInfo = append(netInfo, Network{
			Name:         i.Name,
			HardwareAddr: i.HardwareAddr,
			IPAddress:    ips,
		})
	}
	return netInfo, nil
}

func GetMemoryInfo() (MemoryInfo, error) {
	v, err := mem.VirtualMemory()
	if err != nil {
		return MemoryInfo{}, err
	}
	return MemoryInfo{
		UsedPercent: v.UsedPercent,
		Used:        v.Used / 1024,
		Available:   v.Available / 1024,
		Total:       v.Total / 1024,
	}, nil
}

func NormalizeArch(arch string) string {
	arch = strings.ToLower(arch)
	switch arch {
	case "i386":
		arch = "386"
	case "x86_64", "x86-64":
		arch = "amd64"
	case "aarch64", "arm64":
		arch = "arm64"
	case "armhf":
		arch = "arm"
	case "armel":
		arch = "arm"
	}

	return arch
}
