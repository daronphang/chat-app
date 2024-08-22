package cgroup2

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

var (
	errInvalidStat = errors.New("invalid stat")
	errMissingStat = errors.New("stat value is missing")
)

type ContainerStat struct {
	MemLimit 		uint64 
	MemUsed 		uint64
	SystemCPUUsed	uint64	// Jiffies.
	CPUUsed 		uint64	// Jiffies.
}

func GetContainerStat() (*ContainerStat, error) {
	stat := &ContainerStat{}

	systemCPUUsed, err := GetSystemCPUUsed()
	if err != nil {
		return nil, err
	}
	stat.SystemCPUUsed = systemCPUUsed

	CPUUsed, err := GetContainerCPUUsed()
	if err != nil {
		return nil, err
	}
	stat.CPUUsed = CPUUsed

	memUsed, err := GetContainerMemUsedNoCache()
	if err != nil {
		return nil, err
	}
	stat.MemUsed = memUsed

	memLimit, err := GetContainerMemLimit()
	if err != nil {
		return nil, err
	}
	stat.MemLimit = memLimit
	
	return stat, nil
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