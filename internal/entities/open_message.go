package entities

const OPEN_ACTION = "Open"

type OpenMessage struct {
	Id int
	Action string
}

func (o OpenMessage) Is_open_action() bool {
	return o.Action == OPEN_ACTION
}