// Test FTP filesystem interface
package ftp_test

import (
	"testing"

	"github.com/markus512/rclone/backend/ftp"
	"github.com/markus512/rclone/fstest/fstests"
)

// TestIntegration runs integration tests against the remote
func TestIntegration(t *testing.T) {
	fstests.Run(t, &fstests.Opt{
		RemoteName: "TestFTP:",
		NilObject:  (*ftp.Object)(nil),
	})
}
