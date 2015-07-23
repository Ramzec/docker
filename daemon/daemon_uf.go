// +build !exclude_graphdriver_uf,linux

package daemon

import (
	_ "github.com/docker/docker/daemon/graphdriver/union-federated"
)
