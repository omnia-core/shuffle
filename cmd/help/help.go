package help

import "random-team/consts"

var (
	cmdHelp = map[string]string{
		consts.ShuffleStr: shuffleHelp,
	}

	shuffleHelp = `[[.Bold]]Command:[[.Normal]] shuffle
[[.Bold]]About:[[.Normal]] Shuffle the list of team members
[[.Bold]]Usage:[[.Normal]] random-team shuffle [flags]
[[.Bold]]Flags:[[.Normal]]
[[.Bold]]-n, --names[[.Normal]] string[[.Bold]] (required)[[.Normal]] List of names separated by comma
[[.Bold]]-f, --file[[.Normal]] string[[.Bold]] (optional)[[.Normal]] File containing list of names, one name per line
[[.Bold]]-t, --teams[[.Normal]] int[[.Bold]] (required)[[.Normal]] Number of teams
[[.Bold]]-s, --size[[.Normal]] int[[.Bold]] (required)[[.Normal]] Size of each team

Only one of the flags -n or -f should be provided.

[[.Bold]]Example:[[.Normal]] random-team shuffle -n "Alice,Bob,Charlie,David,Eve" -t 2 -s 2`
)

func GetHelp(cmd string) string {
	return cmdHelp[cmd]
}
