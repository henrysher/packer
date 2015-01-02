package main

import (
	"github.com/henrysher/packer/packer/plugin"
	"github.com/henrysher/packer/post-processor/docker-tag"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterPostProcessor(new(dockertag.PostProcessor))
	server.Serve()
}
