package gql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/purefun/todo-example/server/app/commands"
	"github.com/purefun/todo-example/server/gql/generated"
	"github.com/purefun/todo-example/server/gql/model"
)

func (r *mutationResolver) TodoCreate(ctx context.Context, input model.TodoCreateInput) (*model.Output, error) {
	id := uuid.NewString()
	todo := r.Handlers.Todo(id)
	todo.Create(&commands.CreateTodoCmd{Text: input.Text})
	return &model.Output{ID: id}, nil
}

func (r *mutationResolver) TodoUpdateText(ctx context.Context, id string, text string) (*model.Output, error) {
	todo := r.Handlers.Todo(id)
	todo.UpdateText(&commands.UpdateTodoTextCmd{Text: text})
	return &model.Output{ID: id}, nil
}

func (r *mutationResolver) TodoComplete(ctx context.Context, id string) (*model.Output, error) {
	todo := r.Handlers.Todo(id)
	todo.Complete(&commands.CompleteTodoCmd{})
	return &model.Output{ID: id}, nil
}

func (r *mutationResolver) TodoUncomplete(ctx context.Context, id string) (*model.Output, error) {
	todo := r.Handlers.Todo(id)
	todo.Uncomplete(&commands.UncompleteTodoCmd{})
	return &model.Output{ID: id}, nil
}

func (r *mutationResolver) TodoDelete(ctx context.Context, id string) (*model.Output, error) {
	todo := r.Handlers.Todo(id)
	todo.Delete(&commands.DeleteTodoCmd{})
	return &model.Output{ID: id}, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
