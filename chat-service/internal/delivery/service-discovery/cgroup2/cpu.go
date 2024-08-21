package cgroup2

import (
	"strconv"
	"strings"

	linuxproc "github.com/c9s/goprocinfo/linux"
)

func getSystemTotalCPU() (uint64, error) {
	/*
	cat /proc/stat 
	cpu 1279636934 73759586 192327563 12184330186 543227057 56603 68503253 0 0 
	cpu0 297522664 8968710 49227610 418508635 72446546 56602 24904144 0 0 
	cpu1 227756034 9239849 30760881 424439349 196694821 0 7517172 0 0 
	cpu2 86902920 6411506 12412331 769921453 17877927 0 4809331 0 0 

	Different OS will output different number of columns.

	First line is an aggregate of all CPU cores.

	CPU usage includes only the time spent by CPU cycles in running 
	processes or handling interrupts. When any process is waiting for I/O
	to complete, it is not contributing to CPU cycles. Hence, it is added
	to the idle time instead.
	*/


	// for idx, s := range stat.CPUStats {
	// 	if idx == 0 {
	// 		continue
	// 	}
	// 	idle += int(s.Idle) + int(s.IOWait)
	// 	userTime := int(s.User) - int(s.Guest)
	// 	niceTime := int(s.Nice) + int(s.GuestNice)
	// 	total += userTime + niceTime + int(s.System) + int(s.Steal) + int(s.IRQ) + int(s.SoftIRQ) + int(s.Idle) + int(s.IOWait)
	// }

	stat, err := linuxproc.ReadStat("/proc/stat")
	if err != nil {
		return 0, err
	}

	aggIdle := stat.CPUStats[0].Idle + stat.CPUStats[0].IOWait
	aggUserTime := stat.CPUStats[0].User - stat.CPUStats[0].Guest
	aggNiceTime := stat.CPUStats[0].Nice - stat.CPUStats[0].GuestNice
	aggTotal := aggUserTime + aggNiceTime + stat.CPUStats[0].System + stat.CPUStats[0].Steal + stat.CPUStats[0].IRQ + stat.CPUStats[0].SoftIRQ + aggIdle
	// aggCPUUsage := float64(1 - (float64(aggIdle) / float64(aggTotal)))
	return aggTotal, nil
}

func getContainerCPULimit(totalSystemCPU uint64) (uint64, error) {
	stat, err := linuxproc.ReadStat("/proc/stat")
	if err != nil {
		return 0, err
	}
	systemCores := len(stat.CPUStats)

	strContainerLimit, err := extractStatValue("/sys/fs/cgroup/cpu.max", "")
	if err != nil {
		return 0, err
	}
	temp := strings.Split(strContainerLimit, " ")
	if len(temp) != 2 {
		return 0, errInvalidStat
	}

	containerLimitRef, err := strconv.Atoi(temp[0])
	if err != nil {
		return 0, err
	}

	singleCoreRef, err := strconv.Atoi(temp[1])
	if err != nil {
		return 0, err
	}

	containerCores := float32(containerLimitRef) / float32(singleCoreRef)
	singleCPUCore := float32(totalSystemCPU / uint64(systemCores))
	return uint64(containerCores * singleCPUCore), nil
}

func getContainerCPUUsage() (uint64, error) {
	/*
	/sys/fs/cgroup/cpu.stat
	usage_usec 22498068	(milliseconds)
	user_usec 11991789
	system_usec 10506279
	nr_periods 18781
	nr_throttled 92
	throttled_usec 8822054
	nr_bursts 0
	burst_usec 0

	As CPU in /proc/stat is reported in jiffies (1/100th of second, or 10 milliseconds),
	need to convert it into jiffies i.e. divide by 10.
	*/
	strCPUUsage, err := extractStatValue("/sys/fs/cgroup/cpu.stat", "usage_usec")
	if err != nil {
		return 0, err
	}
	CPUUsage, err := strconv.ParseInt(strCPUUsage, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint64(CPUUsage) / uint64(10), nil
}