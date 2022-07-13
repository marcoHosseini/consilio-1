package main

import (
	"github.com/kevinklinger/open_terraform/noninternal/grpcwrap"
	plugin "github.com/kevinklinger/open_terraform/noninternal/plugin6"
	simple "github.com/kevinklinger/open_terraform/noninternal/provider-simple-v6"
	"github.com/kevinklinger/open_terraform/noninternal/tfplugin6"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		GRPCProviderFunc: func() tfplugin6.ProviderServer {
			return grpcwrap.Provider6(simple.Provider())
		},
	})
}
