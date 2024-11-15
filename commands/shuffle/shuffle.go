package shuffle

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/kevinhwang-dev/shuffle/commands/help"

	"github.com/spf13/cobra"
)

func Cmd(cmd *cobra.Command) {
	filePath, _ := cmd.Flags().GetString("file")
	namesString, _ := cmd.Flags().GetString("names")
	teams, _ := cmd.Flags().GetInt("teams")
	size, _ := cmd.Flags().GetInt("size")

	if err := valid(filePath, namesString, teams, size); err != nil {
		fmt.Println(err)
		return
	}

	var names []string
	if filePath != "" {
		file, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Println(err)
			return
		}
		names = strings.Split(string(file), "\n")
	} else if namesString != "" {
		names = strings.Split(namesString, ",")
	}

	// Shuffle the names
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	r.Shuffle(len(names), func(i, j int) { names[i], names[j] = names[j], names[i] })

	// Print the teams
	for i, team := range makeTeams(teams, size, names) {
		fmt.Printf("Team %d: %v\n", i+1, team)
	}
}

func valid(filePath, names string, teams, size int) error {
	if filePath == "" && names == "" && teams == 0 && size == 0 {
		return fmt.Errorf(help.GetHelp("rootCmd"))
	}
	if filePath == "" && names == "" {
		return fmt.Errorf("either -f or -n flag should be provided")
	}
	if filePath != "" && names != "" {
		return fmt.Errorf("only one of -f or -n flag should be provided")
	}
	if teams == 0 && size == 0 {
		return fmt.Errorf("number of teams should be greater than 0")
	}
	return nil
}

func makeTeams(numOfTeams int, teamSize int, names []string) [][]string {
	if teamSize == 0 && numOfTeams > 0 {
		teamSize = (len(names) + numOfTeams - 1) / numOfTeams // Round up to ensure all names are included
	} else if numOfTeams == 0 && teamSize > 0 {
		numOfTeams = (len(names) + teamSize - 1) / teamSize
	}

	var result [][]string
	result = make([][]string, numOfTeams)
	teamIndex := 0

loop:
	for _, name := range names {
		for len(result[teamIndex]) >= teamSize {
			if teamIndex >= numOfTeams-1 {
				break loop
			}
			teamIndex = (teamIndex + 1) % numOfTeams
		}
		result[teamIndex] = append(result[teamIndex], name)
		teamIndex = (teamIndex + 1) % numOfTeams
	}

	return result
}
