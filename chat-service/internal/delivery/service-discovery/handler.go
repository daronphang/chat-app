package svcdis

import (
	"chat-service/internal"
	"chat-service/internal/delivery/service-discovery/cgroup2"
	"chat-service/internal/domain"
	"context"
	"encoding/json"
	"fmt"
	"net"
	"time"

	"github.com/google/uuid"
	etcd "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
)


var (
	hostIPAddress string
	logger, _ = internal.WireLogger()
	prevStat *cgroup2.ContainerStat
)

func isDockerBridgeIP(ip net.IP) bool {
	// For Docker, IP address starting with 172.17/18 is commonly associated with the
	// default Docker bridge network, which is used for communication between containers
	// on the same Docker host. 
    if ip.To4() == nil {
        return false // Not an IPv4 address
    }
    // Docker's default bridge network subnet
    subnet := net.IPNet{
        IP:   net.ParseIP("172.18.0.0"),
        Mask: net.CIDRMask(16, 32),
    }
    return subnet.Contains(ip)
}

func getOutboundIP() (string, error) {
	if hostIPAddress != "" {
		return hostIPAddress, nil
	}

	interfaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	for _, iface := range interfaces {
		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}

		for _, addr := range addrs {
			ipNet, ok := addr.(*net.IPNet)
			if ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil && !isDockerBridgeIP(ipNet.IP) {
				hostIPAddress = ipNet.IP.String()
				break
			}
		}
	}
    return hostIPAddress, nil
}

func getServerMetadata(uuid string) (domain.ServerMetadata, error) {
	ipAddr, err := getOutboundIP()
	if err != nil {
		return domain.ServerMetadata{}, err
	}

	cpuUsage := float64(0)
	curStat, err := cgroup2.GetContainerStat()
	if err != nil {
		return domain.ServerMetadata{}, err
	}
	if prevStat != nil {
		cpuUsage, err = cgroup2.CalculateCPUUsage(prevStat, curStat)
		if err != nil {
			return domain.ServerMetadata{}, err
		}
	}
	memUsage := cgroup2.CalculateMemUsage(curStat)
	prevStat = curStat

	// api := url.URL{
	// 	Scheme: "http",
	// 	Host: ip.String(),
	// }

	sm := domain.ServerMetadata{
		Name: fmt.Sprintf("chat-server-%v", uuid),
		URL: ipAddr,
		CPU: cpuUsage,
		Memory: memUsage,
	}
	return sm, nil
}

func (sdc *ServiceDiscoveryClient) saveServerMetadataInServiceDiscovery(ctx context.Context, uuid string) error {
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

func (sdc *ServiceDiscoveryClient) SendHeartbeatToServiceDiscovery(ctx context.Context) {
	uuid := uuid.NewString()
	for {
		// For clientv3, need to pass your own context with timeout.
		<- time.After(15 * time.Second)
		if err := sdc.saveServerMetadataInServiceDiscovery(ctx, uuid); err != nil {
			logger.Error(
				"unable to update service discovery with server metadata",
				zap.String("trace", err.Error()),
			)
		}
	}
}