package rules
 
import (
	"fmt"
)
 
const QuestionsPerMatch = 20
const PenaltyScore = 15
 
type Player struct {
	name    string
	health  int
	answers []int
	time    []int
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
 
func evalScore(answerTimeMs int, correct bool) (int, error) {
	if answerTimeMs <= 0 {
		return 0, fmt.Errorf("time returned is below or equal to 0 ms")
	}
	if answerTimeMs > 10000 || !correct {
		return PenaltyScore, nil
	}
	return (answerTimeMs*10 + 9999) / 10000, nil
}
 
// roundInput turns one player's recorded answer into the two values evalScore needs.
// An answer of -1 means the player never answered, which is charged as a late wrong answer.
func (m *Match) roundInput(player *Player, qIndex int) (int, bool) {
	if player.answers[qIndex] == -1 {
		return 10001, false
	}
	return player.time[qIndex], player.answers[qIndex] == m.Questions[qIndex].answerIndex
}
 
// evalRound applies the current question to every player. It checks all inputs
// first, so if any input is bad it returns an error and changes nothing.
func (m *Match) evalRound() error {
	qIndex := m.currentQuestionIndex
	if qIndex >= len(m.Questions) {
		return fmt.Errorf("match is over, no questions left")
	}
 
	losses := make([]int, len(m.Players))
	for i, player := range m.Players {
		if qIndex >= len(player.answers) || qIndex >= len(player.time) {
			return fmt.Errorf("player %s has no input for question %d", player.name, qIndex)
		}
		timeMs, correct := m.roundInput(player, qIndex)
		loss, err := evalScore(timeMs, correct)
		if err != nil {
			return fmt.Errorf("player %s: %w", player.name, err)
		}
		losses[i] = loss
	}
 
	for i, player := range m.Players {
		player.health -= losses[i]
	}
	m.currentQuestionIndex++
	return nil
}
 