package main

import (
	"github.com/henrysher/packer/builder/googlecompute"
	"github.com/henrysher/packer/packer/plugin"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterBuilder(new(googlecompute.Builder))
	server.Serve()
}
