package destination

import (
	"context"
	"fmt"
	"time"

	"github.com/cenkalti/backoff"

	"github.com/rudderlabs/rudder-server/config"
	backendconfig "github.com/rudderlabs/rudder-server/config/backend-config"
	"github.com/rudderlabs/rudder-server/regulation-worker/internal/model"
	"github.com/rudderlabs/rudder-server/utils/logger"
)

var pkgLogger = logger.NewLogger().Child("client")

//go:generate mockgen -source=destination.go -destination=mock_destination_test.go -package=destination github.com/rudderlabs/rudder-server/regulation-worker/internal/Destination/destination
type destinationMiddleware interface {
	Get(ctx context.Context, workspace string) (map[string]backendconfig.ConfigT, error)
}

type DestMiddleware struct {
	Dest destinationMiddleware
}

func (d *DestMiddleware) GetWorkspaceId(ctx context.Context) (string, error) {
	pkgLogger.Debugf("getting destination Id")
	destConfig, err := d.getDestDetails(ctx)
	if err != nil {
		pkgLogger.Errorf("error while getting destination details from backend config: %v", err)
		return "", err
	}
	if len(destConfig) == 1 { // only single workspace configs are supported by regulation worker
		for workspaceID := range destConfig {
			pkgLogger.Debugf("workspaceId=", workspaceID)
			return workspaceID, nil
		}
	}

	pkgLogger.Error("workspaceId not found in config")
	return "", fmt.Errorf("workspaceId not found in config")
}

// GetDestDetails makes api call to get json and then parse it to get destination related details
// like: dest_type, auth details,
// return destination Type enum{file, api}
func (d *DestMiddleware) GetDestDetails(ctx context.Context, destID string) (model.Destination, error) {
	pkgLogger.Debugf("getting destination details for destinationId: %v", destID)
	destConf, err := d.getDestDetails(ctx)
	if err != nil {
		return model.Destination{}, err
	}

	for _, wConf := range destConf {
		for _, source := range wConf.Sources {
			for _, dest := range source.Destinations {
				if dest.ID == destID {
					var destDetail model.Destination
					destDetail.Config = dest.Config
					destDetail.DestinationID = dest.ID
					destDetail.Name = dest.DestinationDefinition.Name
					// Destination Definition Config would most likely be needed
					destDetail.DestDefConfig = dest.DestinationDefinition.Config
					pkgLogger.Debugf("obtained destination detail: %v", destDetail)
					return destDetail, nil
				}
			}
		}
	}
	return model.Destination{}, model.ErrInvalidDestination
}

func (d *DestMiddleware) getDestDetails(ctx context.Context) (map[string]backendconfig.ConfigT, error) {
	pkgLogger.Debugf("getting destination details with exponential backoff")

	maxWait := time.Minute * 10
	var err error
	bo := backoff.NewExponentialBackOff()
	boCtx := backoff.WithContext(bo, ctx)
	bo.MaxInterval = time.Minute
	bo.MaxElapsedTime = maxWait
	var destConf map[string]backendconfig.ConfigT
	if err = backoff.Retry(func() error {
		pkgLogger.Debugf("Fetching backend-config...")
		// TODO : Revisit the Implementation for Regulation Worker in case of MultiTenant Deployment
		destConf, err = d.Dest.Get(ctx, config.GetWorkspaceToken())
		if err != nil {
			return fmt.Errorf("error while getting destination details: %w", err)
		}
		return nil
	}, boCtx); err != nil {
		if bo.NextBackOff() == backoff.Stop {
			pkgLogger.Debugf("reached retry limit...")
			return map[string]backendconfig.ConfigT{}, err
		}
	}
	return destConf, nil
}
