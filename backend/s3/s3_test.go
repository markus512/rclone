// Test S3 filesystem interface
package s3_test

import (
	"testing"

	"github.com/markus512/rclone/backend/s3"
	"github.com/markus512/rclone/fstest/fstests"
)

// TestIntegration runs integration tests against the remote
func TestIntegration(t *testing.T) {
	fstests.Run(t, &fstests.Opt{
		RemoteName: "TestS3:",
		NilObject:  (*s3.Object)(nil),
	})
}
