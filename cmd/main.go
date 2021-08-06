package main

import (
	"github.com/ava-labs/avalanchego/vms/rpcchainvm"
	"github.com/hashicorp/go-plugin"
	"mvm"
)

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: rpcchainvm.Handshake,
		Plugins: map[string]plugin.Plugin{
			"vm": rpcchainvm.New(&mvm.VM{}),
		},

		GRPCServer: plugin.DefaultGRPCServer,
	})
}
