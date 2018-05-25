// Test Sftp filesystem interface

// +build !plan9,go1.9

package sftp_test

import (
	"testing"

	"github.com/markus512/rclone/backend/sftp"
	"github.com/markus512/rclone/fstest/fstests"
)

// TestIntegration runs integration tests against the remote
func TestIntegration(t *testing.T) {
	fstests.Run(t, &fstests.Opt{
		RemoteName: "TestSftp:",
		NilObject:  (*sftp.Object)(nil),
	})
}
