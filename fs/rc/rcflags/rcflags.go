// Package rcflags implements command line flags to set up the remote control
package rcflags

import (
	"github.com/markus512/rclone/cmd/serve/httplib/httpflags"
	"github.com/markus512/rclone/fs/config/flags"
	"github.com/markus512/rclone/fs/rc"
	"github.com/spf13/pflag"
)

// Options set by command line flags
var (
	Opt = rc.DefaultOpt
)

// AddFlags adds the remote control flags to the flagSet
func AddFlags(flagSet *pflag.FlagSet) {
	flags.BoolVarP(flagSet, &Opt.Enabled, "rc", "", false, "Enable the remote control server.")
	httpflags.AddFlagsPrefix(flagSet, "rc-", &Opt.HTTPOptions)
}
