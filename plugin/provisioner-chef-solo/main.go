package main

import (
	"github.com/henrysher/packer/packer/plugin"
	"github.com/henrysher/packer/provisioner/chef-solo"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterProvisioner(new(chefsolo.Provisioner))
	server.Serve()
}
