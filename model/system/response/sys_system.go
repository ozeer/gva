package response

import "github.com/ozeer/gva/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
