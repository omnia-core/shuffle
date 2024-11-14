package flags

import (
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func Bind(desc string, persistent bool, cmd *cobra.Command, flags func(f *pflag.FlagSet)) {
	flagSet := pflag.NewFlagSet(desc, pflag.ContinueOnError)
	flags(flagSet)

	if persistent {
		cmd.PersistentFlags().AddFlagSet(flagSet)
	} else {
		cmd.Flags().AddFlagSet(flagSet)
	}
}
