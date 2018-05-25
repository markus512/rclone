// Test Webdav filesystem interface
package webdav_test

import (
	"testing"

	"github.com/markus512/rclone/backend/webdav"
	"github.com/markus512/rclone/fstest/fstests"
)

// TestIntegration runs integration tests against the remote
func TestIntegration(t *testing.T) {
	fstests.Run(t, &fstests.Opt{
		RemoteName: "TestWebdav:",
		NilObject:  (*webdav.Object)(nil),
	})
}
