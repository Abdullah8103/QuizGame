package rules

import (
	"fmt"
)

const QuestionsPerMatch = 20

type Player struct {
	name   string
	health int
}

type Question struct {
	text        string
	choices     []string
	answerIndex int
}

type Match struct {
	Players              []*Player
	Questions            []Question
	currentQuestionIndex int
}

func newPlayer(name string) *Player {
	var player Player
	player.name = name
	player.health = 100
	return &player
}

func NewMatch(names []string, questions []Question) (*Match, error) {
	var match Match
	match.Questions = questions
	match.currentQuestionIndex = 0

	if len(match.Questions) != QuestionsPerMatch {
		return nil, fmt.Errorf("you can only have %d questions, you added %d", QuestionsPerMatch, len(questions))
	}
	if len(names) <= 1 {
		return nil, fmt.Errorf("you need one or more people to play with")
	}

	for _, name := range names {
		match.Players = append(match.Players, newPlayer(name))
	}

	return &match, nil
}

func score
