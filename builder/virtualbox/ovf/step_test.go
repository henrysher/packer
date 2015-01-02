package ovf

import (
	"bytes"
	"github.com/mitchellh/multistep"
	vboxcommon "github.com/henrysher/packer/builder/virtualbox/common"
	"github.com/henrysher/packer/packer"
	"testing"
)

func testState(t *testing.T) multistep.StateBag {
	state := new(multistep.BasicStateBag)
	state.Put("driver", new(vboxcommon.DriverMock))
	state.Put("ui", &packer.BasicUi{
		Reader: new(bytes.Buffer),
		Writer: new(bytes.Buffer),
	})
	return state
}
