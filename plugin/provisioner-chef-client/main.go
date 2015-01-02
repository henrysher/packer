package main

import (
	"github.com/henrysher/packer/packer/plugin"
	"github.com/henrysher/packer/provisioner/chef-client"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterProvisioner(new(chefclient.Provisioner))
	server.Serve()
}
