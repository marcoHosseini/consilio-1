package main

import (
	"github.com/kevinklinger/open_terraform/noninternal/grpcwrap"
	"github.com/kevinklinger/open_terraform/noninternal/plugin"
	simple "github.com/kevinklinger/open_terraform/noninternal/provider-simple"
	"github.com/kevinklinger/open_terraform/noninternal/tfplugin5"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		GRPCProviderFunc: func() tfplugin5.ProviderServer {
			return grpcwrap.Provider(simple.Provider())
		},
	})
}
