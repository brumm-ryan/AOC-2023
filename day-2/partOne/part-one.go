package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	bag := make(map[string]int)
	bag["red"] = 12
	bag["green"] = 13
	bag["blue"] = 14

	result := 0

	input := getInput("../input.txt")
	for _, line := range input {
		game := parseGame(line)
		result += scoreGame(bag, game)
	}

	fmt.Println("Result: ", result)

}

func parseGame(gameStr string) game {

	gameStruct := game{}

	gameStr, ballStr, _ := strings.Cut(gameStr, ":")
	gameStruct.id = parseGameId(gameStr)

	roundsStr := strings.Split(ballStr, ";")
	for _, roundStr := range roundsStr {
		roundStr = strings.TrimSpace(roundStr)
		ballsStr := strings.Split(roundStr, ",")
		ballMap := make(map[string]int)
		gameStruct.rounds = append(gameStruct.rounds, ballMap)
		for _, ballStr := range ballsStr {
			ballStr = strings.TrimSpace(ballStr)
			countStr, color, _ := strings.Cut(ballStr, " ")
			color = strings.TrimSpace(color)
			count, _ := strconv.ParseInt(countStr, 10, 32)
			ballMap[color] += int(count)
		}
	}
	return gameStruct
}

func parseGameId(gameStr string) int {
	_, gameId, _ := strings.Cut(strings.TrimSpace(gameStr), " ")
	gameId64, _ := strconv.ParseInt(strings.TrimSpace(gameId), 10, 32)
	return int(gameId64)
}

func scoreGame(bag map[string]int, game game) int {
	for _, j := range game.rounds {
		for x, y := range j {
			if y > bag[x] {
				return 0
			}
		}
	}
	return game.id
}

func getInput(filename string) []string {

	readFile, err := os.Open(filename)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()
	return fileLines
}

type game struct {
	id     int
	rounds []map[string]int
}
