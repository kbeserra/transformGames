package representation

type Outcome struct {
	Previous *Outcome
	M        Morphism
	State    GameState
	Awards   []Award
}
