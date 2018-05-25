// Test AzureBlob filesystem interface
package azureblob_test

import (
	"testing"

	"github.com/markus512/rclone/backend/azureblob"
	"github.com/markus512/rclone/fstest/fstests"
)

// TestIntegration runs integration tests against the remote
func TestIntegration(t *testing.T) {
	fstests.Run(t, &fstests.Opt{
		RemoteName: "TestAzureBlob:",
		NilObject:  (*azureblob.Object)(nil),
	})
}
