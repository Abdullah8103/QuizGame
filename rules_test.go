package rules

import (
	"fmt"
	"testing"
)

// makeQuestions builds n simple questions so tests don't need 20 literals.
func makeQuestions(n int) []Question {
	questions := make([]Question, 0, n)
	for i := 0; i < n; i++ {
		questions = append(questions, Question{
			text:        fmt.Sprintf("question %d", i),
			choices:     []string{"a", "b", "c", "d"},
			answerIndex: 0,
		})
	}
	return questions
}

func TestNewMatchRejectsTooFewPlayers(t *testing.T) {
	tests := []struct {
		name  string
		names []string
	}{
		{"no players", nil},
		{"one player", []string{"solo"}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			match, err := NewMatch(tc.names, makeQuestions(QuestionsPerMatch))
			if err == nil {
				t.Error("expected an error, got nil")
			}
			if match != nil {
				t.Error("expected a nil match when there is an error")
			}
		})
	}
}

func TestNewMatchRejectsWrongQuestionCount(t *testing.T) {
	counts := []int{0, QuestionsPerMatch - 1, QuestionsPerMatch + 1}

	for _, n := range counts {
		t.Run(fmt.Sprintf("%d questions", n), func(t *testing.T) {
			match, err := NewMatch([]string{"a", "b", "c"}, makeQuestions(n))
			if err == nil {
				t.Error("expected an error, got nil")
			}
			if match != nil {
				t.Error("expected a nil match when there is an error")
			}
		})
	}
}

func TestNewMatchValid(t *testing.T) {
	names := []string{"alice", "bob", "carol"}

	match, err := NewMatch(names, makeQuestions(QuestionsPerMatch))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if match == nil {
		t.Fatal("expected a match, got nil")
	}

	if len(match.Players) != len(names) {
		t.Fatalf("expected %d players, got %d", len(names), len(match.Players))
	}
	for i, p := range match.Players {
		if p.name != names[i] {
			t.Errorf("player %d: expected name %q, got %q", i, names[i], p.name)
		}
		if p.health != 100 {
			t.Errorf("player %d: expected health 100, got %d", i, p.health)
		}
	}

	if len(match.Questions) != QuestionsPerMatch {
		t.Errorf("expected %d questions, got %d", QuestionsPerMatch, len(match.Questions))
	}
	if match.currentQuestionIndex != 0 {
		t.Errorf("expected current question index 0, got %d", match.currentQuestionIndex)
	}
}
