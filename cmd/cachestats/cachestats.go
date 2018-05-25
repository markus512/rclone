// +build !plan9

package cachestats

import (
	"encoding/json"
	"fmt"

	"github.com/markus512/rclone/backend/cache"
	"github.com/markus512/rclone/cmd"
	"github.com/markus512/rclone/fs"
	"github.com/markus512/rclone/fs/config"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func init() {
	cmd.Root.AddCommand(commandDefinition)
}

var commandDefinition = &cobra.Command{
	Use:   "cachestats source:",
	Short: `Print cache stats for a remote`,
	Long: `
Print cache stats for a remote in JSON format
`,
	Run: func(command *cobra.Command, args []string) {
		cmd.CheckArgs(1, 1, command, args)

		_, configName, _, err := fs.ParseRemote(args[0])
		if err != nil {
			fs.Errorf("cachestats", "%s", err.Error())
			return
		}

		if !config.FileGetBool(configName, "read_only", false) {
			config.FileSet(configName, "read_only", "true")
			defer config.FileDeleteKey(configName, "read_only")
		}

		fsrc := cmd.NewFsSrc(args)
		cmd.Run(false, false, command, func() error {
			var fsCache *cache.Fs
			fsCache, ok := fsrc.(*cache.Fs)
			if !ok {
				unwrap := fsrc.Features().UnWrap
				if unwrap != nil {
					fsCache, ok = unwrap().(*cache.Fs)
				}
				if !ok {
					return errors.Errorf("%s: is not a cache remote", fsrc.Name())
				}
			}
			m, err := fsCache.Stats()
			if err != nil {
				return err
			}

			raw, err := json.MarshalIndent(m, "", "  ")
			if err != nil {
				return err
			}

			fmt.Printf("%s\n", string(raw))
			return nil
		})
	},
}
