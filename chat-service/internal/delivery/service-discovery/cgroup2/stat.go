package cgroup2

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

var (
	errInvalidStat = errors.New("invalid stat")
	errMissingStat = errors.New("stat value is missing")
)

type ContainerStat struct {
	MemLimit 		uint64 
	MemUsage 		uint64
	TotalSystemCPU	uint64 	// Jiffies.
	CPUUsage 		uint64	// Jiffies.
	CPULimit 		uint64
}

func GetContainerStat() (*ContainerStat, error) {
	containerStat := &ContainerStat{}

	// CPU.
	totalSystemCPU, err := getSystemTotalCPU()
	if err != nil {
		return nil, err
	}
	containerStat.TotalSystemCPU = totalSystemCPU

	CPULimit, err := getContainerCPULimit(totalSystemCPU)
	if err != nil {
		return nil, err
	}
	containerStat.CPULimit = CPULimit

	CPUUsage, err := getContainerCPUUsage()
	if err != nil {
		return nil, err
	}
	containerStat.CPUUsage = CPUUsage

	// Memory.
	memUsage, err := getContainerMemUsageNoCache()
	if err != nil {
		return nil, err
	}
	containerStat.MemUsage = memUsage

	strMemLimit, err := extractStatValue("/sys/fs/cgroup/memory.max", "")
	if err != nil {
		return nil, err
	}
	memLimit, err := strconv.ParseInt(strMemLimit, 10, 64)
	if err != nil {
		return nil, err
	}
	containerStat.MemLimit = uint64(memLimit)

	return containerStat, nil
}

func extractStatValue(filePath string, key string) (string, error) {
	/*
	inactive_anon 4878336
	active_anon 5726208
	inactive_file 2113536
	*/

	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var line string
	for scanner.Scan() {
		line = scanner.Text()

		// If key is empty, to return the first line.
		if key == "" {
			return line, nil
		}

		if strings.Contains(line, key) {
			values := strings.Split(line, " ")
			if len(values) != 2 {
				return "", errInvalidStat
			}
			return values[1], nil
		}
	}
	return "", errMissingStat
}