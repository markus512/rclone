// Test Swift filesystem interface
package swift_test

import (
	"testing"

	"github.com/markus512/rclone/backend/swift"
	"github.com/markus512/rclone/fstest/fstests"
)

// TestIntegration runs integration tests against the remote
func TestIntegration(t *testing.T) {
	fstests.Run(t, &fstests.Opt{
		RemoteName: "TestSwift:",
		NilObject:  (*swift.Object)(nil),
	})
}
