package googlecompute

import (
	"github.com/henrysher/packer/packer"
	"testing"
)

func TestArtifact_impl(t *testing.T) {
	var _ packer.Artifact = new(Artifact)
}
