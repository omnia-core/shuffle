package help

import "strings"

const bold = "\033[1m"
const reset = "\033[0m"

var (
	cmdHelp = map[string]string{
		"rootCmd": shuffleHelp,
	}

	shuffleHelp = `[[.Bold]]Command:[[.Normal]] shuffle
[[.Bold]]About:[[.Normal]] Shuffle the list of team members
[[.Bold]]Usage:[[.Normal]] random-team shuffle [flags]
[[.Bold]]Flags:[[.Normal]]
[[.Bold]]-n, --names[[.Normal]] string - List of names separated by comma
[[.Bold]]-f, --file[[.Normal]] string - File containing list of names, one name per line
[[.Bold]]-t, --teams[[.Normal]] int - Number of teams
[[.Bold]]-s, --size[[.Normal]] int - Size of each team

Only one of the flags -n or -f should be provided.

[[.Bold]]Example:[[.Normal]] random-team shuffle -n "Alice,Bob,Charlie,David,Eve" -t 2 -s 2
[[.Bold]]Output:[[.Normal]] 
	Team 1: [Eve Alice]
	Team 2: [Bob Charlie]`
)

func GetHelp(cmd string) string {
	return strings.ReplaceAll(strings.ReplaceAll(cmdHelp[cmd], "[[.Bold]]", bold), "[[.Normal]]", reset)
}
