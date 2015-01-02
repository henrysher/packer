package main

import (
	"github.com/henrysher/packer/packer/plugin"
	"github.com/henrysher/packer/post-processor/atlas"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterPostProcessor(new(atlas.PostProcessor))
	server.Serve()
}
