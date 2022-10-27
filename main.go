package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/bosskrub9992/tennis_scoring_system/loggers"
)

var (
	ErrLessThan2Players          = errors.New("less than 2 players")
	ErrMoreThan2PlayersInData    = errors.New("more than 2 players in data")
	ErrPlayerNotFoundInMatchData = errors.New("player not found in match data")
)

func main() {
	logger := loggers.New()
	reader := bufio.NewReader(os.Stdin)

	// input players
	fmt.Println("Please input tennis players such as 'Lisa VS Jennie'. The 1st player will be the one who serve first.")
	fmt.Print("-> ")
	rawPlayers, err := reader.ReadString('\n')
	if err != nil {
		logger.Error(err)
		return
	}

	rawPlayers = strings.TrimSpace(rawPlayers)
	players := strings.Split(rawPlayers, " VS ")
	if len(players) < 2 {
		logger.Error(ErrLessThan2Players)
		return
	}

	player1 := players[0]
	player2 := players[1]

	// input match data
	fmt.Println("Please input the tennis match data.")
	fmt.Print("-> ")
	raw, err := reader.ReadString(']')
	if err != nil {
		logger.Error(err)
		return
	}

	var match [][]string

	// clean data
	raw = strings.Replace(raw, "[", "", 1)
	raw = strings.Replace(raw, "]", "", 1)
	raw = strings.ReplaceAll(raw, " ", "")
	raw = strings.ReplaceAll(raw, "\t", "")
	raw = strings.ReplaceAll(raw, ",\n", "\n")
	raw = strings.ReplaceAll(raw, ",\r\n", "\n")
	raw = strings.TrimSpace(raw)
	rawMatch := strings.Split(raw, "\n")
	for _, rawSet := range rawMatch {
		rawSet = strings.TrimSpace(rawSet)
		set := strings.Split(rawSet, ",")
		match = append(match, set)
	}

	// check unique player
	uniquePlayers := make(map[string]bool)
	for _, set := range match {
		for _, gameWinner := range set {
			uniquePlayers[gameWinner] = true
		}
	}
	if len(uniquePlayers) > 2 {
		logger.Errorf("error more than 2 player detected: %+v", uniquePlayers)
		return
	}
	if _, found := uniquePlayers[player1]; !found {
		logger.Error(ErrPlayerNotFoundInMatchData)
		return
	}
	if _, found := uniquePlayers[player2]; !found {
		logger.Error(ErrPlayerNotFoundInMatchData)
		return
	}

	// logic
	servePlayer := player1
	player1WinSetCount := 0
	player2WinSetCount := 0

	for _, set := range match {
		player1WinGameCount := 0
		player2WinGameCount := 0
		gameScores := []string{}

		for setIndex, gameWinner := range set {
			player1GameScore := ""
			player2GameScore := ""
			BP := ""
			SP := ""
			if gameWinner == player1 {
				player1WinGameCount++
			} else {
				player2WinGameCount++
			}
			switch player1WinGameCount {
			case 0:
				player1GameScore = "0"
			case 1:
				player1GameScore = "15"
			case 2:
				player1GameScore = "30"
			case 3:
				player1GameScore = "40"
			}
			switch player2WinGameCount {
			case 0:
				player2GameScore = "0"
			case 1:
				player2GameScore = "15"
			case 2:
				player2GameScore = "30"
			case 3:
				player2GameScore = "40"
			}
			if player1WinGameCount >= 3 && player2WinGameCount >= 3 {
				switch {
				case player1WinGameCount < player2WinGameCount:
					player1GameScore = "40"
					player2GameScore = "A"
					if servePlayer == player1 {
						BP = "BP"
					}
					if player1WinSetCount == 5 || player2WinSetCount == 5 {
						SP = "SP"
					}
				case player1WinGameCount > player2WinGameCount:
					player1GameScore = "A"
					player2GameScore = "40"
					if servePlayer == player2 {
						BP = "BP"
					}
					if player1WinSetCount == 5 || player2WinSetCount == 5 {
						SP = "SP"
					}
				default:
					player1GameScore = "40"
					player2GameScore = "40"
				}
			}
			if (servePlayer == player2 && player1WinGameCount == 3 && player2WinGameCount < 3) ||
				(servePlayer == player1 && player1WinGameCount < 3 && player2WinGameCount == 3) {
				BP = "BP"
			}
			if (player1WinGameCount == 3 && player2WinGameCount < 3) ||
				(player1WinGameCount < 3 && player2WinGameCount == 3) {
				if player1WinSetCount == 5 || player2WinSetCount == 5 {
					SP = "SP"
				}
			}
			if setIndex != len(set)-1 {
				gameScores = append(gameScores, fmt.Sprintf("%s:%s%s%s", player1GameScore, player2GameScore, BP, SP))
			}
		}

		if player1WinGameCount > player2WinGameCount {
			player1WinSetCount++
		} else {
			player2WinSetCount++
		}

		fmt.Printf("%s Serve %d-%d\n", servePlayer, player1WinSetCount, player2WinSetCount)
		fmt.Printf("%s\n\n", strings.Join(gameScores, ", "))

		if servePlayer == player1 {
			servePlayer = player2
		} else {
			servePlayer = player1
		}
	}
}
