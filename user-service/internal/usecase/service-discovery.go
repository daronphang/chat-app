package usecase

import (
	"context"
	"errors"
)

var (
	errNoServerMatch = errors.New("unable to find server that meet domain criteria")
)


func (uc *UseCaseService) GetBestServer(ctx context.Context) (string, error) {
	rv, err := uc.ServiceDiscovery.GetServersMetdata(ctx)
	if err != nil {
		return "", err
	}

	// To find the server with CPU < 50% and memory < 80%.
	// First come first serve basis.
	var server string
	for _, s := range rv {
		if s.CPU < 0.5 && s.Memory < 0.8 {
			server = s.URL
			break
		}
	}

	if server == "" {
		return "", errNoServerMatch
	}

	return "", nil
}