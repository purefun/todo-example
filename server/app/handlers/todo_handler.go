package handlers

import (
	"fmt"
	"reflect"

	"github.com/asynkron/protoactor-go/cluster"
	"github.com/purefun/todo-example/server/app"
	"github.com/purefun/todo-example/server/app/commands"
	"github.com/purefun/todo-example/server/domain"
)

// impl github.com/purefun/todo-example/server/app/commands.Todo
type TodoHandler struct {
	app.EventPublisher
	todo domain.Todo
}

func (t *TodoHandler) Init(ctx cluster.GrainContext) {
}

func (t *TodoHandler) Terminate(ctx cluster.GrainContext) {
}

func (t *TodoHandler) ReceiveDefault(ctx cluster.GrainContext) {
	switch ev := ctx.Message().(type) {
	case app.EventEnvelope:
		t.todo.Apply(ev.Event)
	default:
		fmt.Printf("unhandled message: %s\n", reflect.TypeOf(ev).Name())
	}
}

func (t *TodoHandler) Create(cmd *commands.CreateTodoCmd, ctx cluster.GrainContext) (*commands.Empty, error) {
	t.PublishEvent(&domain.TodoCreated{Text: cmd.Text}, ctx)
	return &commands.Empty{}, nil
}

func (t *TodoHandler) UpdateText(cmd *commands.UpdateTodoTextCmd, ctx cluster.GrainContext) (*commands.Empty, error) {
	t.PublishEvent(&domain.TodoTextUpdated{NewText: cmd.Text}, ctx)
	return &commands.Empty{}, nil
}

func (t *TodoHandler) Complete(_ *commands.CompleteTodoCmd, ctx cluster.GrainContext) (*commands.Empty, error) {
	t.PublishEvent(&domain.TodoCompleted{}, ctx)
	return &commands.Empty{}, nil
}

func (t *TodoHandler) Uncomplete(_ *commands.UncompleteTodoCmd, ctx cluster.GrainContext) (*commands.Empty, error) {
	t.PublishEvent(&domain.TodoUncompleted{}, ctx)
	return &commands.Empty{}, nil
}

func (t *TodoHandler) Delete(_ *commands.DeleteTodoCmd, ctx cluster.GrainContext) (*commands.Empty, error) {
	t.PublishEvent(&domain.TodoDeleted{}, ctx)
	return &commands.Empty{}, nil
}
