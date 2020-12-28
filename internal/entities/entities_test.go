package entities

import (
	"testing"
)

func TestMessageId(t *testing.T) {

	var gate_id int = 99
	var action string = "my action"

	var msg = OpenMessage{Id: gate_id, Action: action}

	if msg.Id != gate_id{
		t.Errorf("Message id  mismatch. Wanted %d, got %d", msg.Id, gate_id)
	}

	if msg.Action != action{
		t.Errorf("Message action  mismatch. Wanted %s, got %s", msg.Action, action)
	}
}

func TestOpenAction(t *testing.T) {

	t.Run("Action is open", func(t *testing.T) {

		var action string = OPEN_ACTION
		var msg = OpenMessage{Id: 9, Action: action}

		if !msg.Is_open_action() {
			t.Errorf("Should return true with action open")
		}

	})

	t.Run("Action is not open", func(t *testing.T) {

		var action string = "something"
		var msg = OpenMessage{Id: 9, Action: action}

		if msg.Is_open_action() {
			t.Errorf("Should return false with action open")
		}

	})

}