package main

import (
	"github.com/henrysher/packer/packer/plugin"
	"github.com/henrysher/packer/post-processor/vagrant"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterPostProcessor(new(vagrant.PostProcessor))
	server.Serve()
}
