package blackjack

import (
	"fmt"
	"strings"
)

type Action interface {
	fmt.Stringer
	Help() string
	Do(data *BlackjackData) GameState
}

func NewActions(actions ...Action) []Action {
	var res []Action
	for _, a := range actions {
		res = append(res, a)
	}

	res = append(res, HelpAction{res})
	return res
}

func Prompt(data *BlackjackData, actions []Action) Action {
	fmt.Println("What will you do?")
	fmt.Printf("> ")

	var choice string
	fmt.Scanf("%s", &choice)
	choice = strings.Trim(choice, " ")

	for _, a := range actions {
		if a.String() == choice {
			return a
		}
	}

	return nil
}

type HelpAction struct {
	actions []Action
}

func (ha HelpAction) String() string {
	return "help"
}

func (ha HelpAction) Help() string {
	return "show info about available options"
}

func (ha HelpAction) Do(data *BlackjackData) GameState {
	for _, a := range ha.actions {
		fmt.Printf("\t%s - %s\n", a.String(), a.Help())
	}

	return nil
}

type HitAction struct{}

func (ha HitAction) String() string {
	return "hit"
}

func (ha HitAction) Help() string {
	return "draw a new card"
}

func (ha HitAction) Do(data *BlackjackData) GameState {
	return HitState
}

type StandAction struct{}

func (ha StandAction) String() string {
	return "stand"
}

func (ha StandAction) Help() string {
	return "end turn and proceed with next player"
}

func (ha StandAction) Do(data *BlackjackData) GameState {
	return StandState
}
