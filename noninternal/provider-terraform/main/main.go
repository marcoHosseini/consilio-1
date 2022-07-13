package main

import (
	"github.com/kevinklinger/open_terraform/noninternal/builtin/providers/terraform"
	"github.com/kevinklinger/open_terraform/noninternal/grpcwrap"
	"github.com/kevinklinger/open_terraform/noninternal/plugin"
	"github.com/kevinklinger/open_terraform/noninternal/tfplugin5"
)

func main() {
	// Provide a binary version of the internal terraform provider for testing
	plugin.Serve(&plugin.ServeOpts{
		GRPCProviderFunc: func() tfplugin5.ProviderServer {
			return grpcwrap.Provider(terraform.NewProvider())
		},
	})
}
