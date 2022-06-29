package tournament

import (
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

/*

Algorithm:

- Iterate over the lines of games
- Store the teams in a hashmap with the
number of games, wins, losses, draws and points
- Pass through the elements in the hashmap and store them in an array
- Sort the array by points descending and alphabetically

Time Complexity: O(nlgn)
Space Complexity: O(n)

*/

type Game struct {
	team    string
	matches int
	wins    int
	draws   int
	losses  int
	points  int
}

type TeamGameMap map[string]*Game

func Tally(reader io.Reader, writer io.Writer) error {
	tm := make(TeamGameMap)

	for {
		line, err := readLine(reader)
		line = strings.TrimSpace(line)

		if err == io.EOF && len(line) == 0 {
			break
		}

		// Ignore comments and new lines
		if line == "\n" || len(line) == 0 || line[0] == '#' {
			continue
		}

		team1, team2, gameResult, splitError := splitGameResults(line)

		if splitError != nil {
			return splitError
		}

		addResults(team1, team2, gameResult, tm)

		if err != nil {
			return err
		}
	}

	sortedResults := getSortedResults(tm)
	writeResults(sortedResults, writer)

	return nil
}

func readLine(reader io.Reader) (string, error) {
	sb := strings.Builder{}
	buf := make([]byte, 1)

	for {
		n, err := reader.Read(buf)

		if err == io.EOF {
			return sb.String(), io.EOF
		}

		if err != nil {
			return "", err
		}

		if n > 0 && buf[0] == '\n' {
			return sb.String(), nil
		} else if n > 0 {
			sb.WriteByte(buf[0])
		}
	}
}

func splitGameResults(line string) (string, string, string, error) {
	gamePlayed := strings.Split(line, ";")

	if len(gamePlayed) != 3 {
		return "", "", "", errors.New("Invalid game format")
	}

	team1 := gamePlayed[0]
	team2 := gamePlayed[1]
	gameResult := gamePlayed[2]

	if gameResult != "win" && gameResult != "loss" && gameResult != "draw" {
		return "", "", "", errors.New("Invalid game")
	}

	return team1, team2, gameResult, nil
}

func addResults(team1 string, team2 string, gameResult string, tm TeamGameMap) {
	if _, ok := tm[team1]; !ok {
		tm[team1] = &Game{team: team1, matches: 0, wins: 0, draws: 0, losses: 0, points: 0}
	}

	if _, ok := tm[team2]; !ok {
		tm[team2] = &Game{team: team2, matches: 0, wins: 0, draws: 0, losses: 0, points: 0}
	}

	switch gameResult {
	case "win":
		tm[team1].wins++
		tm[team1].points += 3
		tm[team2].losses++
	case "loss":
		tm[team1].losses++
		tm[team2].wins++
		tm[team2].points += 3
	case "draw":
		tm[team1].draws++
		tm[team1].points++
		tm[team2].draws++
		tm[team2].points++
	}
}

func getSortedResults(tm TeamGameMap) []Game {
	teamsResults := make([]Game, 0, len(tm))

	for _, val := range tm {
		teamsResults = append(teamsResults, *val)
	}

	sort.Slice(teamsResults, func(i, j int) bool {
		// Sort by name in ascending order
		if teamsResults[i].points == teamsResults[j].points {
			return teamsResults[i].team < teamsResults[j].team
		}

		// Sort by points in descending order
		return teamsResults[i].points > teamsResults[j].points
	})

	return teamsResults
}

func writeResults(games []Game, writer io.Writer) {
	header := fmt.Sprintf("Team                           | MP |  W |  D |  L |  P\n")
	writer.Write([]byte(header))

	for _, game := range games {
		row := fmt.Sprintf("%-31s|%3d |%3d |%3d |%3d |%3d\n",
			game.team,
			game.wins+game.draws+game.losses,
			game.wins,
			game.draws,
			game.losses,
			game.points)

		writer.Write([]byte(row))
	}
}
