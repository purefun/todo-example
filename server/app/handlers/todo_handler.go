package handlers

import (
	"github.com/asynkron/protoactor-go/cluster"
	"github.com/purefun/todo-example/server/app/commands"
)

// impl github.com/purefun/todo-example/server/app/commands.Todo
type TodoHandler struct {
}

func (t *TodoHandler) Init(ctx cluster.GrainContext) {
	panic("not implemented") // TODO: Implement
}

func (t *TodoHandler) Terminate(ctx cluster.GrainContext) {
	panic("not implemented") // TODO: Implement
}

func (t *TodoHandler) ReceiveDefault(ctx cluster.GrainContext) {
	panic("not implemented") // TODO: Implement
}

func (t *TodoHandler) Create(_ *commands.CreateTodoCmd, _ cluster.GrainContext) (*commands.Empty, error) {
	panic("not implemented") // TODO: Implement
}

func (t *TodoHandler) UpdateText(_ *commands.UpdateTodoTextCmd, _ cluster.GrainContext) (*commands.Empty, error) {
	panic("not implemented") // TODO: Implement
}

func (t *TodoHandler) Complete(_ *commands.CompleteTodoCmd, _ cluster.GrainContext) (*commands.Empty, error) {
	panic("not implemented") // TODO: Implement
}

func (t *TodoHandler) Uncomplete(_ *commands.UncompleteTodoCmd, _ cluster.GrainContext) (*commands.Empty, error) {
	panic("not implemented") // TODO: Implement
}

func (t *TodoHandler) Delete(_ *commands.DeleteTodoCmd, _ cluster.GrainContext) (*commands.Empty, error) {
	panic("not implemented") // TODO: Implement
}
