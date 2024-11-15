/*
Copyright © 2024 Seunghyun Hwang <kevinhwang1227@gmail.com>
*/
package commands

import (
	"os"

	"github.com/kevinhwang-dev/shuffle/commands/help"
	"github.com/kevinhwang-dev/shuffle/commands/shuffle"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "shuffle",
	Short: "A shuffling tool for creating random teams",
	Long:  help.GetHelp("rootCmd"),
	Run: func(cmd *cobra.Command, args []string) {
		shuffle.Cmd(cmd)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.Flags().StringP("names", "n", "", "List of names separated by comma")
	rootCmd.Flags().StringP("file", "f", "", "File containing list of names, one name per line")
	rootCmd.Flags().IntP("teams", "t", 0, "Number of teams")
	rootCmd.Flags().IntP("size", "s", 0, "Size of each team")

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
