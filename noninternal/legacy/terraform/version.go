package terraform

import (
	"github.com/kevinklinger/open_terraform/version"
)

// Deprecated: Providers should use schema.Provider.TerraformVersion instead
func VersionString() string {
	return version.String()
}
