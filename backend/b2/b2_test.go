// Test B2 filesystem interface
package b2_test

import (
	"testing"

	"github.com/markus512/rclone/backend/b2"
	"github.com/markus512/rclone/fstest/fstests"
)

// TestIntegration runs integration tests against the remote
func TestIntegration(t *testing.T) {
	fstests.Run(t, &fstests.Opt{
		RemoteName: "TestB2:",
		NilObject:  (*b2.Object)(nil),
	})
}
