package main

import (
	"github.com/henrysher/packer/packer/plugin"
	"github.com/henrysher/packer/post-processor/docker-import"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterPostProcessor(new(dockerimport.PostProcessor))
	server.Serve()
}
