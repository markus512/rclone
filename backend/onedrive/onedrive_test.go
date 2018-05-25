// Test OneDrive filesystem interface
package onedrive_test

import (
	"testing"

	"github.com/markus512/rclone/backend/onedrive"
	"github.com/markus512/rclone/fstest/fstests"
)

// TestIntegration runs integration tests against the remote
func TestIntegration(t *testing.T) {
	fstests.Run(t, &fstests.Opt{
		RemoteName: "TestOneDrive:",
		NilObject:  (*onedrive.Object)(nil),
	})
}
