package pvm

import (
	"bytes"
	"github.com/mitchellh/multistep"
	parallelscommon "github.com/henrysher/packer/builder/parallels/common"
	"github.com/henrysher/packer/packer"
	"testing"
)

func testState(t *testing.T) multistep.StateBag {
	state := new(multistep.BasicStateBag)
	state.Put("driver", new(parallelscommon.DriverMock))
	state.Put("ui", &packer.BasicUi{
		Reader: new(bytes.Buffer),
		Writer: new(bytes.Buffer),
	})
	return state
}
