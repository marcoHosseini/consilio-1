package terraform

import (
	"github.com/kevinklinger/open_terraform/noninternal/addrs"
	"github.com/kevinklinger/open_terraform/noninternal/configs"
)

// GraphNodeAttachProvider is an interface that must be implemented by nodes
// that want provider configurations attached.
type GraphNodeAttachProvider interface {
	// ProviderName with no module prefix. Example: "aws".
	ProviderAddr() addrs.AbsProviderConfig

	// Sets the configuration
	AttachProvider(*configs.Provider)
}
