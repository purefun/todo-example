package cluster

import (
	"github.com/asynkron/protoactor-go/actor"
	"github.com/asynkron/protoactor-go/cluster"
	"github.com/asynkron/protoactor-go/cluster/clusterproviders/consul"
	"github.com/asynkron/protoactor-go/cluster/identitylookup/disthash"
	"github.com/asynkron/protoactor-go/remote"
	"github.com/purefun/todo-example/server/app/commands"
	"github.com/purefun/todo-example/server/app/handlers"
)

func NewCluster() *cluster.Cluster {

	system := actor.NewActorSystem()
	provider, _ := consul.New()
	lookup := disthash.New()
	config := remote.Configure("localhost", 0)

	// actor.WithReceiverMiddleware()

	todo := commands.NewTodoKind(func() commands.Todo {
		return &handlers.TodoHandler{}
	}, 0, actor.WithMailbox(actor.UnboundedPriority()))

	clusterConfig := cluster.Configure("my-cluster", provider, lookup, config, cluster.WithKinds(todo))
	c := cluster.New(system, clusterConfig)
	c.StartMember()

	return c
}
