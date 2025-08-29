package service

import (
	context "context"
	"go-service/internal/gen"
)

// Hello is the resolver for the hello field.
func (r *queryResolver) GetHold(ctx context.Context, id int) (*gen.Hold, error) {
	return r.Service.GetHold(ctx, id)
}

func (svc *Service) GetHold(ctx context.Context, id int) (*gen.Hold, error) {
	return &gen.Hold{ID: -1, MarcID: 1, KeepsiteID: 1}, nil
}
