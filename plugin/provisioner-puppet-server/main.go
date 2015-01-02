package main

import (
	"github.com/henrysher/packer/packer/plugin"
	"github.com/henrysher/packer/provisioner/puppet-server"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterProvisioner(new(puppetserver.Provisioner))
	server.Serve()
}
