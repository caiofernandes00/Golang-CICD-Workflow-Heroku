package circuitbreaker

type (
	ChangeState struct {
		From State
		To   State
	}
)
