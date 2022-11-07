package main

type TurnState int

const (
	BeforePlayingAction = iota
	PlayerTurn
	MonsterTurn
)

func GetNextState(state TurnState) TurnState {
	switch state {
	case BeforePlayingAction:
		return PlayerTurn
	case PlayerTurn:
		return MonsterTurn
	case MonsterTurn:
		return PlayerTurn
	default:
		return PlayerTurn
	}
}
