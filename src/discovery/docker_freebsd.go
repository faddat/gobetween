//go:build freebsd

package discovery

/**
 * docker.go - Docker API discovery implementation
 *
 * @author Yaroslav Pogrebnyak <yyyaroslav@gmail.com>
 */

import (
	"fmt"

	"github.com/yyyar/gobetween/config"
	"github.com/yyyar/gobetween/core"
)

/**
 * Create new Discovery with Docker fetch func
 */
func NewDockerDiscovery(_ config.DiscoveryConfig) interface{} {
	d := Discovery{
		fetch: dockerFetch,
	}
	return &d
}

/**
 * Fetch backends from Docker API
 */
func dockerFetch(cfg config.DiscoveryConfig) (*[]core.Backend, error) {
	return nil, fmt.Errorf("Docker discovery not avalilable in freebsd")
}
