package terraform

import (
	backendInit "github.com/kevinklinger/open_terraform/noninternal/backend/init"
)

func init() {
	// Initialize the backends
	backendInit.Init(nil)
}
