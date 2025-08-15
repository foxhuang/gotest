package service

import (
	"context"
	"go-service/gql"
)

func (r *queryResolver) GetReader(ctx context.Context, id int) (*gql.Reader, error) {
	return r.Service.GetReader(ctx, id)
}

func (svc *Service) GetReader(ctx context.Context, id int) (*gql.Reader, error) {
	var r gql.Reader
	query := "SELECT id, name, email FROM reader WHERE id = @p1"
	err1 := svc.DB.QueryRow(query, id).Scan(&r.ID, &r.Name, &r.Email)
	if err1 != nil {
		return &gql.Reader{ID: -1, Name: "error", Email: ""}, nil
	}

	reader := &gql.Reader{
		ID:    r.ID,
		Name:  r.Name,
		Email: r.Email,
	}
	return reader, nil

}
