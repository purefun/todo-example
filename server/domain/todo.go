package domain

import (
	"fmt"
	"reflect"
)

type Todo struct {
	id        string
	text      string
	completed bool
	deleted   bool
}

func (t *Todo) Apply(event interface{}) {
	switch ev := event.(type) {
	case *TodoCreated:
		t.id = ev.ID
		t.text = ev.Text

	case *TodoTextUpdated:
		t.text = ev.NewText

	case *TodoCompleted:
		t.completed = true

	case *TodoUncompleted:
		t.completed = false

	case TodoDeleted:
		t.deleted = true

	default:
		fmt.Printf("unhandled event: %s\n", reflect.TypeOf(event).Name())
	}
}

type TodoCreated struct {
	ID   string
	Text string
}

type TodoCompleted struct {
	ID string
}

type TodoUncompleted struct {
	ID string
}

type TodoTextUpdated struct {
	NewText string
}

type TodoDeleted struct {
	ID string
}