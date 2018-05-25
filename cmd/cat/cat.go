package cat

import (
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/markus512/rclone/cmd"
	"github.com/markus512/rclone/fs/operations"
	"github.com/spf13/cobra"
)

// Globals
var (
	head    = int64(0)
	tail    = int64(0)
	offset  = int64(0)
	count   = int64(-1)
	discard = false
)

func init() {
	cmd.Root.AddCommand(commandDefintion)
	commandDefintion.Flags().Int64VarP(&head, "head", "", head, "Only print the first N characters.")
	commandDefintion.Flags().Int64VarP(&tail, "tail", "", tail, "Only print the last N characters.")
	commandDefintion.Flags().Int64VarP(&offset, "offset", "", offset, "Start printing at offset N (or from end if -ve).")
	commandDefintion.Flags().Int64VarP(&count, "count", "", count, "Only print N characters.")
	commandDefintion.Flags().BoolVarP(&discard, "discard", "", discard, "Discard the output instead of printing.")
}

var commandDefintion = &cobra.Command{
	Use:   "cat remote:path",
	Short: `Concatenates any files and sends them to stdout.`,
	Long: `
rclone cat sends any files to standard output.

You can use it like this to output a single file

    rclone cat remote:path/to/file

Or like this to output any file in dir or subdirectories.

    rclone cat remote:path/to/dir

Or like this to output any .txt files in dir or subdirectories.

    rclone --include "*.txt" cat remote:path/to/dir

Use the --head flag to print characters only at the start, --tail for
the end and --offset and --count to print a section in the middle.
Note that if offset is negative it will count from the end, so
--offset -1 --count 1 is equivalent to --tail 1.
`,
	Run: func(command *cobra.Command, args []string) {
		usedOffset := offset != 0 || count >= 0
		usedHead := head > 0
		usedTail := tail > 0
		if usedHead && usedTail || usedHead && usedOffset || usedTail && usedOffset {
			log.Fatalf("Can only use one of  --head, --tail or --offset with --count")
		}
		if head > 0 {
			offset = 0
			count = head
		}
		if tail > 0 {
			offset = -tail
			count = -1
		}
		cmd.CheckArgs(1, 1, command, args)
		fsrc := cmd.NewFsSrc(args)
		var w io.Writer = os.Stdout
		if discard {
			w = ioutil.Discard
		}
		cmd.Run(false, false, command, func() error {
			return operations.Cat(fsrc, w, offset, count)
		})
	},
}
