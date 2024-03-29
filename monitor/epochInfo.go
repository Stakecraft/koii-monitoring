package monitor

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Stakecraft/koii-mission-control/config"
	"github.com/Stakecraft/koii-mission-control/types"
	"github.com/Stakecraft/koii-mission-control/utils"
)

// GetEpochInfo returns information about the current epoch
func GetEpochInfo(cfg *config.Config, node string) (types.EpochInfo, error) {
	log.Println("Getiing EpochInfo...")
	ops := types.HTTPOptions{
		Method: http.MethodPost,
		Body:   types.Payload{Jsonrpc: "2.0", Method: "getEpochInfo", ID: 1},
	}

	if node == utils.Network {
		ops.Endpoint = cfg.Endpoints.NetworkRPC
	} else if node == utils.Validator {
		ops.Endpoint = cfg.Endpoints.RPCEndpoint
	} else {
		ops.Endpoint = cfg.Endpoints.RPCEndpoint
	}

	var result types.EpochInfo
	resp, err := HitHTTPTarget(ops)
	if err != nil {
		log.Printf("Error: %v", err)
		return result, err
	}

	err = json.Unmarshal(resp.Body, &result)
	if err != nil {
		log.Printf("Error: %v", err)
		return result, err
	}

	return result, nil
}
