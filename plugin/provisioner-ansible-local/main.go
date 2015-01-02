package main

import (
	"github.com/henrysher/packer/packer/plugin"
	"github.com/henrysher/packer/provisioner/ansible-local"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterProvisioner(new(ansiblelocal.Provisioner))
	server.Serve()
}
