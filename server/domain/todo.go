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
		fmt.Printf("unhandled event: %s, %+v\n", reflect.TypeOf(event).Elem().Name(), event)
	}

	fmt.Printf("%+v\n", t)
}

func (t *Todo) ID() string {
	return t.id
}

func (t *Todo) Text() string {
	return t.text
}

func (t *Todo) Completed() bool {
	return t.completed
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
	ID      string
	NewText string
}

type TodoDeleted struct {
	ID string
}
