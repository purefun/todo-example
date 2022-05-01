package handlers

import (
	"github.com/asynkron/protoactor-go/cluster"
	"github.com/purefun/todo-example/server/app/commands"
)

type Handlers struct {
	cluster *cluster.Cluster
}

func (h *Handlers) Todo(id string) *commands.TodoGrainClient {
	return commands.GetTodoGrainClient(h.cluster, id)
}

func NewHandlers(c *cluster.Cluster) *Handlers {
	return &Handlers{cluster: c}
}
