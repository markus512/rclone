package tree

import (
	"bytes"
	"testing"

	"github.com/a8m/tree"
	"github.com/markus512/rclone/fs"
	"github.com/markus512/rclone/fstest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	_ "github.com/markus512/rclone/backend/local"
)

func TestTree(t *testing.T) {
	fstest.Initialise()

	buf := new(bytes.Buffer)

	f, err := fs.NewFs("testfiles")
	require.NoError(t, err)
	err = Tree(f, buf, new(tree.Options))
	require.NoError(t, err)
	assert.Equal(t, `/
├── file1
├── file2
├── file3
└── subdir
    ├── file4
    └── file5

1 directories, 5 files
`, buf.String())
}
