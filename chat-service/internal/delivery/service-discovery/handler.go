package svcdis

import (
	"chat-service/internal"
	"chat-service/internal/domain"
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/url"
	"sync"
	"time"

	linuxproc "github.com/c9s/goprocinfo/linux"
	"github.com/google/uuid"
	etcd "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
)


var (
	syncOnceConfig sync.Once
	hostIPAddress net.IP
	logger, _ = internal.WireLogger()
)

func getOutboundIP() (net.IP, error) {
	var e error
	syncOnceConfig.Do(func() {
		conn, err := net.Dial("udp", "8.8.8.8:80")
		if err != nil {
			e = err
			return
		}
		defer conn.Close()
	
		localAddr := conn.LocalAddr().(*net.UDPAddr)
		hostIPAddress = localAddr.IP
	})
    return hostIPAddress, e 
}

func calculateCPUUsage() (float32, error) {
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

	stat, err := linuxproc.ReadStat("/proc/stat")
	if err != nil {
		return 0, err
	}

	aggIdle := stat.CPUStats[0].Idle + stat.CPUStats[0].IOWait
	aggUserTime := stat.CPUStats[0].User - stat.CPUStats[0].Guest
	aggNiceTime := stat.CPUStats[0].Nice - stat.CPUStats[0].GuestNice
	aggTotal := aggUserTime + aggNiceTime + stat.CPUStats[0].System + stat.CPUStats[0].Steal + stat.CPUStats[0].IRQ + stat.CPUStats[0].SoftIRQ + aggIdle
	aggCPUUsage := float32(1 - (aggIdle / aggTotal))

	// idle := 0
	// total := 0
	// for idx, s := range stat.CPUStats {
	// 	if idx == 0 {
	// 		continue
	// 	}
	// 	idle += int(s.Idle) + int(s.IOWait)
	// 	userTime := int(s.User) - int(s.Guest)
	// 	niceTime := int(s.Nice) + int(s.GuestNice)
	// 	total += userTime + niceTime + int(s.System) + int(s.Steal) + int(s.IRQ) + int(s.SoftIRQ) + int(s.Idle) + int(s.IOWait)
	// }

	// testUsage := float32(1 - (idle / total))
	// fmt.Println(testUsage)

	return aggCPUUsage, nil
}

func calculateMemUsage() (float32, error) {
	stat, err := linuxproc.ReadMemInfo("/proc/meminfo")
	if err != nil {
		return 0, err
	}
	memUsage := float32( stat.MemFree / stat.MemTotal)
	return memUsage, nil
}

func getServerMetadata(uuid string) (domain.ServerMetadata, error) {
	_, err := getOutboundIP()
	if err != nil {
		return domain.ServerMetadata{}, err
	}

	// cpu, err := calculateCPUUsage()
	// if err != nil {
	// 	return domain.ServerMetadata{}, err
	// }
	
	// mem, err := calculateMemUsage()
	// if err != nil {
	// 	return domain.ServerMetadata{}, err
	// }

	api := url.URL{
		Scheme: "ws",
		Host: "localhost:8080", // ip.String()
		Path: "api/v1/ws",
	}
	
	sm := domain.ServerMetadata{
		Name: fmt.Sprintf("chat-server-%v", uuid),
		URL: api.String(),
		CPU: 0.423,
		Memory: 0.675,
	}
	return sm, nil
}

func (sdc *ServiceDiscoveryClient) updateServerMetadata(ctx context.Context, uuid string) error {
	rv, err := getServerMetadata(uuid)
	if err != nil {
		return err
	}

	b, err := json.Marshal(rv)
	if err != nil {
		return err
	}

	// For minimum lease TTL in seconds.
	resp, err := sdc.cli.Grant(context.TODO(), 15)
	if err != nil {
		return err
	}

	_, err = sdc.cli.Put(ctx, rv.Name, string(b), etcd.WithLease(resp.ID))
	if err != nil {
		return err
	}
	return nil
}

func (sdc *ServiceDiscoveryClient) ServiceDiscoveryHeartbeat(ctx context.Context) {
	uuid := uuid.NewString()
	for {
		// For clientv3, need to pass your own context with timeout.
		<- time.After(5 * time.Second)
		if err := sdc.updateServerMetadata(ctx, uuid); err != nil {
			logger.Error(
				"unable to update service discovery with server metadata",
				zap.String("trace", err.Error()),
			)
		}
	}
}