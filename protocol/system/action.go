package system

type Action string

const (
	ActionCreate = Action("c")
	ActionRead   = Action("r")
	ActionUpdate = Action("u")
	ActionDelete = Action("d")
)

type ActionBody[T any] struct {
	Action Action `json:"action"`
	Data   T      `json:"data"`
}

func NewActionMessage[D any](tags []string, action Action, data D) Message[ActionBody[D]] {
	return Message[ActionBody[D]]{
		Tags: tags,
		Type: MessageAction,
		Body: ActionBody[D]{
			Action: action,
			Data:   data,
		},
	}
}
