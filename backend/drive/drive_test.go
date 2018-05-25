// Test Drive filesystem interface
package drive_test

import (
	"testing"

	"github.com/markus512/rclone/backend/drive"
	"github.com/markus512/rclone/fstest/fstests"
)

// TestIntegration runs integration tests against the remote
func TestIntegration(t *testing.T) {
	fstests.Run(t, &fstests.Opt{
		RemoteName: "TestDrive:",
		NilObject:  (*drive.Object)(nil),
	})
}
