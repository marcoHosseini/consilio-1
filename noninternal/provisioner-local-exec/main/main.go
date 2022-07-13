package main

import (
	localexec "github.com/kevinklinger/open_terraform/noninternal/builtin/provisioners/local-exec"
	"github.com/kevinklinger/open_terraform/noninternal/grpcwrap"
	"github.com/kevinklinger/open_terraform/noninternal/plugin"
	"github.com/kevinklinger/open_terraform/noninternal/tfplugin5"
)

func main() {
	// Provide a binary version of the internal terraform provider for testing
	plugin.Serve(&plugin.ServeOpts{
		GRPCProvisionerFunc: func() tfplugin5.ProvisionerServer {
			return grpcwrap.Provisioner(localexec.New())
		},
	})
}
