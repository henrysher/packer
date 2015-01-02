package main

import (
	"github.com/henrysher/packer/builder/virtualbox/ovf"
	"github.com/henrysher/packer/packer/plugin"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterBuilder(new(ovf.Builder))
	server.Serve()
}
