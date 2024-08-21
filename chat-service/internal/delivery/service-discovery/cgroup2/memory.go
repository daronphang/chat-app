package cgroup2

import (
	"strconv"

	linuxproc "github.com/c9s/goprocinfo/linux"
)

func getContainerMemUsageNoCache() (uint64, error) {
	// Need to subtract cache (inactive_file) from total usage.
	// https://github.com/docker/cli/blob/d47c36debb4221194b98fa0f72573669e6ac0bc7/cli/command/container/stats_helpers.go#L105
	strMemUsage, err := extractStatValue("/sys/fs/cgroup/memory.current", "")
	if err != nil {
		return 0, err
	}
	memUsage, err := strconv.ParseInt(strMemUsage, 10, 64)
	if err != nil {
		return 0, err
	}

	strCache, err := extractStatValue("/sys/fs/cgroup/memory.stat", "inactive_file")
	if err != nil {
		return 0, err
	}
	cache, err := strconv.ParseInt(strCache, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint64(memUsage - cache), nil
}

func GetSystemMemUsage() (float64, error) {
	stat, err := linuxproc.ReadMemInfo("/proc/meminfo")
	if err != nil {
		return 0, err
	}
	memUsage := float64(1 - (float64(stat.MemFree) / float64(stat.MemTotal)))
	return memUsage, nil
}