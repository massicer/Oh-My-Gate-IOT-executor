package iot_adapter

import (
	"fmt"
	"bytes"
	"testing"
	ent "github.com/massicer/Oh-My-Gate-IOT-executor/internal/entities"
)

func get_valid_open_msg(t *testing.T) ent.OpenMessage {
	t.Helper()
	var gate_id int = 99
	var action string = ent.OPEN_ACTION

	var msg = ent.OpenMessage{Id: gate_id, Action: action}
    return msg
}

func TestSdoutAdapter(t *testing.T) {
	var b bytes.Buffer
	s := Sdout_iot_adapter{W: &b}

	msg := get_valid_open_msg(t)

    if err := s.Open(msg); err != nil {
        t.Fatalf("s.get_valid_open_msg() gave error: %s", err)
    }
    got := b.String()
    want := fmt.Sprintf("Going to open iot device with id: %d", msg.Id)
    if got != want {
        t.Errorf("s.Open() = %q, want %q", got, want)
    }
}
