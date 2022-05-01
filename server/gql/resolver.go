package gql

import (
	"github.com/purefun/todo-example/server/app/handlers"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Handlers *handlers.Handlers
}
