package response

import "elim5/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
