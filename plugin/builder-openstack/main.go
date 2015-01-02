package main

import (
	"github.com/henrysher/packer/builder/openstack"
	"github.com/henrysher/packer/packer/plugin"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterBuilder(new(openstack.Builder))
	server.Serve()
}
