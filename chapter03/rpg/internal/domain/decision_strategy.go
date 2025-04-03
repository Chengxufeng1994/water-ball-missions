package domain

type DecisionStrategy interface {
	ChooseAction(actions []Action, requiredMagicPoint int) Action
	ChooseTargets(targets []Unit, requiredTargets int) []Unit
}
