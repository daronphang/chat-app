package cgroup2

import (
	"strconv"

	linuxproc "github.com/c9s/goprocinfo/linux"
)

func GetContainerMemUsedNoCache() (uint64, error) {
	// Need to subtract cache (inactive_file) from total usage.
	// https://github.com/docker/cli/blob/d47c36debb4221194b98fa0f72573669e6ac0bc7/cli/command/container/stats_helpers.go#L105
	strMemUsed, err := extractStatValue("/sys/fs/cgroup/memory.current", "")
	if err != nil {
		return 0, err
	}
	memUsed, err := strconv.ParseInt(strMemUsed, 10, 64)
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
	return uint64(memUsed - cache), nil
}

func GetSystemMemUsed() (uint64, error) {
	stat, err := linuxproc.ReadMemInfo("/proc/meminfo")
	if err != nil {
		return 0, err
	}
	memUsed := stat.MemTotal - stat.MemFree
	return memUsed, nil
}

func GetContainerMemLimit() (uint64, error) {
	strMemLimit, err := extractStatValue("/sys/fs/cgroup/memory.max", "")
	if err != nil {
		return 0, err
	}
	memLimit, err := strconv.ParseInt(strMemLimit, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint64(memLimit), nil
}

func CalculateMemUsage(stat *ContainerStat) float64 {
	return float64(stat.MemUsed) / float64(stat.MemLimit)
}