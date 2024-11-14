/*
Copyright © 2024 Seunghyun Hwang <kevinhwang1227@gmail.com>
*/
package commands

import (
	"os"

	"github.com/kevinhwang-dev/random-team/commands/flags"
	"github.com/kevinhwang-dev/random-team/commands/help"
	"github.com/kevinhwang-dev/random-team/commands/shuffle"
	"github.com/kevinhwang-dev/random-team/consts"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "random-team",
	Short: "A shuffling tool for random team",
	Long: `A shuffling tool for random team.
The tool will shuffle the list of names and divide them into specified number of teams of the specified size.

Example:
random-team shuffle -n "Alice,Bob,Charlie,David,Eve" -t 2 -s 2`,
}

var cfgFile string

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	shuffleCmd := &cobra.Command{
		Use:   consts.ShuffleStr,
		Short: "Shuffle the list of names and divide them into specified number of teams of the specified size",
		Long:  help.GetHelp(consts.ShuffleStr),
		Args:  cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			shuffle.Cmd(cmd)
		},
	}
	flags.Bind(consts.ShuffleStr, false, shuffleCmd, func(f *pflag.FlagSet) {
		f.StringP("names", "n", "", "List of names separated by comma")
		f.StringP("file", "f", "", "File containing list of names, one name per line")
		f.IntP("teams", "t", 0, "Number of teams")
		f.IntP("size", "s", 0, "Size of each team")
	})
	rootCmd.AddCommand(shuffleCmd)

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "./config/config.yaml", "config file (default is ./config/config.yaml)")
}
