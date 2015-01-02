package main

import (
	"github.com/henrysher/packer/builder/null"
	"github.com/henrysher/packer/packer/plugin"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterBuilder(new(null.Builder))
	server.Serve()
}
