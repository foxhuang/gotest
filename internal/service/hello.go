package service

import (
	context "context"
	"fmt"
)
// Hello is the resolver for the hello field.
func (r *queryResolver) Hello(ctx context.Context) (string, error) {
	return "Hello, GraphQL!", nil
}

// SayHello is the resolver for the sayHello field.
func (r *queryResolver) SayHello(ctx context.Context, name string) (string, error) {
	return fmt.Sprintf("Hello, %s!", name), nil
}
