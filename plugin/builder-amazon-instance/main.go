package main

import (
	"github.com/henrysher/packer/builder/amazon/instance"
	"github.com/henrysher/packer/packer/plugin"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterBuilder(new(instance.Builder))
	server.Serve()
}
